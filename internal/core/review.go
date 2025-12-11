package core

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/genericrepo"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/endpoint"
)

func (c *Core) FetchReviews(ctx context.Context, req *models.Request) (resp []entities.Review, meta *models.PaginationResponse, err error) {
	repo := genericrepo.New[entities.Review](c.db)
	reviews, meta, err := repo.GetAll(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return *reviews, meta, nil
}

func (c *Core) StartReview(ctx context.Context, req *endpoint.StartReviewRequest) (resp *endpoint.StartReviewResponse, err error) {
	repo := genericrepo.New[entities.Review](c.db)

	words, err := c.repo.Fetch(req.FromProficiency, req.ToProficiency, 1000, 0)
	if err != nil {
		return
	}

	review := entities.Review{
		FromProficiency: req.FromProficiency,
		ToProficiency:   req.ToProficiency,
		Total:           len(words),
	}

	if err = repo.Add(&review, ctx); err != nil {
		return
	}

	resp = &endpoint.StartReviewResponse{
		ID:         review.ID,
		TotalWords: review.Total,
	}

	return resp, nil
}
