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
	rwrepo := genericrepo.New[entities.ReviewWords](c.db)

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

	for idx, word := range words {
		reviewWord := entities.ReviewWords{
			ReviewID: review.ID,
			WordID:   word.ID,
			Status:   entities.None,
			Index:    idx,
		}

		if err = rwrepo.Add(&reviewWord, ctx); err != nil {
			return
		}
	}

	resp = &endpoint.StartReviewResponse{
		ID:         review.ID,
		TotalWords: review.Total,
	}

	return resp, nil
}

func (c *Core) GetReview(ctx context.Context, reviewID models.IID) (resp *endpoint.ReviewResponse, meta *models.PaginationResponse, err error) {
	review, err := c.repo.GetReview(ctx, reviewID)
	if err != nil {
		return
	}

	resp = &endpoint.ReviewResponse{
		Review:           *review,
		CurrentWordIndex: -1,
		//Words:
	}

	for idx, word := range review.Words {
		if word.Status == entities.None {
			resp.CurrentWordIndex = idx
			break
		}
	}

	meta = models.MakePaginationResponse(int64(len(resp.Words)), int64(resp.CurrentWordIndex+1), 1)

	return
}

func (c *Core) NextWord(ctx context.Context, req *endpoint.NextWordRequest) (resp *endpoint.ReviewResponse, meta *models.PaginationResponse, err error) {
	review, err := c.repo.GetReview(ctx, req.ReviewID)
	if err != nil {
		return
	}

	var rw *entities.ReviewWords

	for _, word := range review.Words {
		if word.Status == entities.None {
			rw = word
			break
		}
	}

	if err = c.repo.SubmitWord(ctx, rw.ID, req.Status); err != nil {
		return
	}

	return c.GetReview(ctx, req.ReviewID)
}
