package models

type FollowGetFollowersRequestModel struct {
	Username string `json:"username" validate:"required"`
	Page     int    `json:"page" validate:"required,gt=0"`
	PageSize int    `json:"pageSize" validate:"required,gt=0,lt=20"`
}
