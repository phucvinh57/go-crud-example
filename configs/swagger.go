package configs

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	apidocs "github.com/phucvinh57/go-crud-example/api"
)

func InitSwagger(r *gin.Engine) {
	apidocs.SwaggerInfo.Title = "API Documentation"
	apidocs.SwaggerInfo.Version = "1.0"
	apidocs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
