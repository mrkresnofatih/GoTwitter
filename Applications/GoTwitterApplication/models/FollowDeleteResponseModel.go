package models

type FollowDeleteResponseModel struct {
	Follower string `json:"follower"`
	Username string `json:"username"`
}
