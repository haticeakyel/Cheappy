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
	collection := repository.client.Database("product").Collection("products")
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

	return &productData, nil
}

func (repository Repository) GetProducts() ([]model.Product, error) {
	collection := repository.client.Database("product").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	products := []model.Product{}
	for cur.Next(ctx) {
		var product model.Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)

	}

	return products, nil
}

func (repository Repository) CreateBrand(brandData model.Brand) (*model.Brand, error) {
	collection := repository.client.Database("product").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, brandData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert brand: %w", err)
	}

	return &brandData, nil
}

func (repository Repository) GetBrands() ([]model.Brand, error) {
	collection := repository.client.Database("product").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	brands := []model.Brand{}
	for cur.Next(ctx) {
		var brand model.Brand
		err := cur.Decode(&brand)
		if err != nil {
			log.Fatal(err)
		}

		brands = append(brands, brand)

	}

	return brands, nil
}

func (repository *Repository) GetBrand(ID string) (*model.Brand, error) {
	collection := repository.client.Database("product").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := &model.Brand{}
	filters := bson.M{"id": ID}
	err := collection.FindOne(ctx, filters).Decode(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (repository Repository) CreateWebsite(webData model.Website) (*model.Website, error) {
	collection := repository.client.Database("product").Collection("websites")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, webData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert website: %w", err)
	}

	return &webData, nil
}

func (repository Repository) GetWebsites() ([]model.Website, error) {
	collection := repository.client.Database("product").Collection("websites")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	webs := []model.Website{}
	for cur.Next(ctx) {
		var web model.Website
		err := cur.Decode(&web)
		if err != nil {
			log.Fatal(err)
		}

		webs = append(webs, web)

	}

	return webs, nil
}

func (repository *Repository) GetWebsite(ID string) (*model.Website, error) {
	collection := repository.client.Database("product").Collection("websites")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := &model.Website{}
	filters := bson.M{"id": ID}
	err := collection.FindOne(ctx, filters).Decode(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (repository Repository) CreateCategory(categoryData model.Category) (*model.Category, error) {
	collection := repository.client.Database("product").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, categoryData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert category: %w", err)
	}

	return &categoryData, nil
}

func (repository Repository) GetCategories() ([]model.Category, error) {
	collection := repository.client.Database("product").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	categories := []model.Category{}
	for cur.Next(ctx) {
		var category model.Category
		err := cur.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)

	}

	return categories, nil
}

func (repository *Repository) GetCategory(ID string) (*model.Category, error) {
	collection := repository.client.Database("product").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entity := &model.Category{}
	filters := bson.M{"id": ID}
	err := collection.FindOne(ctx, filters).Decode(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
