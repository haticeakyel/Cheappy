package main

import (
	"context"
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

func NewRepository() *Repository {
	uri := "mongodb+srv://haticeakyel:cheappytest@cluster0.tn9d3bu.mongodb.net/"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
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
		return nil, err
	}

	product, err := repository.GetProduct(productData.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
