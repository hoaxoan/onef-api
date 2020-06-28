package repository

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_auth"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	DbName  = "onef"
	ColName = "user"
)

type userRepository struct {
	Client *mongo.Client
}

func NewRepository(Client *mongo.Client) onef_auth.Repository {
	return &userRepository{Client}
}

func (repo *userRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *userRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	cur, err := repo.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) Get(id string) (*model.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *model.User
	filter := bson.M{"_id": objId}
	if err := repo.collection().FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*model.User, error) {
	var user *model.User
	filter := bson.M{"email": email}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) GetByUserName(userName string) (*model.User, error) {
	var user *model.User
	filter := bson.M{"username": userName}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository) Create(user *model.User) error {
	_, err := repo.collection().InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Update(user *model.User) error {
	filter := bson.M{"email": user.Email}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) IsEmailToken(email string) bool {
	var user *model.User
	filter := bson.M{"email": email}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	}

	if user != nil && user.Email == email {
		return true
	}

	return false
}
