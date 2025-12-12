package repository

import (
	"context"
	"lexilift/internal/repository/entities"
)

func (r *Repo) GetAllWords(ctx context.Context) (words []string, err error) {
	if err = r.db.WithContext(ctx).
		Model(&entities.Word{}).
		Select("Word").
		Scan(&words).Error; err != nil {
		return nil, err
	}
	return
}
