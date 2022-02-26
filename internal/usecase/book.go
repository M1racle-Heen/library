package usecase

import "github.com/M1racle-Heen/library/internal/domain"

type BookUsecase interface {
	CreateBook() (*domain.Book, error)
	GetBook(id string) (*domain.Book, error)
	GetMany(limit, offset int) ([]domain.Book, error)
}
