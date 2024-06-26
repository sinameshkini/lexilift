package core

import (
	"bufio"
	"errors"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"github.com/common-nighthawk/go-figure"
	"lexilift/internal/models"
	"lexilift/internal/repository"
	"lexilift/pkg/dictionary"
	"log/slog"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"
)

type Core struct {
	repo *repository.Repo
	dict *dictionary.API
	//ply   *player.Player
	level int
	debug bool
}

// func New(repo *repository.Repo, dict *dictionary.API, ply *player.Player, debug bool) *Core {

func New(repo *repository.Repo, dict *dictionary.API, debug bool) *Core {
	return &Core{
		repo: repo,
		dict: dict,
		//ply:   ply,
		debug: debug,
		level: 3,
	}
}

func (c *Core) Handler() (err error) {
	var (
		input rune
	)

	fmt.Print(":::> ")

	input, err = inputChar()
	if err != nil {
		return err
	}

	if c.debug {
		fmt.Printf("Input: %c\n", input)
	}

	switch input {
	case '1':
		return c.Review()
	case '2':
		return c.AddNewWord()
	case '3':
		return c.AddWordsList()
	case '4':
		return c.ReviewHistory()
	case '5':
		return c.NewTag()
	case '6':
		return c.Tags()
	case 'm':
		return c.Menu()
	case 'd':
		return c.Dashboard()
	case 'c':
		return clearConsole()
	case 'a':
		return c.About()
	case 'q':
		fmt.Println("See you soon, goodbye :)")
		os.Exit(0)
	default:
		return
	}

	return
}

func (c *Core) About() (err error) {
	banner := figure.NewFigure("LexiLift", "", true).String()
	fmt.Println(banner)
	fmt.Println("LexiLift is a free and open-source CLI app designed to help you learn any English word you want!")
	fmt.Println("")
	return nil
}

func (c *Core) Menu() (err error) {
	fmt.Println(strings.Repeat(">", 32), "LexiLift", strings.Repeat("<", 32))
	fmt.Println("Menu:")
	fmt.Println("\t1- Review my words")
	fmt.Println("\t2- Add a new word to my words")
	fmt.Println("\t3- Add words list to my words")
	fmt.Println("\t4- Review history")
	fmt.Println("\t5- Add a new tag")
	fmt.Println("\t6- Tags list")
	fmt.Println("\tm- Menu")
	fmt.Println("\td- Dashboard")
	fmt.Println("\tc- Clear")
	fmt.Println("\tq- close the app")
	fmt.Println("Press the character corresponding to the action you want to perform")
	return
}

func (c *Core) Dashboard() (err error) {
	var (
		allWords      []*models.Word
		allReviews    []*models.Review
		knowMap       = make(map[int]int)
		sorted        []int
		totalDuration time.Duration
		totalScore    int
	)

	fmt.Println(strings.Repeat(">", 32), "Dashboard", strings.Repeat("<", 32))

	if allWords, err = c.repo.GetAll(); err != nil {
		return err
	}

	fmt.Println("Score Ranking:")
	for idx, w := range allWords {
		knowMap[w.Proficiency] += 1

		if idx < 5 {
			fmt.Printf("\t%d- %s\t(Score: %d, Proficiency: %d)\n", idx+1, w.Word, w.Score, w.Proficiency)
		}

		if idx == 5 {
			fmt.Println("\t.")
		}

		if idx >= len(allWords)-5 && idx >= 5 {
			fmt.Printf("\t%d- %s\t(Score: %d, Proficiency: %d)\n", idx+1, w.Word, w.Score, w.Proficiency)
		}
	}

	for k := range knowMap {
		sorted = append(sorted, k)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	fmt.Println("\nMy Words:")
	for _, i := range sorted {
		kn := knowMap[i]
		fmt.Printf("\tProficiency: %d\t Count: %d\n", i, kn)
	}

	fmt.Printf("Total: %d\n", len(allWords))

	if allReviews, err = c.repo.GetAllReviews(); err != nil {
		return err
	}

	fmt.Println("\nMy Reviews:")
	for idx, r := range allReviews {
		if idx < 5 {
			if err = c.ShowReview(len(allReviews)-idx, r); err != nil {
				slog.Error(err.Error())
				continue
			}
		}

		totalDuration += r.Duration
		totalScore += r.Score
	}
	fmt.Printf("Total: %d, Duration: %s, Score: %d\n",
		len(allReviews), totalDuration.Round(time.Second).String(), totalScore)
	fmt.Println("")

	return nil
}

func (c *Core) Save(word string) (w *models.Word, err error) {
	var (
		mean  string
		dict  *models.Dictionary
		sound string
	)

	fmt.Println("Getting Persian translation from the Google Translate API...")
	if mean, err = gt.Translate(word, "en", "fa"); err != nil {
		return
	}

	fmt.Println("Getting translation from the dictionary API...")
	if dict, sound, err = c.dict.Find(word); err != nil {
		return
	}

	w = &models.Word{
		Word:      word,
		Mean:      mean,
		Dict:      dict,
		SoundFile: sound,
	}

	if err = c.repo.Create(*w); err != nil {
		return
	}

	return
}

func shuffle(array []*models.Word) {
	for i := len(array) - 1; i > 0; i-- { //run the loop from the end till the start
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i] //swap the random element with the current element
	}
}

