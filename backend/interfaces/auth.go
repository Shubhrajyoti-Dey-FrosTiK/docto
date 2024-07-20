package interfaces

type TokenVerifyResponse struct {
	Token  string  `json:"token"`
	Role   *string `json:"role"`
	UserId *uint   `json:"userId"`
}
