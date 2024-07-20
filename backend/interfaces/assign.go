package interfaces

type AssignPatientRequest struct {
	PatientId string `json:"patientId" validate:"required"`
}

type AssignDoctorRequest struct {
	DoctorId string `json:"doctorId" validate:"required"`
}
