package models

type OauthOpenIdResponseModel struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	ImageUrl string `json:"imageUrl"`
}
