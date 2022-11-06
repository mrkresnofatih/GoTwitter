package models

type PlayerGetResponseModel struct {
	Username  string `bson:"username" json:"username"`
	FullName  string `bson:"fullName" json:"fullName"`
	ImageUrl  string `bson:"imageUrl" json:"imageUrl"`
	Bio       string `bson:"bio" json:"bio"`
	CreatedAt string `bson:"createdAt" json:"createdAt"`
}
