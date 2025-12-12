package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/sinameshkini/microkit/models"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/endpoint"
	"lexilift/pkg/scoring"
	"time"
)

func (r *Repo) GetReview(ctx context.Context, id models.IID) (review *entities.Review, err error) {
	if err = r.db.WithContext(ctx).
		Preload("Words.Word").
		First(&review, id).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *Repo) SubmitWord(ctx context.Context, req *endpoint.NextWordRequest) (err error) {
	var (
		next entities.ReviewWords
		now  = time.Now().Local()
		// 1. Explicitly start the transaction
		tx = r.db.WithContext(ctx).Begin()
	)

	// Ensure we always clean up. If tx is committed, Rollback does nothing.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	review, err := r.GetReview(ctx, req.ReviewID)
	if err != nil {
		tx.Rollback()
		return err // Must return error
	}

	current := review.CurrentWord()
	if current.Status != entities.RWWaiting {
		tx.Rollback()
		return errors.New("review already committed")
	}

	// Update Current Word
	current.Status = req.Status
	current.FinishedAt = &now
	if err = tx.Save(&current).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Try to find the next word
	err = tx.Where("review_id = ? AND index = ?", current.ReviewID, current.Index+1).First(&next).Error

	if err != nil {
		// --- SCENARIO: Review Finished (No more words) ---

		// Reload review with all words and actual word details
		if err = tx.Preload("Words.Word").First(&review, review.ID).Error; err != nil {
			tx.Rollback()
			return err
		}

		for i := range review.Words {
			word := review.Words[i]

			switch word.Status {
			case entities.RWNone:
				tx.Rollback()
				return fmt.Errorf("word not submitted: %s", word.Word.Word)
			case entities.RWKnown:
				review.Know++
				word.Word.Proficiency++
				score := scoring.CalculateScore(word.Duration(), word.Word.Proficiency, word.Word.ReviewCount)
				word.Word.Score += score
				review.Score += score
			case entities.RWUnknown:
				review.NotKnow++
				word.Word.Proficiency--
			case entities.RWSkipped:
				// No score changes
			}

			word.Word.ReviewCount++
			// 2. IMPORTANT: Use tx here, not r.Update(), to stay in transaction
			if err = tx.Save(&word.Word).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		// Finalize Review Metadata
		updateData := map[string]interface{}{
			"duration":      now.Sub(review.StartedAt),
			"current_index": -1,
			"status":        entities.Finished,
			"know":          review.Know,
			"not_know":      review.NotKnow,
			"score":         review.Score,
		}

		if err = tx.Model(&entities.Review{}).Where("id = ?", review.ID).Updates(updateData).Error; err != nil {
			tx.Rollback()
			return err
		}

	} else {
		// --- SCENARIO: Next Word Exists ---

		next.StartedAt = &now
		next.Status = entities.RWWaiting
		if err = tx.Save(&next).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err = tx.Model(&entities.Review{}).Where("id = ?", review.ID).Update("current_index", next.Index).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 3. Success! Commit the transaction
	return tx.Commit().Error
}

//func (r *Repo) SubmitWord(ctx context.Context, req *endpoint.NextWordRequest) (err error) {
//	var (
//		next entities.ReviewWords
//		tx   = r.db.WithContext(ctx)
//		now  = time.Now().Local()
//	)
//
//	review, err := r.GetReview(ctx, req.ReviewID)
//	if err != nil {
//		tx.Rollback()
//		return
//	}
//
//	current := review.CurrentWord()
//
//	if current.Status != entities.RWWaiting {
//		tx.Rollback()
//		return errors.New("review already commited")
//	}
//
//	current.Status = req.Status
//
//	// Set Results
//	current.FinishedAt = &now
//
//	if err = tx.Updates(&current).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	// Get next review word
//	if err = tx.
//		Where("review_id = ?", current.ReviewID).
//		Where("index = ?", current.Index+1).
//		First(&next).Error; err != nil {
//
//		// review finished
//		// Get review again
//		err = tx.Preload("Words.Word").
//			First(&review, review.ID).Error
//		if err != nil {
//			tx.Rollback()
//			return
//		}
//
//		for _, word := range review.Words {
//			switch word.Status {
//			case entities.RWNone:
//				tx.Rollback()
//				return fmt.Errorf("no submitted word exist: %s", word.Word.Word)
//			case entities.RWSkipped:
//			case entities.RWKnown:
//				review.Know++
//				word.Word.Proficiency++
//				score := scoring.CalculateScore(word.Duration(), word.Word.Proficiency, word.Word.ReviewCount)
//				word.Word.Score += score
//				review.Score += score
//			case entities.RWUnknown:
//				review.NotKnow++
//				word.Word.Proficiency--
//			default:
//			}
//
//			word.Word.ReviewCount++
//			if err = r.Update(word.Word); err != nil {
//				logrus.Errorln("can not update word", err.Error())
//			}
//		}
//
//		if err = tx.Model(&entities.Review{}).
//			Where("id = ?", review.ID).
//			Update("duration", now.Sub(review.StartedAt)).
//			Update("current_index", -1).
//			Update("status", entities.Finished).
//			Update("know", review.Know).
//			Update("not_know", review.NotKnow).
//			Update("score", review.Score).
//			Error; err != nil {
//			tx.Rollback()
//			return err
//		}
//
//	} else {
//		//  Next existed
//
//		//	Update next
//		next.StartedAt = &now
//		next.Status = entities.RWWaiting
//		if err = tx.Updates(&next).Error; err != nil {
//			tx.Rollback()
//			return err
//		}
//
//		// Update review.current_index
//		if err = tx.Model(&entities.Review{}).
//			Where("id = ?", review.ID).
//			Update("current_index", next.Index).
//			Error; err != nil {
//			tx.Rollback()
//			return err
//		}
//	}
//
//	return tx.Commit().Error
//}
