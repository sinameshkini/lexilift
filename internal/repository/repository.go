package repository

import (
	"gorm.io/gorm"
	"lexilift/internal/repository/entities"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) (repo *Repo, err error) {
	repo = &Repo{
		db: db,
	}

	return
}

func (r *Repo) Get(word string) (w *entities.Word, err error) {
	if err = r.db.Where("word = ?", word).First(&w).Error; err != nil || w.ID == 0 {
		return nil, err
	}

	return
}

func (r *Repo) Create(w entities.Word) (err error) {
	return r.db.Create(&w).Error
}

func (r *Repo) CreateReview(review entities.Review) (err error) {
	return r.db.Create(&review).Error
}

func (r *Repo) Count() (total int64, err error) {
	if err = r.db.Model(&entities.Word{}).Count(&total).Error; err != nil {
		return
	}

	return
}

func (r *Repo) Fetch(fromKnw, toKnw, limit, offset int) (words []*entities.Word, err error) {
	if err = r.db.Where("proficiency >= ? AND proficiency <= ?", fromKnw, toKnw).
		Limit(limit).Offset(offset).Find(&words).Error; err != nil {
		return
	}

	return
}

func (r *Repo) GetAll() (words []*entities.Word, err error) {
	if err = r.db.Order("score desc").Find(&words).Error; err != nil {
		return
	}

	return
}

func (r *Repo) GetAllReviews() (reviews []*entities.Review, err error) {
	if err = r.db.Order("started_at desc").Find(&reviews).Error; err != nil {
		return
	}

	return
}

func (r *Repo) Update(word *entities.Word) (err error) {
	if err = r.db.Updates(word).Error; err != nil {
		return err
	}

	return err
}

func (r *Repo) GetAllTags() (tags []*entities.Tag, err error) {
	if err = r.db.Find(&tags).Error; err != nil {
		return
	}

	return
}

func (r *Repo) CreateTag(t entities.Tag) (err error) {
	return r.db.Create(&t).Error
}
