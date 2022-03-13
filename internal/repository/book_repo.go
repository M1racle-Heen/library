package repository

import "github.com/M1racle-Heen/library/internal/domain"

type BookRepo interface {
	Store(book domain.Book) (*domain.Book, error)
	Get(id string) (*domain.Book, error)
	Remove(id string) error
}
