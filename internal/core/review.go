package core

import (
	"context"
	"errors"

	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/genericrepo"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/endpoint"
	"time"
)

func (c *Core) FetchReviews(ctx context.Context, req *models.Request) (resp []endpoint.Review, meta *models.PaginationResponse, err error) {
	repo := genericrepo.New[entities.Review](c.db)
	reviews, meta, err := repo.GetAll(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	for _, r := range *reviews {
		resp = append(resp, endpoint.MakeReview(r))
	}

	return resp, meta, nil
}

func (c *Core) StartReview(ctx context.Context, req *endpoint.StartReviewRequest) (resp *endpoint.StartReviewResponse, err error) {
	var (
		repo    = genericrepo.New[entities.Review](c.db)
		rwrepo  = genericrepo.New[entities.ReviewWords](c.db)
		started = time.Now()
	)

	words, err := c.repo.Fetch(req.FromProficiency, req.ToProficiency, 1000, 0)
	if err != nil {
		return
	}

	if len(words) == 0 {
		return nil, errors.New("no word existed in this range")
	}

	if req.Shuffle {
		shuffle(words)
	}

	review := entities.Review{
		StartedAt:       started,
		FromProficiency: req.FromProficiency,
		ToProficiency:   req.ToProficiency,
		Total:           len(words),
		Status:          entities.Started,
	}

	if err = repo.Add(&review, ctx); err != nil {
		return
	}

	for idx, word := range words {
		reviewWord := entities.ReviewWords{
			ReviewID: review.ID,
			WordID:   word.ID,
			Status:   entities.RWNone,
			Index:    idx,
		}

		if idx == 0 {
			reviewWord.Status = entities.RWWaiting
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
	}

	total := int64(len(resp.Words))

	if current := resp.CurrentWord(); current != nil {
		resp.CurrentWordIndex = current.Index
		meta = models.MakePaginationResponse(total, int64(resp.CurrentWordIndex+1), 1)
	} else {
		meta = models.MakePaginationResponse(total, total, 1)
	}

	return
}

func (c *Core) NextWord(ctx context.Context, req *endpoint.NextWordRequest) (resp *endpoint.ReviewResponse, meta *models.PaginationResponse, err error) {
	if err = c.repo.SubmitWord(ctx, req); err != nil {
		return
	}

	resp, meta, err = c.GetReview(ctx, req.ReviewID)
	if err != nil {
		return
	}

	return
}
