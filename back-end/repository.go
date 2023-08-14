package main

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/haticeakyel/back-end/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() (*Repository, error) {
    uri := "mongodb+srv://haticeakyel:cheappytest@cluster0.tn9d3bu.mongodb.net/"
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }

    return &Repository{client}, nil
}


func (repository *Repository) GetProduct(ID string) (*model.Product, error) {
	collection := repository.client.Database("product").Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := &model.Product{}
	filters := bson.M{"id": ID}
	err := collection.FindOne(ctx, filters).Decode(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (repository Repository) CreateProduct(productData model.Product) (*model.Product, error) {
    collection := repository.client.Database("product").Collection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, productData)
    if err != nil {
        return nil, fmt.Errorf("failed to insert product: %w", err)
    }

    // Return the newly created product data directly
    return &productData, nil
}


func (repository Repository) GetProducts() ([]model.Product, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"_id": -1})

	cur, err := collection.Find(ctx, options)
	if err != nil {
		return nil, err
	}

	entities := []model.Product{}

	for cur.Next(ctx) {

		entity := model.Product{}
		err := cur.Decode(&entity)
		if err != nil {
			log.Fatal(err)
		}
		entities = append(entities, entity)
	}

	if err != nil {
		return nil, err
	}
	return entities, nil
}
