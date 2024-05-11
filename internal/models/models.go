package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	StartedAt       time.Time
	Duration        time.Duration
	FromProficiency int
	ToProficiency   int
	Total           int
	Know            int
	NotKnow         int
	Score           int `gorm:"default=0"`
}

type Word struct {
	gorm.Model
	Word        string
	Mean        string
	SoundFile   string
	Dict        *Dictionary
	Proficiency int
	ReviewCount int `gorm:"default=0"`
	Score       int `gorm:"default=0"`
}

type Dictionary struct {
	Word       string     `json:"word"`
	Phonetics  []Phonetic `json:"phonetics"`
	Meanings   []Meaning  `json:"meanings"`
	License    License    `json:"license"`
	SourceUrls []string   `json:"sourceUrls"`
}

type Phonetic struct {
	Text      string `json:"text"`
	Audio     string `json:"audio"`
	SourceURL string `json:"sourceUrl"`
	License   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    string   `json:"example"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (s *Dictionary) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Dictionary{}
	err := json.Unmarshal(bytes, &result)
	*s = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (s Dictionary) Value() (driver.Value, error) {
	//if len(s) == 0 {
	//	return nil, nil
	//}
	return json.Marshal(s)
}