func (c *Core) Review() (err error) {
	var (
		words          []*models.Word
		input          rune
		fromKnw, toKnw int
		startedAt      = time.Now().Local()
		know, notKnow  int
		totalScore     int
		counter        int
	)

	// TODO review history

	fmt.Print("From proficiency count: ")
	if fromKnw, err = inputInt(); err != nil {
		return err
	}

	fmt.Print("To proficiency count: ")
	if toKnw, err = inputInt(); err != nil {
		return err
	}

	if words, err = c.repo.Fetch(fromKnw, toKnw, 1000, 0); err != nil {
		return err
	}

	shuffle(words)

	printDiv()

	fmt.Printf("%d words found\n", len(words))
	if len(words) == 0 {
		return nil
	}

	fmt.Println("Starting ...")
	fmt.Println("Press Enter to view word meaning")

	for idx, word := range words {
		counter += 1
	review:
		started := time.Now()
		if err = c.ShowWord(idx, word); err != nil {
			slog.Error(err.Error())
		}

		fmt.Println("(1) I know, (2) I don't know, (0) nothing, (a) add a new word, (q) stop")
		if input, err = inputChar(); err != nil {
			return err
		}

		dur := int(time.Now().Sub(started).Round(time.Second).Seconds())

		if input == 'q' {
			break
		}

		switch input {
		case '1':
			word.Proficiency += 1
			word.ReviewCount += 1
			score := c.CalculateScore(dur, word.Proficiency, word.ReviewCount)
			fmt.Printf("Score: %d\n", score)
			totalScore += score
			word.Score += score
			know += 1
		case '2':
			word.Proficiency -= 1
			word.ReviewCount += 1
			//score := CalculateScore(c.level, dur, word.Proficiency, word.ReviewCount, -1)
			//fmt.Printf("Score: %d", score)
			//totalScore += score
			notKnow += 1
		case 'a':
			if err = c.AddNewWord(); err != nil {
				slog.Error(err.Error())
			}

			goto review
		default:
			continue
		}

		if err = c.repo.Update(word); err != nil {
			return err
		}
	}
	fmt.Println("**")
	fmt.Println("****")
	fmt.Println("****** Congratulations, review completed")
	fmt.Println("****")
	fmt.Println("**")
	fmt.Print("Please enter a comment: ")
	comment, err := inputString()
	if err != nil {
		slog.Error(err.Error())
	}

	review := models.Review{
		StartedAt:       startedAt,
		Duration:        time.Now().Sub(startedAt),
		FromProficiency: fromKnw,
		ToProficiency:   toKnw,
		Total:           counter,
		Know:            know,
		NotKnow:         notKnow,
		Score:           totalScore,
		Comment:         comment,
	}

	if err = c.repo.CreateReview(review); err != nil {
		return
	}

	return
}

func (c *Core) ShowReview(idx int, review *models.Review) (err error) {
	fmt.Printf("%d- %s  %s\tFP:%d\tTP:%d\tCNT:%d\tKNW:%d\tNK:%d\tS:%d\tCM:%s\n",
		idx,
		review.StartedAt.Format("2006-01-02 15:04"),
		review.Duration.Round(time.Second).String(),
		review.FromProficiency,
		review.ToProficiency,
		review.Total,
		review.Know,
		review.NotKnow,
		review.Score,
		review.Comment,
	)
	return nil
}

func (c *Core) ShowWord(idx int, word *models.Word) (err error) {
	printDiv()
	fmt.Printf("\t%d- %s\n",
		idx+1, word.Word)

	if word.Dict != nil && len(word.Dict.Phonetics) != 0 {
		fmt.Printf("Phonetic: %s, ", word.Dict.Phonetics[0].Text)
	}

	fmt.Printf("Created at: %s, Reviewed: %d, Proficiency: %d)\n",
		word.CreatedAt.Format("2006-01-02 15:04"), word.ReviewCount, word.Proficiency)
	//if word.SoundFile != "" {
	//	if err = c.ply.Play(word.SoundFile); err != nil {
	//		slog.Error(err.Error())
	//	}
	//}
	_, err = inputChar()
	if err != nil {
		return err
	}

	fmt.Println("\tMeanings:")
	if word.Dict != nil && len(word.Dict.Meanings) != 0 {
		for _, meaning := range word.Dict.Meanings {
			fmt.Printf("\t(%s)\n", meaning.PartOfSpeech)
			for _, m := range meaning.Definitions {
				fmt.Printf("\t- %s\n", m.Definition)
				if m.Example != "" {
					fmt.Printf("\tExample: %s\n", m.Example)
				}
			}
			if len(meaning.Synonyms) != 0 {
				fmt.Println("\tSynonyms:")
				for _, s := range meaning.Synonyms {
					fmt.Printf("\t\t- %s\n", s)
				}
				fmt.Println("")
			}
			if len(meaning.Antonyms) != 0 {
				fmt.Println("Antonyms:")
				for _, s := range meaning.Antonyms {
					fmt.Printf("- %s", s)
				}
				fmt.Println("")
			}
		}
	}
	fmt.Printf("\tPersian: %s\n", word.Mean)

	return nil
}

