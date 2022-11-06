package entities

type Follow struct {
	Follower  string `bson:"follower"`
	Username  string `bson:"username"`
	CreatedAt string `bson:"createdAt"`
}
