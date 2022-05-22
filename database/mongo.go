package database

import (
	"context"
	"errors"
	"log"
	"todolist/todolist"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB = "todolists"

type Mongo struct {
	database *mongo.Database
}

func NewMongoDatabase() *Mongo {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	d := client.Database(DB)

	return &Mongo{
		database: d,
	}
}

func (d *Mongo) Create(t todolist.Todolist) (todolist.Todolist, error) {
	t.Id = primitive.NewObjectID().Hex()
	_, err := d.database.Collection("todolists").InsertOne(context.TODO(), t)
	return t, err
}

func (d *Mongo) GetAll() (todolist.Todolists, error) {
	collection, _ := d.database.Collection("todolists").Find(context.TODO(), bson.M{})

	var results []todolist.Todolist
	collection.All(context.TODO(), &results)

	var todolists todolist.Todolists
	todolists.Todolists = results

	return todolists, nil
}

func (d *Mongo) Update(id string, t todolist.Todolist) (todolist.Todolist, error) {
	collection := d.database.Collection("todolists")

	result, _ := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"task", t.Task}}}, // @todo : refactor this part
		},
	)

	if result.MatchedCount == 0 {
		return t, errors.New("id not found")
	}

	return t, nil
}

func (d *Mongo) Delete(id string) (bool, error) {
	collection := d.database.Collection("todolists")

	result, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if result.DeletedCount == 0 {
		return false, errors.New("id not found")
	}

	return true, nil
}
