package moduls

import (
	"gorm.io/gorm"
)

type Repo interface {
    Create(b Book) (Book, error)
    Update(id int, b Book) (Book, error)
    First() (Book, error)
    Last() (Book, error)
    FindByID(ID int) (Book, error)
    All() ([]Book, error)
    Delete(ID int) error
    // Filter(Rating int) ([]Book, error)
}

type repo struct {
    db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
    return &repo{db}
}

func (r *repo) All() ([]Book, error) {
    var books []Book
    err := r.db.Find(&books).Error

    return books, err
}

func (r *repo) First() (Book, error) {
    var book Book
    err := r.db.First(&book).Error

    return book, err
}

func (r *repo) Last() (Book, error) {
    var book Book
    err := r.db.Last(&book).Error

    return book, err
}

func (r *repo) Delete(ID int) error{
    var book Book
    return r.db.Where("deleted_at IS NULL AND ID = ?", ID).Delete(&book).Error
}

func (r *repo) FindByID(ID int) (Book, error){
    var book Book
    err := r.db.Where("deleted_at IS NULL AND ID = ?", ID).First(&book).Error

    return book, err
}

func (r *repo) Create(b Book) (Book, error) {
    err := r.db.Create(&b).Error

    return b, err
}

func (r *repo) Update(id int, b Book) (Book, error) {
    err := r.db.Where("deleted_at IS NULL AND id = ?", id).Updates(&b).Error

    return b, err
}

