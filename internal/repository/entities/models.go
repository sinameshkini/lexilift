package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
	"time"
)

var AllTables = []interface{}{
	Word{},
	Review{},
	ReviewWords{},
	Tag{},
}

type Review struct {
	ID              models.IID
	StartedAt       time.Time
	Duration        time.Duration
	FromProficiency int
	ToProficiency   int
	Total           int
	CurrentIndex    int
	Know            int
	NotKnow         int
	Score           int `gorm:"default=0"`
	Comment         string
	Status          ReviewStatus
	Words           []*ReviewWords
}

func (m Review) CurrentWord() *ReviewWords {
	for _, w := range m.Words {
		if w.Status == RWWaiting {
			return w
		}
	}

	return nil
}

type ReviewStatus string

const (
	Started  ReviewStatus = "started"
	Finished ReviewStatus = "finished"
)

type Word struct {
	models.ModelIID
	Word        string      `json:"word"`
	Mean        string      `json:"mean"`
	SoundFile   string      `json:"soundFile"`
	Dict        *Dictionary `json:"dict"`
	Proficiency int         `json:"proficiency"`
	ReviewCount int         `json:"reviewCount" gorm:"default=0"`
	Score       int         `json:"score" gorm:"default=0"`
	Tags        []*Tag      `json:"tags" gorm:"many2many:word_tag;"`
}

type ReviewWords struct {
	ID         models.IID
	CreatedAt  time.Time
	StartedAt  *time.Time
	FinishedAt *time.Time
	ReviewID   models.IID
	Review     *Review
	WordID     models.IID
	Word       *Word
	Status     ReviewWordStatus
	Index      int
}

func (m *ReviewWords) Duration() int {
	if m.StartedAt == nil || m.FinishedAt == nil {
		return 0
	}

	return int(m.FinishedAt.Sub(*m.StartedAt).Round(time.Second).Seconds())
}

type ReviewWordStatus string

const (
	RWNone    ReviewWordStatus = ""
	RWWaiting ReviewWordStatus = "waiting"
	RWKnown   ReviewWordStatus = "known"
	RWUnknown ReviewWordStatus = "unknown"
	RWSkipped ReviewWordStatus = "skipped"
)

func (m *Word) AfterFind(tx *gorm.DB) (err error) {
	if m.SoundFile != "" {
		m.SoundFile = "http://localhost:5050/" + m.SoundFile
	}

	return nil
}

type Tag struct {
	gorm.Model
	Name  string  `json:"name"`
	Words []*Word `json:"words" gorm:"many2many:word_tag;"`
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
