package endpoint

type AddWordRequest struct {
	Word string `json:"word"`
}

//
//type Word struct {
//	ID          string
//	CreatedAt   time.Time
//	UpdatedAt   time.Time
//	Word        string
//	Mean        string
//	SoundFile   string
//	Dict        *entities.Dictionary
//	Proficiency int
//	ReviewCount int
//	Score       int
//	Tags        []*Tag
//}
//
//type Tag struct{}
