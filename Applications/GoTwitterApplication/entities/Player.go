package entities

type Player struct {
	Username  string `bson:"username"`
	FullName  string `bson:"fullName"`
	ImageUrl  string `bson:"imageUrl"`
	Bio       string `bson:"bio"`
	CreatedAt string `bson:"createdAt"`
}
