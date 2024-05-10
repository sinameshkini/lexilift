package app

import (
	"lexilift/internal/core"
	"lexilift/internal/repository"
	"lexilift/pkg/dictionary"
	"log/slog"
)

func Run(debug bool) error {
	var (
		err  error
		repo *repository.Repo
		dict *dictionary.API
		//ply  *player.Player
		c *core.Core
		//word  *models.Word
		//repo  *repository.Repo
	)

	if repo, err = repository.New(debug); err != nil {
		return err
	}

	dict = dictionary.New(debug)
	//
	//if ply, err = player.New(); err != nil {
	//	return err
	//}

	//c = core.New(repo, dict, ply, debug)
	c = core.New(repo, dict, debug)

	if err = c.About(); err != nil {
		return err
	}

	if err = c.Menu(); err != nil {
		return err
	}

	for {
		if err = c.Handler(); err != nil {
			slog.Error(err.Error())
		}
	}
}
