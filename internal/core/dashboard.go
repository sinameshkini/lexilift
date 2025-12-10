package core

import (
	"fmt"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/endpoint"
	"log/slog"
	"sort"
	"time"
)

func (c *Core) UserDashboard() (resp *endpoint.Dashboard, err error) {
	var (
		allWords      []*entities.Word
		allReviews    []*entities.Review
		knowMap       = make(map[int]int)
		sorted        []int
		totalDuration time.Duration
		totalScore    int
	)

	resp = &endpoint.Dashboard{
		LearningChart: make(map[int]int),
	}

	if allWords, err = c.repo.GetAll(); err != nil {
		return
	}

	for idx, w := range allWords {
		knowMap[w.Proficiency] += 1

		if idx < 5 {
			resp.ScoreRanking = append(resp.ScoreRanking, w)
		}

		if idx >= len(allWords)-5 && idx >= 5 {
			resp.ScoreRanking = append(resp.ScoreRanking, w)
		}
	}

	for k := range knowMap {
		sorted = append(sorted, k)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	fmt.Println("\nMy Words:")
	for _, i := range sorted {
		resp.LearningChart[i] = knowMap[i]
	}

	resp.TotalWords = len(allWords)

	if allReviews, err = c.repo.GetAllReviews(); err != nil {
		return
	}

	for idx, r := range allReviews {
		if idx < 5 {
			if err = c.ShowReview(len(allReviews)-idx, r); err != nil {
				slog.Error(err.Error())
				continue
			}
		}

		totalDuration += r.Duration
		totalScore += r.Score
	}

	resp.TotalReviews = len(allReviews)
	resp.ReviewDuration = totalDuration.Round(time.Second).String()
	resp.TotalScore = totalScore

	return
}
