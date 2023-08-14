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

func (s *Service) CreateProduct(productDTO model.ProductDTO) (*model.Product, error) { // Use models.ProductDTO and models.Product
	productCreate := model.Product{ // Use models.Product
		ID:          GenerateUUID(8),
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Stock:       5,
		BrandID:     productDTO.BrandID,
		WebsitePrices: []model.WebsitePrice{ // Use models.WebsitePrice
			{WebsiteID: "trendyol", Price: 199.99, Stock: 50},
			{WebsiteID: "hepsiburada", Price: 219.99, Stock: 30},
		},
	}

	productCreated, err := s.Repository.CreateProduct(productCreate)
	if err != nil {
		return nil, err
	}

	return productCreated, nil
}
