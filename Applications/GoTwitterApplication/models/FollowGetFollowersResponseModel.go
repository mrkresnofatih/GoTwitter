package models

type FollowGetFollowersResponseModel struct {
	Username  string                           `json:"username"`
	Followers []FollowGetFollowerResponseModel `json:"followers"`
}
