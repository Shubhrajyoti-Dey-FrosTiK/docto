export interface CreatePatientRequest {
  name: string;
  password: string;
  email: string;
}

export interface GetPatientStruct
  extends Omit<CreatePatientRequest, "password"> {
  id: string;
}

export interface SearchPatientResponse {
  patients: GetPatientStruct[];
}

export interface GetPatientWithConnections {
  patient: GetPatientStruct;
  connected: boolean;
}
