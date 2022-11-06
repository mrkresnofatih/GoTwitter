package models

type AuthLoginRequestModel struct {
	GrantId string `json:"grantId" validate:"required"`
}
