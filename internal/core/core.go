package core

import (
	"bufio"
	"errors"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"lexilift/internal/models"
	"lexilift/internal/repository"
	"lexilift/pkg/dictionary"
	"log/slog"
	"os"
	"sort"
	"strings"
	"time"
)

type Core struct {
	repo *repository.Repo
	dict *dictionary.API
	//ply   *player.Player
	debug bool
}

// func New(repo *repository.Repo, dict *dictionary.API, ply *player.Player, debug bool) *Core {

func New(repo *repository.Repo, dict *dictionary.API, debug bool) *Core {
	return &Core{
		repo: repo,
		dict: dict,
		//ply:   ply,
		debug: debug,
	}
}

func (c *Core) Dashboard() (err error) {
	var (
		allWords      []*models.Word
		allReviews    []*models.Review
		knowMap       = make(map[int]int)
		sorted        []int
		totalDuration time.Duration
	)

	fmt.Println(strings.Repeat(">", 32), "Dashboard", strings.Repeat("<", 32))

	if allWords, err = c.repo.GetAll(); err != nil {
		return err
	}

	for _, w := range allWords {
		knowMap[w.Proficiency] += 1
	}

	for k, _ := range knowMap {
		sorted = append(sorted, k)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	fmt.Println("My Words:")
	for _, i := range sorted {
		kn := knowMap[i]
		fmt.Printf("Proficiency: %d\t Count: %d\n", i, kn)
	}

	fmt.Printf("Total: %d\n", len(allWords))

	if allReviews, err = c.repo.GetAllReviews(); err != nil {
		return err
	}

	for _, r := range allReviews {
		totalDuration += r.Duration
	}

	fmt.Println("\nMy Reviews:")
	fmt.Printf("Total: %d, Duration: %s\n", len(allReviews), totalDuration.Round(time.Second).String())
	for idx, r := range allReviews {
		if err = c.ShowReview(len(allReviews)-idx, r); err != nil {
			slog.Error(err.Error())
			continue
		}

		if idx == 2 {
			break
		}
	}

	fmt.Println("")

	return nil
}

func (c *Core) ReviewHistory() (err error) {
	var (
		allReviews    []*models.Review
		totalDuration time.Duration
	)

	if allReviews, err = c.repo.GetAllReviews(); err != nil {
		return err
	}

	fmt.Println("\nMy Reviews:")
	for idx, r := range allReviews {
		totalDuration += r.Duration
		if err = c.ShowReview(len(allReviews)-idx, r); err != nil {
			slog.Error(err.Error())
			continue
		}
	}
	fmt.Printf("Total: %d, Duration: %s\n", len(allReviews), totalDuration.Round(time.Second).String())

	fmt.Println("")

	return nil
}

func (c *Core) Menu() (err error) {
	var (
		input rune
	)
	fmt.Println(strings.Repeat(">", 32), "LexiLift", strings.Repeat("<", 32))
	fmt.Println("Menu:")
	fmt.Println("\t0- Dashboard")
	fmt.Println("\t1- Review my words")
	fmt.Println("\t2- Add a new word to my words")
	fmt.Println("\t3- Add words list to my words")
	fmt.Println("\t4- Review history")
	fmt.Println("\tq- close the app")
	fmt.Printf("press number of action you want do:")
	input, err = inputChar()
	if err != nil {
		return err
	}

	if c.debug {
		fmt.Printf("input: %c\n", input)
	}

	switch input {
	case '0':
		return c.Dashboard()
	case '1':
		return c.Review()
	case '2':
		return c.AddNewWord()
	case '3':
		return c.AddWordsList()
	case '4':
		return c.ReviewHistory()
	case 'q':
		fmt.Println("See you soon, bye :)")
		os.Exit(0)
	default:
		return c.Menu()
	}

	return
}

func (c *Core) Save(word string) (w *models.Word, err error) {
	var (
		mean  string
		dict  *models.Dictionary
		sound string
	)

	fmt.Println("getting persian translate from google translate API ...")
	if mean, err = gt.Translate(word, "en", "fa"); err != nil {
		return
	}

	fmt.Println("getting translate from dictionary API ...")
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
		slog.Error("can not save to DB:", err.Error())
	}

	return
}

