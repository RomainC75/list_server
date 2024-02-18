package responses

type AuthLoginResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthVerifyResponse struct {
	Id    float64 `json:"id"`
	Email string  `json:"email"`
}
