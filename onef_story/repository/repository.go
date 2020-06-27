package repository

import (
	"github.com/hoaxoan/onef-api/onef_story/delivery/http"
	"github.com/hoaxoan/onef-api/onef_story/repository"
	"github.com/hoaxoan/onef-api/onef_story/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)
)

const (
	DbName  = "onef"
	ColName = "story"
)

type storyRepository struct {
	Client *mongo.Client
}

func NewRepository(Client *mongo.Client) onef_story.Repository {
	return &storyRepository{Client}
}

func (repo *storyRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *storyRepository) GetAll() ([]*model.Story, error) {
	var stoies []*model.Story
	cur, err := repo.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &stoies)
	if err != nil {
		return nil, err
	}
	return stoies, nil
}

func (repo *storyRepository) Get(id int) (*model.Story, error) {
	var story *model.Story
	story.Id = id
	filter := bson.M{"id": id}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func (repo *storyRepository) Create(story *model.Story) error {
	_, err := repo.collection().InsertOne(context.TODO(), story)
	if err != nil {
		return err
	}
	return nil
}

func (repo *storyRepository) Update(story *model.Story) error {
	filter := bson.M{"_id": story.Id}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": story})
	if err != nil {
		return err
	}
	return nil
}
