package mapper

import (
	"docto/auth"
	"docto/interfaces"
	"docto/models"
	"strconv"
)

func DoctorModelToCreateDoctorResponse(doctorModel *models.Doctor) *interfaces.CreateDoctorResponse {
	token, _ := auth.GenerateToken(true, false, doctorModel.ID)

	return &interfaces.CreateDoctorResponse{
		Name:        doctorModel.Name,
		Designation: doctorModel.Designation,
		Headline:    doctorModel.Designation,
		Id:          strconv.FormatUint(uint64(doctorModel.ID), 10),
		Email:       doctorModel.Email,
		Token:       token,
	}
}

func DoctorsModelToSearchDoctorsResponse(doctorModels *[]models.Doctor) *interfaces.SearchDoctorsResponse {
	var doctors []interfaces.SearchDoctorStruct

	for _, doctor := range *doctorModels {
		doctors = append(doctors, interfaces.SearchDoctorStruct{
			ID:       strconv.FormatInt(int64(doctor.ID), 10),
			Email:    doctor.Email,
			Name:     doctor.Name,
			Headline: doctor.Headline,
		})
	}

	return &interfaces.SearchDoctorsResponse{
		Doctors: doctors,
	}
}

func CreateGetDoctorByConnectionResponse(doctor *models.Doctor, connected bool) *interfaces.GetDoctorAndConnectionsResponse {
	res := &interfaces.GetDoctorAndConnectionsResponse{
		Connected: connected,
	}

	if doctor != nil {
		res.Doctor = &interfaces.GetDoctorStruct{
			ID:          strconv.FormatInt(int64(doctor.ID), 10),
			Email:       doctor.Email,
			Name:        doctor.Name,
			Headline:    doctor.Headline,
			Designation: doctor.Designation,
		}
	}

	return res
}

func DoctorModelToUserMapper(doctor *models.Doctor) *interfaces.User {
	return &interfaces.User{
		ID:       strconv.FormatInt(int64(doctor.ID), 10),
		Email:    doctor.Email,
		Name:     doctor.Name,
		UserType: "DOCTOR",
	}
}

func DoctorModelsToUserMapper(doctors *[]models.Doctor) *[]interfaces.User {
	users := []interfaces.User{}

	for _, doctor := range *doctors {
		users = append(users, *DoctorModelToUserMapper(&doctor))
	}

	return &users
}
