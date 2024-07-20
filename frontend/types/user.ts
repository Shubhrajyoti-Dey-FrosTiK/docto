import { UserType } from "./auth";

export interface User {
  id: string;
  name: string;
  userType: UserType;
  email: string;
}

export interface GetAssociatedUsersResponse {
  users: User[];
}
