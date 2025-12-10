package core

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/genericrepo"
	"lexilift/internal/repository/entities"
)

func (c *Core) FetchWords(ctx context.Context, req *models.Request) (resp []entities.Word, meta *models.PaginationResponse, err error) {
	repo := genericrepo.New[entities.Word](c.db)
	words, meta, err := repo.GetAll(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return *words, meta, nil
}

func (c *Core) AddWord(ctx context.Context, input string) (word *entities.Word, err error) {
	if word, err = c.repo.Get(input); err == nil {
		return word, models.ErrAlreadyExist
	}

	if word, err = c.Save(input); err != nil {
		return
	}

	return
}
