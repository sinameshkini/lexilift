package core

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/genericrepo"
	"lexilift/internal/repository/entities"
)

func (c *Core) FetchReviews(ctx context.Context, req *models.Request) (resp []entities.Review, meta *models.PaginationResponse, err error) {
	repo := genericrepo.New[entities.Review](c.db)
	reviews, meta, err := repo.GetAll(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return *reviews, meta, nil
}
