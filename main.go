package main

import (
	"github.com/acxe21/crud-app.git/config"
	"github.com/gin-gonic/gin"
)

func main() {

	db := config.InitDB()

	r := gin.Default()

	ProductHandler := ProductHandler(db)

	r.GET("/product", ProductHandler.FindAllProduct)
	r.POST("/product", ProductHandler.CreateProduct)
	r.GET("/product/:id", ProductHandler.FindProductById)
	r.PUT("/product/:id", ProductHandler.UpdateProduct)
	r.DELETE("/product/:id", ProductHandler.DeleteProduct)

	//	r.productRoute(r, ProductHandler)

	err := r.Run("localhost:3000")
	if err != nil {
		panic(err)
	}

}
