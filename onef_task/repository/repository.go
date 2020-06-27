package repository

import (
	"context"

	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_task"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	DbName  = "onef"
	ColName = "task"
)

type taskRepository struct {
	Client *mongo.Client
}

func NewRepository(Client *mongo.Client) onef_task.Repository {
	return &taskRepository{Client}
}

func (repo *taskRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *taskRepository) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	cur, err := repo.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *taskRepository) Get(id int) (*model.Task, error) {
	var task *model.Task
	task.Id = id
	filter := bson.M{"id": id}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *taskRepository) Create(task *model.Task) error {
	_, err := repo.collection().InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func (repo *taskRepository) Update(task *model.Task) error {
	filter := bson.M{"_id": task.Id}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": task})
	if err != nil {
		return err
	}
	return nil
}
