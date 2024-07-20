package interfaces

type CreatePatientRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,strongpassword"`
}

type CreatePatientResponse struct {
	Id    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

/* ---- DOCTOR LOGIN -----*/

type LoginPatientRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginPatientResponse struct {
	Token string `json:"token" validate:"required"`
}

/* ---- DOCTOR LOGIN -----*/

type SearchPatientStruct struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SearchPatientsResponse struct {
	Patients []SearchPatientStruct `json:"patients"`
}

/* ---- GET PATIENT WITH CONNECTION -----*/

type GetPatientStruct struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetPatientAndConnectionsResponse struct {
	Patient   *GetPatientStruct `json:"patient"`
	Connected bool              `json:"connected"`
}
