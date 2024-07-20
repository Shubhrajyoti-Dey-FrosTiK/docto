package auth

import (
	"docto/constants"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(doctor bool, patient bool, id uint) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"doctor":  doctor,
		"patient": patient,
		"exp":     time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	// Return the signed token
	return token.SignedString([]byte(constants.AUTH_JWT_SECRET))
}

func (auth *Auth) GetId() (*uint, error) {
	user := auth.Ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("corrupted token")
	}

	uintId := uint(id)

	return &uintId, nil
}

func (auth *Auth) IsDoctor() bool {
	user := auth.Ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["doctor"].(bool)
}

func (auth *Auth) IsPatient() bool {
	user := auth.Ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["patient"].(bool)
}
