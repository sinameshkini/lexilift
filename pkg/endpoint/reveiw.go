package endpoint

import (
	"github.com/sinameshkini/microkit/models"
	"lexilift/internal/repository/entities"
	"time"
)

type StartReviewRequest struct {
	FromProficiency int `json:"from_proficiency"`
	ToProficiency   int `json:"to_proficiency"`
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
