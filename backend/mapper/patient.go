package mapper

import (
	"docto/auth"
	"docto/interfaces"
	"docto/models"
	"strconv"
)

func PatientModelToCreatePatientResponse(patientModel *models.Patient) *interfaces.CreatePatientResponse {
	token, _ := auth.GenerateToken(false, true, patientModel.ID)

	return &interfaces.CreatePatientResponse{
		Name:  patientModel.Name,
		Email: patientModel.Email,
		Id:    strconv.FormatUint(uint64(patientModel.ID), 10),
		Token: token,
	}
}

func PatientsModelToSearchDoctorsResponse(patientsModels *[]models.Patient) *interfaces.SearchPatientsResponse {
	var patients []interfaces.SearchPatientStruct

	for _, doctor := range *patientsModels {
		patients = append(patients, interfaces.SearchPatientStruct{
			ID:    strconv.FormatInt(int64(doctor.ID), 10),
			Email: doctor.Email,
			Name:  doctor.Name,
		})
	}

	return &interfaces.SearchPatientsResponse{
		Patients: patients,
	}
}

func CreateGetPatientByConnectionResponse(doctor *models.Patient, connected bool) *interfaces.GetPatientAndConnectionsResponse {
	res := &interfaces.GetPatientAndConnectionsResponse{
		Connected: connected,
	}

	if doctor != nil {
		res.Patient = &interfaces.GetPatientStruct{
			ID:    strconv.FormatInt(int64(doctor.ID), 10),
			Email: doctor.Email,
			Name:  doctor.Name,
		}
	}

	return res
}

func PatientModelToUserMapper(patient *models.Patient) *interfaces.User {
	return &interfaces.User{
		ID:       strconv.FormatInt(int64(patient.ID), 10),
		Email:    patient.Email,
		Name:     patient.Name,
		UserType: "PATIENT",
	}
}

func PatientModelsToUserMapper(patients *[]models.Patient) *[]interfaces.User {
	users := []interfaces.User{}

	for _, patient := range *patients {
		users = append(users, *PatientModelToUserMapper(&patient))
	}

	return &users
}
