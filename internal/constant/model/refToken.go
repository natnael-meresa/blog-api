package model

type TokenReqBody struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
