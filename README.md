# Docto

A platform to connect doctors and patient

## Prerequisites

1. PostgresSQL
2. NodeJS
3. Golang
4. Yarn
5. AWS Account


## Setup

## DB setup

Start the `pgsql` server and create a table named `docto` and you are ready to start setting up the `backend`

## Backend setup
Go into the `backend` diretory and set the following env variables

```
export PORT="8084"
export DB_PORT="..." // Default is 5432
export DB_NAME="docto"
export DB_HOST="localhost"
export AWS_ACCESS_KEY_ID="..."
export AWS_SECRET_ACCESS_KEY="..." // Can get from AWS IAM Console
export AWS_BUCKET_NAME="..." // Can get from AWS IAM Console
export JWT_SECRET="..." // can be any string
```

## Frontend setup
Go into the `frontend` directory and run `yarn` to install the dependencies
```
yarn
```

Now start the server with 

```
yarn dev
```

Now the server will be running at `http://localhost:3000`

### Note
The `BACKEND_URL` is already set to `8084` in the `frontend` env as well as the `backend` env


## Features
1. Authentication (Signup + Login)
2. Connect to doctors/patient (doctors if a patient and opposite otherwise)
3. View connections
4. File Upload

## Walkthrough

1. Homepage (Not logged in)

![Capture-2024-07-20-192109](https://github.com/user-attachments/assets/429e5769-b32e-4f9c-a71a-1a54cea12125)

2. Signup Page

![Capture-2024-07-20-192624](https://github.com/user-attachments/assets/94bca98b-b06b-4c7e-a8bb-9e51ae52ffd8)

3. Login Page

![Capture-2024-07-20-192732](https://github.com/user-attachments/assets/e351c5ee-bc47-4167-bc91-2209f4b8b21a)

4. Home Page (Logged in)

![Capture-2024-07-20-192824](https://github.com/user-attachments/assets/e467476d-188e-4670-bf82-98ffffd1ee49)

5. Search doctors/patients

![Capture-2024-07-20-192942](https://github.com/user-attachments/assets/a25dcd27-8f97-43e2-8fef-7210db84d8a5)

6. View connections

![Capture-2024-07-20-193010](https://github.com/user-attachments/assets/3d256b5e-2760-46c2-a30f-095c9afdbf36)

7. File explorer

![Capture-2024-07-20-193242](https://github.com/user-attachments/assets/25f97602-0103-4d97-8969-6b2c52e6250d)


## Implementation

### DB Schema

```.go

type Doctor struct {
	gorm.Model

	ID          uint      `gorm:"primaryKey"` // id of the doctor
	Name        string    // name of the doctor
	Email       string    `gorm:"uniqueIndex"` // email of the patient
	Designation string    // eg. MBBS etc
	Headline    string    // eg. Chief of surgery at AIMS Delhi
	Patients    []Patient `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_patients;"`
	Files       []File    `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_files;"`

	// Sensitive
	Password string // Will be hashed before being inserted

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}


type File struct {
	gorm.Model

	ID       uint `gorm:"primaryKey"` // id of the file
	Url      string
	Key      string
	FileName string

	DoctorID  uint
	PatientID uint

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}


type Patient struct {
	gorm.Model

	ID      uint     `gorm:"primaryKey"` // id of the patient
	Name    string   // name of the patient
	Email   string   `gorm:"uniqueIndex"` // email of the patient
	Files   []File   `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:patient_files;"`
	Doctors []Doctor `gorm:"OnDelete:SET ARRAY[]::varchar[]; many2many:doctor_patients;"`

	// Sensitive
	Password string // Will be hashed before being inserted

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

```

## Endpoints
### Non-Authenticated
```
/health  -> health of the server
/metrics -> metrics of usage of the server

/doctor/create -> Creates a doctor
/doctor/login  -> Logs in a doctor

/patient/create -> Creates a patient
/patient/login  -> Logs in a patient
```

### Authenticated Routes
```
/assign/patient -> assigns a patient to a doctor (invoker must be a doctor)
/assign/doctor  -> assigns a doctor to a patient (invoker must be a patient)

/token/verify -> provides information like validity, userId, userType (doctor/patient) from a token

/doctor/upload               -> to upload files and attach it to a doctor
/doctor/populated/patient    -> to fetch the doctor and its associated patients
/doctor/populated/connection -> to fetch the doctor info and check if the invoker (patient) is connected witn the doctor
/doctor/search               -> search doctors using name, email and id
/doctor/connectedPatients    -> fetch the connected patients of a doctor
/doctor/files                -> fetch the files of a doctor

/patient/upload              -> to upload files and attach it to a patient
/patient/populated/connection-> to fetch the patient info and check if the invoker (doctor) is connected witn the patient
/patient/search              -> search patients using name, email and id
/patient/connectedDoctors    -> fetch the connected doctors of a patient
/patient/files               -> fetch the files of a patient
```
## Authentication

A simple `JWT` token based authenticated is used in the project for simplicity which can be enhanced by using firebase in the project in future to have `Google` login and other types of login.
