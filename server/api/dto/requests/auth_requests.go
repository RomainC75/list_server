package requests

import "time"

type SignupRequest struct {
	Email    string     `json:"email" binding:"required,email"`
	Password string     `json:"password" binding:"required,min=6"`
	Birthday *time.Time `json:"birthday" `
}
