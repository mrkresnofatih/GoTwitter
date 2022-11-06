package models

type FollowCreateRequestModel struct {
	Follower string `json:"follower" validate:"required"`
	Username string `json:"username" validate:"required"`
}
