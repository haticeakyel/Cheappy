package main

import (
	"strings"

	"github.com/google/uuid"
	model "github.com/haticeakyel/back-end/models"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) Service {
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
		ID:           GenerateUUID(8),
		Name:         productDTO.Name,
		Description:  productDTO.Description,
		BrandID:      productDTO.BrandID,
		CategoryID:   productDTO.CategoryID,
		ProductImage: productDTO.ProductImage,
		WebsitePrices: []model.WebsitePrice{
			{WebsiteID: productDTO.WebsitePrices[0].WebsiteID, Price: productDTO.WebsitePrices[0].Price, Stock: productDTO.WebsitePrices[0].Stock},
			{WebsiteID: productDTO.WebsitePrices[1].WebsiteID, Price: productDTO.WebsitePrices[1].Price, Stock: productDTO.WebsitePrices[1].Stock},
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

func (s *Service) GetProduct(ID string) (*model.Product, error) {

	gotProduct, err := s.Repository.GetProduct(ID)

	if err != nil {
		return nil, err
	}

	return gotProduct, nil
}

func (s *Service) CreateBrand(brandDTO model.BrandDTO) (*model.Brand, error) {
	brandCreate := model.Brand{
		ID:   GenerateUUID(8),
		Name: brandDTO.Name,
	}

	brandCreated, err := s.Repository.CreateBrand(brandCreate)
	if err != nil {
		return nil, err
	}

	return brandCreated, nil
}

func (s *Service) GetBrands() ([]model.Brand, error) {
	brandList, err := s.Repository.GetBrands()

	if err != nil {
		return nil, err
	}

	return brandList, nil
}

func (s *Service) GetBrand(ID string) (*model.Brand, error) {

	gotBrand, err := s.Repository.GetBrand(ID)

	if err != nil {
		return nil, err
	}

	return gotBrand, nil
}

func (s *Service) CreateWebsite(websiteDTO model.WebsiteDTO) (*model.Website, error) {
	websiteCreate := model.Website{
		ID:   GenerateUUID(8),
		Name: websiteDTO.Name,
		URL:  websiteDTO.URL,
	}

	websiteCreated, err := s.Repository.CreateWebsite(websiteCreate)
	if err != nil {
		return nil, err
	}

	return websiteCreated, nil
}

func (s *Service) GetWebsites() ([]model.Website, error) {
	websiteList, err := s.Repository.GetWebsites()

	if err != nil {
		return nil, err
	}

	return websiteList, nil
}

func (s *Service) GetWebsite(ID string) (*model.Website, error) {

	gotWeb, err := s.Repository.GetWebsite(ID)

	if err != nil {
		return nil, err
	}

	return gotWeb, nil
}

func (s *Service) CreateCategory(categoryDTO model.CategoryDTO) (*model.Category, error) {
	categoryCreate := model.Category{
		ID:   GenerateUUID(8),
		Name: categoryDTO.Name,
	}

	categoryCreated, err := s.Repository.CreateCategory(categoryCreate)
	if err != nil {
		return nil, err
	}

	return categoryCreated, nil
}

func (s *Service) GetCategories() ([]model.Category, error) {
	categoryList, err := s.Repository.GetCategories()

	if err != nil {
		return nil, err
	}

	return categoryList, nil
}

func (s *Service) GetCategory(ID string) (*model.Category, error) {

	gotCategory, err := s.Repository.GetCategory(ID)

	if err != nil {
		return nil, err
	}

	return gotCategory, nil
}

func (s *Service) DeleteProduct(ID string) error {

	err := s.Repository.DeleteProduct(ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteBrand(ID string) error {

	err := s.Repository.DeleteBrand(ID)

	if err != nil {
		return err
	}

	return nil
}
