export interface TokenVerifyResponse {
  token: string;
  role: UserType;
  userId: string;
}

export enum UserType {
  UNKNOWN = "UNKNOWN",
  DOCTOR = "DOCTOR",
  PATIENT = "PATIENT",
}
