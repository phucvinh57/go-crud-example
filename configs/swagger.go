package configs

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"github.com/phucvinh57/go-crud-example/docs"
)

func InitSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
