package model

type Website struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

type WebsiteDTO struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

type Product struct {
	ID            string         `json:"id" bson:"id"`
	Name          string         `json:"name" bson:"name"`
	Description   string         `json:"description" bson:"description"`
	BrandID       string         `json:"brandId" bson:"brandId"`
	WebsitePrices []WebsitePrice `json:"websitePrices" bson:"websitePrices"`
}

type ProductDTO struct {
	Name          string         `json:"name" bson:"name"`
	Description   string         `json:"description" bson:"description"`
	BrandID       string         `json:"brandId" bson:"brandId"`
	WebsitePrices []WebsitePrice `json:"websitePrices" bson:"websitePrices"`
}

type WebsitePrice struct {
	WebsiteID string  `json:"websiteId" bson:"websiteId"`
	Price     float64 `json:"price" bson:"price"`
	Stock     int     `json:"stock" bson:"stock"`
}

type WebsitePriceDTO struct {
	Price float64 `json:"price" bson:"price"`
	Stock int     `json:"stock" bson:"stock"`
}

type Brand struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type BrandDTO struct {
	Name string `json:"name" bson:"name"`
}
