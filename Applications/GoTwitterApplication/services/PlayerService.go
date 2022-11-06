package services

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/entities"
	"github.com/mrkresnofatih/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlayerService interface {
	CreatePlayer(ctx context.Context, model models.PlayerCreateRequestModel) (models.PlayerCreateResponseModel, error)
	GetPlayer(ctx context.Context, model models.PlayerGetRequestModel) (models.PlayerGetResponseModel, error)
	UpdatePlayer(ctx context.Context, model models.PlayerUpdateRequestModel) (models.PlayerUpdateResponseModel, error)
}

type PlayerService struct {
	MongoDb *mongo.Database
}

func (p *PlayerService) CreatePlayer(ctx context.Context, model models.PlayerCreateRequestModel) (models.PlayerCreateResponseModel, error) {
	newPlayer := entities.Player{
		Username:  model.Username,
		ImageUrl:  model.ImageUrl,
		FullName:  model.FullName,
		Bio:       model.Bio,
		CreatedAt: model.CreatedAt,
	}
	_, err := p.MongoDb.Collection(PlayerCollectionName).InsertOne(ctx, newPlayer)
	if err != nil {
		log.Error("failed to insert player to mongodb")
		return *new(models.PlayerCreateResponseModel), errors.New("failed to insert player to mongodb")
	}
	return models.PlayerCreateResponseModel{
		Username:  newPlayer.Username,
		FullName:  newPlayer.FullName,
		Bio:       newPlayer.Bio,
		ImageUrl:  newPlayer.ImageUrl,
		CreatedAt: newPlayer.CreatedAt,
	}, nil
}

func (p *PlayerService) GetPlayer(ctx context.Context, model models.PlayerGetRequestModel) (models.PlayerGetResponseModel, error) {
	var player entities.Player

	err := p.MongoDb.Collection(PlayerCollectionName).
		FindOne(ctx, bson.D{{"username", model.Username}}).
		Decode(&player)
	if err != nil {
		log.Error("failed to find player with username")
		return *new(models.PlayerGetResponseModel), errors.New("failed to find player to mongodb")
	}
	return models.PlayerGetResponseModel{
		Username:  player.Username,
		FullName:  player.FullName,
		Bio:       player.Bio,
		ImageUrl:  player.ImageUrl,
		CreatedAt: player.CreatedAt,
	}, nil
}

func (p *PlayerService) UpdatePlayer(ctx context.Context, model models.PlayerUpdateRequestModel) (models.PlayerUpdateResponseModel, error) {
	foundPlayer, err := p.GetPlayer(ctx, models.PlayerGetRequestModel{
		Username: model.Username,
	})
	if err != nil {
		log.Error("failed to find player with provided username, cannot update")
		return *new(models.PlayerUpdateResponseModel), nil
	}

	filter := bson.D{{"username", model.Username}}
	update := bson.D{
		{"fullName", model.FullName},
		{"imageUrl", model.ImageUrl},
		{"bio", model.Bio},
	}
	_, err = p.MongoDb.Collection(PlayerCollectionName).
		UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("failed to update player to mongodb")
		return *new(models.PlayerUpdateResponseModel), nil
	}
	return models.PlayerUpdateResponseModel{
		Username:  model.Username,
		FullName:  model.FullName,
		Bio:       model.Bio,
		ImageUrl:  model.ImageUrl,
		CreatedAt: foundPlayer.CreatedAt,
	}, nil
}

const (
	PlayerCollectionName = "players"
)
