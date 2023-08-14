package main

import (
	"strings"

	"github.com/google/uuid"
	"github.com/haticeakyel/back-end/models"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) Service { // Renamed parameter for clarity
	return Service{
		Repository: repository,
	}
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}

func (s *Service) CreateProduct(productDTO model.ProductDTO) (*model.Product, error) { 
	productCreate := model.Product{ 
		ID:          GenerateUUID(8),
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Stock:       5,
		BrandID:     productDTO.BrandID,
		WebsitePrices: []model.WebsitePrice{
			{WebsiteID: GenerateUUID(8), Price: 199.99, Stock: 50},
			{WebsiteID: GenerateUUID(8), Price: 219.99, Stock: 30},
		},
	}

	productCreated, err := s.Repository.CreateProduct(productCreate)
	if err != nil {
		return nil, err
	}

	return productCreated, nil
}

func (s *Service) GetProducts() ([]model.Product, error) {
	productList, err := s.Repository.GetProducts()

	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (s *Service) GetProduct( ID string) (*model.Product, error) {

	gotProduct, err := s.Repository.GetProduct( ID)

	if err != nil {
		return nil, err
	}

	return gotProduct, nil
}