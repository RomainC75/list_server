package responses

type AuthLoginResponse struct {
	Id    int32  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthVerifyResponse struct {
	Id    int32  `json:"id"`
	Email string `json:"email"`
}
