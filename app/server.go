package app

import (
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"lexilift/internal/api"
	"lexilift/internal/config"
	"lexilift/internal/core"
	"lexilift/internal/repository"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/dictionary"
	"log/slog"
)

func Server(conf *config.Config) (err error) {
	db, err := database.NewDBWithDsn("host=localhost user=admin password=admin dbname=lexilift port=5432 sslmode=disable", true)
	if err != nil {
		return err
	}

	if err = database.Migrate(db.DB(), entities.AllTables); err != nil {
		slog.Error(err.Error())
	}

	repo, err := repository.New(db.DB())
	if err != nil {
		return err
	}

	dict := dictionary.New(conf.Debug)

	c := core.New(db.DB(), repo, dict, nil, conf.Debug)

	return api.Init(conf, c)
}
