package interfaces

/* --- CREATE DOCTOR ---*/

type CreateDoctorRequest struct {
	Name        string `json:"name" validate:"required"`
	Headline    string `json:"headline" validate:"required"`
	Designation string `json:"designation" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,strongpassword"`
}

type CreateDoctorResponse struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Headline    string `json:"headline" validate:"required"`
	Designation string `json:"designation" validate:"required"`
	Email       string `json:"email" required:"true"`
	Token       string `json:"token" validate:"required"`
}

/* ---- DOCTOR LOGIN -----*/

type LoginDoctorRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginDoctorResponse struct {
	Token string `json:"token" validate:"required"`
}

/* ---- DOCTOR SEARCH -----*/

type SearchDoctorStruct struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Headline string `json:"headline"`
	Email    string `json:"email"`
}

type SearchDoctorsResponse struct {
	Doctors []SearchDoctorStruct `json:"doctors"`
}

/* ---- GET DOCTOR AND CONNECTION -----*/

type GetDoctorStruct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Headline    string `json:"headline"`
	Email       string `json:"email"`
	Designation string `json:"designation"`
}

type GetDoctorAndConnectionsResponse struct {
	Doctor    *GetDoctorStruct `json:"doctor"`
	Connected bool             `json:"connected"`
}
