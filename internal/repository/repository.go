package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lexilift/internal/models"
	"log"
	"log/slog"
	"os"
	"time"
)

type Repo struct {
	db *gorm.DB
}

func New(debug bool) (repo *Repo, err error) {
	var (
		newLogger logger.Interface
	)

	if debug {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	} else {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  logger.Error, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,         // Disable color
			},
		)
	}
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return
	}

	if err = db.Migrator().AutoMigrate(
		&models.Word{},
		&models.Review{},
	); err != nil {
		slog.Error(err.Error())
	}

	repo = &Repo{
		db: db,
	}

	return
}

func (r *Repo) Get(word string) (w *models.Word, err error) {
	if err = r.db.Where("word = ?", word).First(&w).Error; err != nil || w.ID == 0 {
		return nil, err
	}

	return
}

func (r *Repo) Create(w models.Word) (err error) {
	return r.db.Create(&w).Error
}

func (r *Repo) CreateReview(review models.Review) (err error) {
	return r.db.Create(&review).Error
}

func (r *Repo) Count() (total int64, err error) {
	if err = r.db.Model(&models.Word{}).Count(&total).Error; err != nil {
		return
	}

	return
}

func (r *Repo) Fetch(fromKnw, toKnw, limit, offset int) (words []*models.Word, err error) {
	if err = r.db.Where("proficiency >= ? AND proficiency <= ?", fromKnw, toKnw).
		Limit(limit).Offset(offset).Find(&words).Error; err != nil {
		return
	}

	return
}

func (r *Repo) GetAll() (words []*models.Word, err error) {
	if err = r.db.Order("score desc").Find(&words).Error; err != nil {
		return
	}

	return
}

func (r *Repo) GetAllReviews() (reviews []*models.Review, err error) {
	if err = r.db.Order("started_at desc").Find(&reviews).Error; err != nil {
		return
	}

	return
}

func (r *Repo) Update(word *models.Word) (err error) {
	if err = r.db.Updates(word).Error; err != nil {
		return err
	}

	return err
}

func (r *Repo) GetAllTags() (tags []*models.Tag, err error) {
	if err = r.db.Find(&tags).Error; err != nil {
		return
	}

	return
}

func (r *Repo) CreateTag(t models.Tag) (err error) {
	return r.db.Create(&t).Error
}
