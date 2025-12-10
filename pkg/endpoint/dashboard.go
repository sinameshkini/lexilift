package endpoint

import (
	"lexilift/internal/repository/entities"
)

type Dashboard struct {
	ScoreRanking   []*entities.Word   `json:"scoreRanking"`
	LearningChart  map[int]int        `json:"learningChart"`
	TotalWords     int                `json:"totalWords"`
	Reviews        []*entities.Review `json:"reviews"`
	TotalReviews   int                `json:"totalReviews"`
	ReviewDuration string             `json:"reviewDuration"`
	TotalScore     int                `json:"totalScore"`
}
