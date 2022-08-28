package service

import (
	"github.com/acxe21/crud-app.git/models"
	"github.com/acxe21/crud-app.git/repository"
)

type TestProductService interface {
	CreateProduct(product models.Product) models.Product
	FindAllProduct() []models.Product
	FindProductById(id string) []models.Product
	UpdateProduct(id string, product models.Product) models.Product
	DeleteProduct(id string)
}
type ProductServiceImpl struct {
	ProductRepository repository.TestProductRepository
}

func ProvideProductService(pr repository.TestProductRepository) TestProductService {
	return &ProductServiceImpl{ProductRepository: pr}
}

func (p *ProductServiceImpl) FindAllProduct() []models.Product {
	return p.ProductRepository.FindAllProduct()
}

func (p *ProductServiceImpl) CreateProduct(product models.Product) models.Product {

	p.ProductRepository.CreateProduct(product)
	return product

}

func (p *ProductServiceImpl) FindProductById(id string) []models.Product {

	return p.ProductRepository.FindProductById(id)
}

func (p *ProductServiceImpl) UpdateProduct(id string, product models.Product) models.Product {

	return p.ProductRepository.UpdateProduct(id, product)
}

func (p *ProductServiceImpl) DeleteProduct(id string) {

	p.ProductRepository.DeleteProduct(id)
}
