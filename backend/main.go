package main

import (
	"fmt"
	"os"

	"docto/constants"
	"docto/db"
	_ "docto/docs"
	"docto/handler"
	"docto/interfaces"
	"docto/middleware"
	"docto/s3"
	"docto/util"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	godotenv.Load()

	s3 := s3.Connect()
	db := db.Connect()

	app := fiber.New(fiber.Config{
		Prefork:           true,
		JSONEncoder:       json.Marshal,
		EnablePrintRoutes: true,
		JSONDecoder:       json.Unmarshal,
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(interfaces.GetGenericResponse(false, "ERROR", nil, err))
		},
	})

	// Allow origin
	app.Use(cors.New(util.DefaultCors()))

	// Recover from panics || Comment this out to check panic logs
	// app.Use(recover.New())

	// Rate limiting
	app.Use(limiter.New(limiter.Config{Max: constants.REQUEST_RATE}))

	// Compress responses
	app.Use(compress.New())

	// Security
	app.Use(helmet.New())

	// Health check
	app.Use(healthcheck.New())

	handler := &handler.Handler{
		DB: db,
		S3: s3,
	}

	/* ---- UNAUTHENTICATD ROUTES ---- */

	// Health
	app.Get("/health", handler.HealthCheckHandler)

	// Monitor
	app.Get("/metrics", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Doctor
	app.Post("/doctor/create", middleware.ValidateBody[interfaces.CreateDoctorRequest], handler.HandleCreateDoctor)
	app.Post("/doctor/login", middleware.ValidateBody[interfaces.LoginDoctorRequest], handler.HandleLoginDoctor)

	// Patient
	app.Post("/patient/create", middleware.ValidateBody[interfaces.CreatePatientRequest], handler.HandleCreatePatient)
	app.Post("/patient/login", middleware.ValidateBody[interfaces.LoginPatientRequest], handler.HandleLoginPatient)

	/* ---- AUTH MIDDLEWARE ----*/

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(constants.AUTH_JWT_SECRET)},
	}))

	/* ---- UTHENTICATD ROUTES ---- */

	// Assign
	app.Post("/assign/patient", middleware.CheckDoctor, middleware.ValidateBody[interfaces.AssignPatientRequest], handler.HandleAssignPatient)
	app.Post("/assign/doctor", middleware.CheckPatient, middleware.ValidateBody[interfaces.AssignDoctorRequest], handler.HandleAssignDoctor)

	// Auth
	app.Get("/token/verify", handler.HandleVerifyToken)

	// Doctor
	app.Post("/doctor/upload", middleware.CheckDoctor, handler.HandleDoctorUploadFile)
	app.Get("/doctor/populated/patient", middleware.CheckDoctor, handler.HandleGetDoctorWithPatientPopulated)
	app.Get("/doctor/populated/connection", middleware.CheckPatient, handler.HandleGetDoctorWithConnectionPopulated)
	app.Get("/doctor/search", handler.HandleSearchDoctors)
	app.Get("/doctor/connectedPatients", middleware.CheckDoctor, handler.HandleGetAssociatedUsersForDoctors)
	app.Get("/doctor/files", middleware.CheckDoctor, handler.HandleGetAssociatedFilesForDoctors)

	// Patient
	app.Post("/patient/upload", middleware.CheckPatient, handler.HandlePatientUploadFile)
	app.Get("/patient/populated/connection", middleware.CheckDoctor, handler.HandleGetPatientWithConnectionPopulated)
	app.Get("/patient/search", handler.HandleSearchPatients)
	app.Get("/patient/connectedDoctors", middleware.CheckPatient, handler.HandleGetAssociatedUsersForPatients)
	app.Get("/patient/files", middleware.CheckPatient, handler.HandleGetAssociatedFilesForPatients)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting Server on PORT : ", port)

	app.Listen(":" + port)
}
