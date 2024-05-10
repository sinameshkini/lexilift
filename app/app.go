package app

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"lexilift/internal/core"
	"lexilift/internal/repository"
	"lexilift/pkg/dictionary"
	"lexilift/pkg/player"
	"log/slog"
)

func Run(debug bool) error {
	var (
		err  error
		repo *repository.Repo
		dict *dictionary.API
		ply  *player.Player
		c    *core.Core
		//word  *models.Word
		//repo  *repository.Repo
	)

	if repo, err = repository.New(debug); err != nil {
		return err
	}

	dict = dictionary.New(debug)

	if ply, err = player.New(); err != nil {
		return err
	}

	c = core.New(repo, dict, ply, debug)

	banner := figure.NewFigure("LexiLift", "", true).String()
	fmt.Println(banner)

	if err = c.Dashboard(); err != nil {
		return err
	}

	for {
		if err = c.Menu(); err != nil {
			slog.Error(err.Error())
		}
	}
}
