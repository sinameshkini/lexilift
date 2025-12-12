package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sinameshkini/microkit/models"
	"lexilift/internal/repository/entities"
)

func (r *Repo) GetReview(ctx context.Context, id models.IID) (review *entities.Review, err error) {
	if err = r.db.WithContext(ctx).
		Preload("Words.Word").
		First(&review, id).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *Repo) SubmitWord(ctx context.Context, reviewWordID models.IID, status entities.ReviewWordStatus) (err error) {
	var (
		rw entities.ReviewWords
	)

	if err = r.db.WithContext(ctx).
		First(&rw, reviewWordID).Error; err != nil {
		return err
	}

	if rw.Status != entities.None {
		return errors.New("review already commited")
	}

	rw.Status = status

	if err = r.db.WithContext(ctx).
		Updates(&rw).Error; err != nil {
		return err
	}

	return
}
