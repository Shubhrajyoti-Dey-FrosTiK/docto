export interface CreateDoctorRequest {
  name: string;
  designation: string;
  headline: string;
  password: string;
  email: string;
}

export interface CreateDoctorResponse {
  id: string;
  name: string;
  designation: string;
  headline: string;
  token: string;
  email: string;
}

export interface LoginUserRequest {
  email: string;
  password: string;
}

export interface DoctorSearchStruct {
  id: string;
  name: string;
  headline: string;
  email: string;
}

export interface DoctorSearchResponse {
  doctors?: DoctorSearchStruct[];
}

export interface GetDoctorStruct extends DoctorSearchStruct {
  designation: string;
}

export interface GetDoctorWithConnectionsResponse {
  connected: boolean;
  doctor?: GetDoctorStruct;
}
