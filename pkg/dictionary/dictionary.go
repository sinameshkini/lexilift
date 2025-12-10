package dictionary

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"lexilift/internal/repository/entities"
	"lexilift/pkg/utils"
	"log/slog"
)

type API struct {
	rc *resty.Client
}

func New(debug bool) *API {
	return &API{
		rc: resty.New().
			SetBaseURL("https://api.dictionaryapi.dev/api/v2/entries/en").
			SetDebug(debug),
	}
}

func (d *API) Find(word string) (dict *entities.Dictionary, sound string, err error) {

	var dictResp []*entities.Dictionary

	if _, err = d.rc.R().SetResult(&dictResp).Get(fmt.Sprintf("/%s", word)); err != nil {
		return
	}

	if len(dictResp) != 0 {
		dict = dictResp[0]
		if len(dict.Phonetics) != 0 {
			link := dict.Phonetics[0].Audio
			if link != "" {
				fmt.Printf("downloading sound file from: %s\n", link)
				sound = fmt.Sprintf("audios/%s.mp3", word)
				if err = utils.DownloadFile(link, sound); err != nil {
					slog.Error(err.Error())
					sound = ""
					err = nil
				}
			}
		}
	}

	return
}
