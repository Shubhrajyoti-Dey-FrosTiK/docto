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






