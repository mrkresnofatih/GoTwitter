package models

type FollowDeleteRequestModel struct {
	Follower string `json:"follower" validate:"required"`
	Username string `json:"username" validate:"required"`
}