func (c *Core) AddNewWord() (err error) {
	var (
		input string
		word  *models.Word
	)

	printDiv()
	fmt.Print("Type the word: ")
	if input, err = inputString(); err != nil {
		return err
	}

	if word, err = c.repo.Get(input); err == nil {
		return errors.New("already exist")
	}

	if word, err = c.Save(input); err != nil {
		return err
	}

	fmt.Printf("*** %s added successfuly!\n", strings.ToUpper(word.Word))

	if err = c.ShowWord(0, word); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func (c *Core) AddWordsList() (err error) {
	var (
		input []string
		word  *models.Word
	)

	printDiv()
	fmt.Print("type each word on a new line, submit empty line to finish: \n")
	if input, err = inputStrings(); err != nil {
		return err
	}

	for _, w := range input {
		if word, err = c.repo.Get(w); err == nil {
			slog.Warn(fmt.Sprintf("%s already exist", w))
			continue
		}

		if word, err = c.Save(w); err != nil {
			slog.Error(fmt.Sprintf("cannot save %s, error:", w), err.Error(), "")
			continue
		}

		fmt.Printf("*** %s added successfully!\n", strings.ToUpper(word.Word))
	}

	return nil
}

func (c *Core) ReviewHistory() (err error) {
	var (
		allReviews    []*models.Review
		totalDuration time.Duration
		totalScore    int
	)

	if allReviews, err = c.repo.GetAllReviews(); err != nil {
		return err
	}

	fmt.Println("\nMy Reviews:")
	for idx, r := range allReviews {
		totalDuration += r.Duration
		totalScore += r.Score
		if err = c.ShowReview(len(allReviews)-idx, r); err != nil {
			slog.Error(err.Error())
			continue
		}
	}
	fmt.Printf("Total: %d, Duration: %s Score: %d\n",
		len(allReviews), totalDuration.Round(time.Second).String(), totalScore)

	fmt.Println("")

	return nil
}

func (c *Core) NewTag() (err error) {
	fmt.Print("Enter new tag name: ")
	name, err := inputString()
	if err != nil {
		return
	}

	if err = c.repo.CreateTag(models.Tag{Name: name}); err != nil {
		return
	}

	fmt.Println("\ttag added successfully")

	return nil
}

func (c *Core) Tags() (err error) {
	tags, err := c.repo.GetAllTags()
	if err != nil {
		return
	}

	fmt.Println("Tags List:")
	for _, t := range tags {
		fmt.Printf("\t1- %s (%d words)\n", t.Name, len(t.Words))
	}

	return nil
}

func (c *Core) CalculateScore(t, pfc, rc int) (score int) {
	score += proficiencyScore(pfc)
	score += timeScore(t, 20)
	score += timeScore(rc, 10)
	score /= 3

	return
}

//
//func reviewCountScore(rc, limit int) int {
//	if rc > limit {
//		return 0
//	}
//}

func timeScore(t, limit int) int {
	if t > limit {
		return 0
	} else if t == limit {
		return 1
	}

	x := float64(10) / float64(limit)
	return 10 - int(float64(t)*x)
}

func proficiencyScore(pfc int) int {
	if pfc <= -16 {
		return 10
	} else if pfc >= 16 {
		return 0
	}

	return (-1 * pfc / 3) + 5
}

func printDiv() {
	fmt.Println(strings.Repeat("*", 64))
}

func inputChar() (input rune, err error) {
	reader := bufio.NewReader(os.Stdin)
	input, _, err = reader.ReadRune()
	if err != nil {
		return
	}

	return
}

func inputString() (input string, err error) {
	_, err = fmt.Scanln(&input)
	if err != nil {
		return
	}

	return
}

func inputInt() (input int, err error) {
	_, err = fmt.Scan(&input)
	if err != nil {
		return
	}

	return

}

func inputStrings() (input []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		input = append(input, line)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func clearConsole() (err error) {
	// Clearing console based on the operating system
	switch _os := runtime.GOOS; _os {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}

	return
}

func random(low, hi int64) int64 {
	return low + rand.Int63n(hi-low)
}
