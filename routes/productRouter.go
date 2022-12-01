package routes

import (
	"github.com/gin-gonic/gin"
	ProductController "github.com/lordnr/learn-redis/controllers/product"
)

func ProductRouter(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/products", ProductController.Products)

}
