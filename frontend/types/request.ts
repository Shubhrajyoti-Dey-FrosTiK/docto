export interface GenericResponse<T = any> {
  success: boolean;
  message: string;
  details: T;
  error: any;
}
