package repository

import (
	"fmt"
	"log"

	"github.com/acxe21/crud-app.git/models"
	"gorm.io/gorm"
)

type TestProductRepository interface {
	CreateProduct(product models.Product) models.Product
	FindAllProduct() []models.Product
	FindProductById(id string) []models.Product
	UpdateProduct(id string, product models.Product) models.Product
	DeleteProduct(id string)
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func ProvideProductRepository(db *gorm.DB) TestProductRepository {
	return &ProductRepositoryImpl{Db: db}
}

func (p *ProductRepositoryImpl) FindAllProduct() []models.Product {

	var products []models.Product

	p.Db.Find(&products)

	return products

}

func (p *ProductRepositoryImpl) CreateProduct(product models.Product) models.Product {

	err := p.Db.Create(&product)

	if err != nil {

		log.Println("error")
	}

	return product

}

func (p *ProductRepositoryImpl) FindProductById(id string) []models.Product {

	var products []models.Product

	p.Db.First(&products, id)

	return products
}

func (p *ProductRepositoryImpl) UpdateProduct(id string, product models.Product) models.Product {

	test := p.Db.Where("id = ?", id).First(&product)
	fmt.Println(test)
	p.Db.Save(&product)

	return product
}

func (p *ProductRepositoryImpl) DeleteProduct(id string) {

	var products []models.Product

	p.Db.Delete(&products, id)

}
