package models

type GenerateResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string
}
