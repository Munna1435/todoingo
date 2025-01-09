package tododatastore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"todoingo/models"
	"todoingo/repositories/db"
)

const (
	DATABASE        = "practice"
	TODO_COLLECTION = "todos"
)

type TodoDataStore struct {
	collection *mongo.Collection
}

func New() *TodoDataStore {
	client := db.GetMongoClient(os.Getenv("MONGODB_URI"))
	return &TodoDataStore{collection: client.Database(DATABASE).Collection(TODO_COLLECTION)}
}

func (r *TodoDataStore) GetAllTodos(userId string) ([]models.TodoData, error) {
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"userId", objectID}, {"isDeleted", false}}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var todos []TodoData
	if err = cursor.All(context.TODO(), &todos); err != nil {
		return nil, err
	}

	return mapTodosData(todos), nil
}

func (r *TodoDataStore) GetTodo(userId string, todoId string) (*models.TodoData, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	todoObjectID, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"userId", userObjectID}, {"_id", todoObjectID}, {"isDeleted", false}}

	var todo TodoData

	err = r.collection.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return mapTodoData(&todo), nil
}

func (r *TodoDataStore) CreateToDo(todo models.TodoData) (*models.TodoData, error) {
	objectID, err := primitive.ObjectIDFromHex(todo.UserId)
	if err != nil {
		return nil, err
	}
	todoModel := TodoData{
		UserId:      objectID,
		Task:        todo.Task,
		IsDeleted:   false,
		IsCompleted: false,
	}

	result, err := r.collection.InsertOne(context.TODO(), todoModel)
	if err != nil {
		return nil, err
	}
	todoModel.Id = (result.InsertedID).(primitive.ObjectID)
	return mapTodoData(&todoModel), nil
}

func (r *TodoDataStore) UpdateToDo(todo models.TodoData) error {
	todoId, err := primitive.ObjectIDFromHex(todo.Id)
	userId, err := primitive.ObjectIDFromHex(todo.UserId)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", todoId}, {"userId", userId}, {"isDeleted", false}}
	update := bson.M{"$set": bson.M{
		"task":        todo.Task,
		"isCompleted": todo.IsCompleted,
	}}
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoDataStore) DeleteToDo(todo models.TodoData) error {
	todoId, err := primitive.ObjectIDFromHex(todo.Id)
	userId, err := primitive.ObjectIDFromHex(todo.UserId)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", todoId}, {"userId", userId}}
	update := bson.M{"$set": bson.M{
		"isDeleted": true,
	}}

	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
