package repository

import (
	"log/slog"

	"github.com/sinameshkini/microkit/pkg/clients/database"
	"gorm.io/gorm"

	"lexilift/internal/config"
	"lexilift/internal/models"
)

type Repo struct {
	db *gorm.DB
}

func New(conf *config.Config) (repo *Repo, err error) {
	db, err := database.NewSQLite(conf.DatabasePath, conf.Debug)
	if err != nil {
		return nil, err
	}

	if err = database.Migrate(db.DB(), models.AllTables); err != nil {
		slog.Error(err.Error())
	}

	repo = &Repo{
		db: db.DB(),
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
