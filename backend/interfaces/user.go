package interfaces

/* ----- GET USER ----- */

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserType string `json:"userType"`
}

/* ---- GET PATIENTS AS USERS -----*/

type GetAssociatedUsersResponse struct {
	Users []User `json:"users"`
}
