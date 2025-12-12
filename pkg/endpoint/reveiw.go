package endpoint

import (
	"github.com/sinameshkini/microkit/models"
	"lexilift/internal/repository/entities"
	"time"
)

type StartReviewRequest struct {
	FromProficiency int  `json:"from_proficiency"`
	ToProficiency   int  `json:"to_proficiency"`
	Shuffle         bool `json:"shuffle"`
}

type StartReviewResponse struct {
	ID         models.IID `json:"id"`
	TotalWords int        `json:"total_words"`
	CreatedAt  time.Time  `json:"created_at"`
}

type ReviewResponse struct {
	entities.Review
	CurrentWordIndex int `json:"current_word_index"`
}

type NextWordRequest struct {
	ReviewID models.IID                `json:"review_id"`
	Index    int                       `json:"index"`
	Status   entities.ReviewWordStatus `json:"status"`
}

type Review struct {
	ID              models.IID
	StartedAt       string
	Duration        string
	FromProficiency int
	ToProficiency   int
	Total           int
	CurrentIndex    int
	Know            int
	NotKnow         int
	Score           int
	Comment         string
	Status          string
	Words           []*entities.ReviewWords
}

func MakeReview(review entities.Review) Review {
	return Review{
		ID:              review.ID,
		StartedAt:       review.StartedAt.Format(time.DateTime),
		Duration:        review.Duration.Round(time.Second).String(),
		FromProficiency: review.FromProficiency,
		ToProficiency:   review.ToProficiency,
		Total:           review.Total,
		CurrentIndex:    review.CurrentIndex,
		Know:            review.Know,
		NotKnow:         review.NotKnow,
		Score:           review.Score,
		Comment:         review.Comment,
		Status:          string(review.Status),
		Words:           review.Words,
	}
}
