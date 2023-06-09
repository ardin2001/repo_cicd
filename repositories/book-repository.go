package repositories

import (
	"unit_testing/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooksRepository() ([]*models.Book, error)
	GetBookRepository(id string) (*models.Book, error)
	CreateRepository(Book models.Book) (*models.Book, error)
	UpdateRepository(id string, BookBody models.Book) (*models.Book, error)
	DeleteRepository(id string) error
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) BookRepository {
	return &BookRepo{
		DB: DB,
	}
}

func (b *BookRepo) GetBooksRepository() ([]*models.Book, error) {
	var Books []*models.Book

	if err := b.DB.Find(&Books).Error; err != nil {
		return nil, err
	}

	return Books, nil
}

func (b *BookRepo) GetBookRepository(id string) (*models.Book, error) {
	var Book *models.Book

	if err := b.DB.Where("ID = ?", id).Take(&Book).Error; err != nil {
		return nil, err
	}

	return Book, nil
}

func (b *BookRepo) CreateRepository(Book models.Book) (*models.Book, error) {
	if err := b.DB.Save(&Book).Error; err != nil {
		return nil, err
	}

	return &Book, nil
}

func (b *BookRepo) UpdateRepository(id string, BookBody models.Book) (*models.Book, error) {
	Book, err := b.GetBookRepository(id)
	if err != nil {
		return nil, err
	}

	err = b.DB.Where("ID = ?", id).Updates(models.Book{Title: BookBody.Title, Author: BookBody.Author, Description: BookBody.Description}).Error
	if err != nil {
		return nil, err
	}

	Book.Title = BookBody.Title
	Book.Author = BookBody.Author
	Book.Description = BookBody.Description

	return Book, nil
}

func (b *BookRepo) DeleteRepository(id string) error {
	_, err := b.GetBookRepository(id)
	if err != nil {
		return err
	}

	if err := b.DB.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}

	return nil
}
