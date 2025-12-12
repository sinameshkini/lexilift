package app

import (
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"lexilift/internal/core"
	"lexilift/internal/repository"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/dictionary"
	"lexilift/pkg/ollama"
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
	)

	db, err := database.NewDBWithDsn("host=localhost user=admin password=admin dbname=lexilift port=5432 sslmode=disable", debug)
	//db, err := database.NewSQLite("gorm.db", debug)
	if err != nil {
		return err
	}

	if err = database.Migrate(db.DB(), entities.AllTables); err != nil {
		slog.Error(err.Error())
	}

	if repo, err = repository.New(db.DB()); err != nil {
		return err
	}

	dict = dictionary.New(debug)

	if ply, err = player.New(); err != nil {
		return err
	}

	llm := ollama.New("http://localhost:11434/api", "gemma3", true)

	c = core.New(db.DB(), repo, dict, ply, llm, debug)

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
