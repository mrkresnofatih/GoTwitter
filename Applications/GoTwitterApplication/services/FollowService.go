package services

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/entities"
	"github.com/mrkresnofatih/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IFollowService interface {
	Follow(ctx context.Context, followReq models.FollowCreateRequestModel) (models.FollowCreateResponseModel, error)
	Unfollow(ctx context.Context, unfollowReq models.FollowDeleteRequestModel) (models.FollowDeleteResponseModel, error)
	GetFollowers(ctx context.Context, getFollowersReq models.FollowGetFollowersRequestModel) (models.FollowGetFollowersResponseModel, error)
	GetFollowings(ctx context.Context, getFollowingsReq models.FollowGetFollowingsRequestModel) (models.FollowGetFollowingsResponseModel, error)
}

type FollowService struct {
	MongoDb *mongo.Database
}

func (f *FollowService) Follow(ctx context.Context, followReq models.FollowCreateRequestModel) (models.FollowCreateResponseModel, error) {
	log.Info("Start follow")
	newFollow := entities.Follow{
		Follower:  followReq.Follower,
		Username:  followReq.Username,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	_, err := f.MongoDb.Collection(FollowCollectionName).InsertOne(ctx, newFollow)
	if err != nil {
		log.Error("failed to insert follow to mongodb")
		return *new(models.FollowCreateResponseModel), errors.New("failed to insert follow to mongodb")
	}
	log.Info("finished follow")
	return models.FollowCreateResponseModel{
		Username: newFollow.Username,
		Follower: newFollow.Follower,
	}, nil
}

func (f *FollowService) Unfollow(ctx context.Context, unfollowReq models.FollowDeleteRequestModel) (models.FollowDeleteResponseModel, error) {
	log.Info("Start unfollow")
	_, err := f.MongoDb.Collection(FollowCollectionName).
		DeleteOne(ctx, bson.D{
			{"follower", unfollowReq.Follower},
			{"username", unfollowReq.Username},
		})
	if err != nil {
		log.Error("failed to delete follow from mongodb")
		return *new(models.FollowDeleteResponseModel), errors.New("failed to delete follow from mongodb")
	}
	log.Info("finished unfollow")
	return models.FollowDeleteResponseModel{
		Username: unfollowReq.Username,
		Follower: unfollowReq.Follower,
	}, nil
}

func (f *FollowService) GetFollowers(ctx context.Context, getFollowersReq models.FollowGetFollowersRequestModel) (models.FollowGetFollowersResponseModel, error) {
	log.Info("Start GetFollowers")
	matchStage := bson.D{{"$match", bson.D{{"username", getFollowersReq.Username}}}}
	skipStage := bson.D{{"$skip", int64(getFollowersReq.PageSize * (getFollowersReq.Page - 1))}}
	limitStage := bson.D{{"$limit", int64(getFollowersReq.PageSize)}}
	lookupStage := bson.D{{"$lookup", bson.D{
		{"from", PlayerCollectionName},
		{"localField", "follower"},
		{"foreignField", "username"},
		{"as", "profile"},
	}}}
	unwindStage := bson.D{{"$unwind", "$profile"}}

	cursor, err := f.MongoDb.Collection(FollowCollectionName).
		Aggregate(ctx, mongo.Pipeline{matchStage, skipStage, limitStage, lookupStage, unwindStage})
	if err != nil {
		log.Error(err)
		log.Error("failed to aggregate follow docs with player docs")
		return *new(models.FollowGetFollowersResponseModel), errors.New("failed to aggregate follow docs with player docs")
	}

	var followerGetProfiles []models.FollowGetFollowerResponseModel
	if err = cursor.All(ctx, &followerGetProfiles); err != nil {
		log.Error(err)
		log.Error("failed to parse cursor docs with follower-get-profiles")
		return *new(models.FollowGetFollowersResponseModel), errors.New("failed to parse cursor docs with follower-get-profiles")
	}

	log.Info(followerGetProfiles)

	log.Info("finished GetFollowers")
	return models.FollowGetFollowersResponseModel{
		Username:  getFollowersReq.Username,
		Followers: followerGetProfiles,
	}, nil
}

func (f *FollowService) GetFollowings(ctx context.Context, getFollowingsReq models.FollowGetFollowingsRequestModel) (models.FollowGetFollowingsResponseModel, error) {
	//TODO implement me
	panic("implement me")
}

const (
	FollowCollectionName = "follows"
)
