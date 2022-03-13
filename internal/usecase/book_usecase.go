package usecase

import "github.com/M1racle-Heen/library/internal/domain"

type BookUsecase interface {
	Create(book domain.Book) (*domain.Book, error)
	Get(id string) (*domain.Book, error)
	Delete(id string) error
}
