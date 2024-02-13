package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/phucvinh57/go-crud-example/configs"
	sqlc "github.com/phucvinh57/go-crud-example/db/sqlc"
	"github.com/phucvinh57/go-crud-example/internal/app/routes"

	_ "github.com/lib/pq"
	dbPkg "github.com/phucvinh57/go-crud-example/db"
)

// @contact.name   Nguyen Phuc Vinh
// @contact.email  npvinh0507@gmail.com
func main() {
	ctx := context.Background()
	app := gin.Default()

	router := app.Group("/api")
	db := sqlc.New(dbPkg.Init())

	routes.ArticleRoute(router, db, ctx)

	configs.InitSwagger(app)
	app.Run("localhost:8080")
}
