package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"html"
	"log"
	"strings"
	"time"
)

const CollectionName = "TodoList"

type TodoList struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title, omitempty"`
	Description string             `json:"description" bson:"description, omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at, omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at, omitempty"`
}

func (tl *TodoList) Prepare() {
	tl.Title = html.EscapeString(strings.TrimSpace(tl.Title))
	tl.Description = html.EscapeString(strings.TrimSpace(tl.Description))
	tl.CreatedAt = time.Now()
	tl.UpdatedAt = time.Now()
}

func (tl *TodoList) Validate() map[string]string {
	var err error
	var errorMessages = make(map[string]string)

	if tl.Title == "" {
		err = errors.New("required title")
		errorMessages["Required_title"] = err.Error()
	}
	if tl.Description == "" {
		err = errors.New("required description")
		errorMessages["Required_description"] = err.Error()
	}

	return errorMessages
}

func (tl *TodoList) CreateTodoList(context context.Context, database *mongo.Database) (*TodoList, error) {
	collection := database.Collection(CollectionName)
	res, err := collection.InsertOne(context, &tl)
	if err != nil {
		log.Fatal("insert one failed")
		return &TodoList{}, err
	}
	// Type assertions: https://tour.golang.org/methods/15
	tl.ID = res.InsertedID.(primitive.ObjectID)
	return tl, nil
}

func (tl *TodoList) FindAllTodoLists(context context.Context, database *mongo.Database) (*[]TodoList, error) {
	var lists []TodoList
	collection := database.Collection(CollectionName)
	cur, err := collection.Find(context, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context) {
		var list TodoList
		err := cur.Decode(&list)
		if err != nil {
			log.Fatal(err)
		}
		lists = append(lists, list)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context)

	return &lists, nil
}
