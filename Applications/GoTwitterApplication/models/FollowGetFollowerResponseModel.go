package models

type FollowGetFollowerResponseModel struct {
	Follower string                 `json:"follower" bson:"follower"`
	Profile  PlayerGetResponseModel `json:"profile" bson:"profile"`
}
