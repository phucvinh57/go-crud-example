package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	sqlc "github.com/phucvinh57/go-crud-example/db/sqlc"
	"github.com/phucvinh57/go-crud-example/internal/app/controllers"
)

func ArticleRoute(rg *gin.RouterGroup, db *sqlc.Queries, ctx context.Context) {
	router := rg.Group("articles")
	postCtrler := controllers.NewArticleCrtler(db, ctx)
	router.GET("", postCtrler.GetArticles)
}
