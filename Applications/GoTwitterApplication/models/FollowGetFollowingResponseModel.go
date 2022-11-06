package models

type FollowGetFollowingResponseModel struct {
	Username string                 `json:"following" bson:"username"`
	Profile  PlayerGetResponseModel `json:"profile" bson:"profile"`
}
