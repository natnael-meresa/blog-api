package model

type ResetEmail struct {
	Email string `json:"email" binding:"required"`
}

type PasswordReset struct {
	Password string `json:"password" binding:"required"`
	Confirm  string `json:"confirm" binding:"required"`
}
