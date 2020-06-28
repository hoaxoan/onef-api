package repository

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_common"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName            = "onef"
	ColCategoryName   = "category"
	ColMoodName       = "mood"
	ColColorRangeName = "color_range"
)

type commonRepository struct {
	Client *mongo.Client
}

func NewRepository(Client *mongo.Client) onef_common.Repository {
	return &commonRepository{Client}
}

func (repo *commonRepository) collectionCategory() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColCategoryName)
}

func (repo *commonRepository) collectionMood() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColMoodName)
}

func (repo *commonRepository) collectionColorRange() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColColorRangeName)
}

// Categories
func (repo *commonRepository) GetCategories(req *model.CategoryRequest) ([]*model.Category, error) {
	var categories []*model.Category
	cur, err := repo.collectionCategory().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *commonRepository) CreateCategories(categories []*model.Category) error {
	for i := range categories {
		_, err := repo.collectionCategory().InsertOne(context.TODO(), categories[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *commonRepository) CreateCategory(category *model.Category) error {
	_, err := repo.collectionCategory().InsertOne(context.TODO(), category)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) UpdateCategory(category *model.Category) error {
	filter := bson.M{"_id": category.Id}
	_, err := repo.collectionCategory().UpdateOne(context.TODO(), filter, bson.M{"$set": category})
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) DeleteCategory(category *model.Category) error {
	filter := bson.M{"_id": category.Id}
	_, err := repo.collectionCategory().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// Moods
func (repo *commonRepository) GetMoods(req *model.MoodRequest) ([]*model.Mood, error) {
	var moods []*model.Mood
	cur, err := repo.collectionMood().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &moods)
	if err != nil {
		return nil, err
	}
	return moods, nil
}

func (repo *commonRepository) CreateMoods(moods []*model.Mood) error {
	for i := range moods {
		_, err := repo.collectionMood().InsertOne(context.TODO(), moods[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *commonRepository) CreateMood(mood *model.Mood) error {
	_, err := repo.collectionMood().InsertOne(context.TODO(), mood)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) UpdateMood(mood *model.Mood) error {
	filter := bson.M{"_id": mood.Id}
	_, err := repo.collectionMood().UpdateOne(context.TODO(), filter, bson.M{"$set": mood})
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) DeleteMood(mood *model.Mood) error {
	filter := bson.M{"_id": mood.Id}
	_, err := repo.collectionMood().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// Color Ranges
func (repo *commonRepository) GetColorRanges(req *model.ColorRangeRequest) ([]*model.ColorRange, error) {
	var colorRanges []*model.ColorRange
	cur, err := repo.collectionColorRange().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &colorRanges)
	if err != nil {
		return nil, err
	}
	return colorRanges, nil
}

func (repo *commonRepository) CreateColorRanges(colorRanges []*model.ColorRange) error {
	for i := range colorRanges {
		_, err := repo.collectionColorRange().InsertOne(context.TODO(), colorRanges[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *commonRepository) CreateColorRange(colorRange *model.ColorRange) error {
	_, err := repo.collectionColorRange().InsertOne(context.TODO(), colorRange)
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) UpdateColorRange(colorRange *model.ColorRange) error {
	filter := bson.M{"_id": colorRange.Id}
	_, err := repo.collectionColorRange().UpdateOne(context.TODO(), filter, bson.M{"$set": colorRange})
	if err != nil {
		return err
	}
	return nil
}

func (repo *commonRepository) DeleteColorRange(colorRange *model.ColorRange) error {
	filter := bson.M{"_id": colorRange.Id}
	_, err := repo.collectionColorRange().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
