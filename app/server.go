package app

import (
	"lexilift/internal/api"
	"lexilift/internal/config"
	"lexilift/internal/core"
	"lexilift/internal/repository"
	"lexilift/pkg/dictionary"
)

func Server(conf *config.Config) (err error) {
	repo, err := repository.New(conf)
	if err != nil {
		return err
	}

	dict := dictionary.New(conf.Debug)

	c := core.New(repo, dict, nil, conf.Debug)

	return api.Init(conf, c)
}
