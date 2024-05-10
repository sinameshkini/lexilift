package dictionary

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"lexilift/internal/models"
	"net/http"
	"os"
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

func (d *API) Find(word string) (dict *models.Dictionary, sound string, err error) {

	var dictResp []*models.Dictionary

	if _, err = d.rc.R().SetResult(&dictResp).Get(fmt.Sprintf("/%s", word)); err != nil {
		return
	}

	if len(dictResp) != 0 {
		dict = dictResp[0]
		//if len(dict.Phonetics) != 0 {
		//	link := dict.Phonetics[0].Audio
		//	fmt.Printf("downloading sound file from: %s\n", link)
		//	sound = fmt.Sprintf("audios/%s.mp3", word)
		//	if err = downloadFile(link, sound); err != nil {
		//		slog.Error(err.Error())
		//		sound = ""
		//	}
		//}
	}

	return
}

func downloadFile(url, filepath string) error {
	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
