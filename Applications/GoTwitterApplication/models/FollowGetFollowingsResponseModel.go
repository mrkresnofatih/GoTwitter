package models

type FollowGetFollowingsResponseModel struct {
	Username   string                            `json:"username"`
	Followings []FollowGetFollowingResponseModel `json:"followings"`
}
