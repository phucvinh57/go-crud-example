// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	GetArticles(ctx context.Context) ([]Article, error)
}

var _ Querier = (*Queries)(nil)
