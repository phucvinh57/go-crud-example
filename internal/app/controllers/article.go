package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	db "github.com/phucvinh57/go-crud-example/db/sqlc"
)

type PostDTO struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ArticleCtrler struct {
	db  *db.Queries
	ctx context.Context
}

// GetPost godoc
// @Summary Create a new post
// @Description Create a new post with the input payload
// @Tags         Post
// @Accept       json
// @Produce      json
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts [post]

func NewArticleCrtler(db *db.Queries, ctx context.Context) *ArticleCtrler {
	return &ArticleCtrler{db, ctx}
}

func (pc *ArticleCtrler) GetArticles(c *gin.Context) {
	posts, err := pc.db.GetArticles(pc.ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, posts)
}