func (c *Core) Review() (err error) {
	var (
		words          []*models.Word
		input          rune
		fromKnw, toKnw int
		startedAt      = time.Now().Local()
		know, notKnow  int
	)

	// TODO review history

	fmt.Print("from know count: ")
	if fromKnw, err = inputInt(); err != nil {
		return err
	}

	fmt.Print("to know count: ")
	if toKnw, err = inputInt(); err != nil {
		return err
	}

	if words, err = c.repo.Fetch(fromKnw, toKnw, 1000, 0); err != nil {
		return err
	}

	printDiv()

	fmt.Printf("%d words founded\n", len(words))
	if len(words) != 0 {
		fmt.Println("Starting ...")
		fmt.Println("press Enter to view word meaning")

		for idx, word := range words {
			word.ReviewCount += 1
		review:
			if err = c.ShowWord(idx, word); err != nil {
				slog.Error(err.Error())
			}

			fmt.Println("(1) I know, (2) I don't know, (0) nothing, (a) add a new word, (q) stop")
			if input, err = inputChar(); err != nil {
				return err
			}

			if input == 'q' {
				break
			}

			switch input {
			case '1':
				word.Proficiency += 1
				know += 1
			case '2':
				word.Proficiency -= 1
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
	}

	review := models.Review{
		StartedAt:       startedAt,
		Duration:        time.Now().Sub(startedAt),
		FromProficiency: fromKnw,
		ToProficiency:   toKnw,
		Total:           len(words),
		Know:            know,
		NotKnow:         notKnow,
	}

	if err = c.repo.CreateReview(review); err != nil {
		return
	}

	return
}

func (c *Core) ShowReview(idx int, review *models.Review) (err error) {
	fmt.Printf("%d- %s  %s\tFP:%d\tTP:%d\tCNT:%d\tKNW:%d\tNK:%d\n",
		idx,
		review.StartedAt.Format("2006-01-02 15:04"),
		review.Duration.Round(time.Second).String(),
		review.FromProficiency,
		review.ToProficiency,
		review.Total,
		review.Know,
		review.NotKnow,
	)
	return nil
}

func (c *Core) ShowWord(idx int, word *models.Word) (err error) {
	printDiv()
	fmt.Printf("\t%d- %s\ncreated at: %s, reviewd: %d, proficiency: %d)\n",
		idx+1, word.Word, word.CreatedAt.Format("2006-01-02 15:04"), word.ReviewCount, word.Proficiency)

	//if word.SoundFile != "" {
	//	if err = c.ply.Play(word.SoundFile); err != nil {
	//		slog.Error(err.Error())
	//	}
	//}
	_, err = inputChar()
	if err != nil {
		return err
	}

	fmt.Println("\tmeanings:")
	if word.Dict != nil && len(word.Dict.Meanings) != 0 {
		for _, m := range word.Dict.Meanings[0].Definitions {
			fmt.Printf("\t- %s\n", m.Definition)
			if m.Example != "" {
				fmt.Printf("\texample: %s\n", m.Example)
			}
		}
	}
	fmt.Printf("\tpersian: %s\n", word.Mean)

	return nil
}

func (c *Core) AddNewWord() (err error) {
	var (
		input string
		word  *models.Word
	)

	printDiv()
	fmt.Print("type the word: ")
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
	fmt.Print("type each word in new line, submit empty line to finish: \n")
	if input, err = inputStrings(); err != nil {
		return err
	}

	for _, w := range input {
		if word, err = c.repo.Get(w); err == nil {
			slog.Warn(fmt.Sprintf("%s already exist", w))
			continue
		}

		if word, err = c.Save(w); err != nil {
			slog.Error(fmt.Sprintf("can not save %s, error:", w), err.Error())
			continue
		}

		fmt.Printf("*** %s added successfuly!\n", strings.ToUpper(word.Word))
	}

	return nil
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
