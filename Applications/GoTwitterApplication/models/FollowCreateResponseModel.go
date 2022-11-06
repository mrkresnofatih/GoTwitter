package models

type FollowCreateResponseModel struct {
	Follower string `json:"follower"`
	Username string `json:"username"`
}
