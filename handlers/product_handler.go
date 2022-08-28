package handler

import (
	"net/http"

	"github.com/acxe21/crud-app.git/models"
	service "github.com/acxe21/crud-app.git/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService service.TestProductService
}

func ProvideProductHandle(p service.TestProductService) ProductHandler {
	return ProductHandler{ProductService: p}
}

func (p *ProductHandler) FindAllProduct(c *gin.Context) {

	products := p.ProductService.FindAllProduct()

	c.JSON(200, gin.H{
		"products": products,
	})
}

func (p *ProductHandler) CreateProduct(c *gin.Context) {

	var products models.Product

	err := c.ShouldBindJSON(&products)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error()})
		return
	}

	createdProduct := p.ProductService.CreateProduct(products)

	c.JSON(http.StatusOK, gin.H{
		"products": createdProduct,
	})

}

func (p *ProductHandler) FindProductById(c *gin.Context) {

	id := c.Param("id")

	product := p.ProductService.FindProductById(id)

	c.JSON(http.StatusOK, gin.H{
		"products": product,
	})

}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {

	id := c.Param("id")

	var product models.Product

	c.ShouldBindJSON(&product)

	updateProduct := p.ProductService.UpdateProduct(id, product)

	c.JSON(http.StatusOK, gin.H{
		"products": updateProduct,
	})
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	p.ProductService.DeleteProduct(id)

	c.JSON(http.StatusOK, gin.H{
		"products": "succesfully deleted",
	})
}
