package endpoint

import (
	"github.com/sinameshkini/microkit/models"
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
