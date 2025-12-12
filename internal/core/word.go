package core

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/genericrepo"
	"github.com/sirupsen/logrus"
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

	c.allWords = append(c.allWords, input)

	return
}

func (c *Core) AddRandomWord(ctx context.Context) (word *entities.Word, err error) {
	var (
		retry int
		input string
	)

	for retry = 0; retry < 3; retry++ {
		input, err = c.GetRandomWord(ctx)
		if err != nil {
			logrus.Warnln("can not generate random word", err.Error())
			continue
		}

		if word, err = c.AddWord(ctx, input); err != nil {
			logrus.Warnln("can not save generated random word", input, err.Error())
			continue
		}
		break
	}

	return
}

func (c *Core) GetRandomWord(ctx context.Context) (word string, err error) {
	return c.llm.GetRandomWord(ctx, c.allWords)
}

func (c *Core) ReloadWords() (err error) {
	words, err := c.repo.GetAllWords(context.Background())
	if err != nil {
		return
	}

	c.allWords = words

	return nil
}
