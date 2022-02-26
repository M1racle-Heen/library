package repository

import (
	"context"

	"github.com/M1racle-Heen/library/internal/domain"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book domain.Book) (*domain.Book, error)
	GetByID(ctx context.Context, id string) (*domain.Book, error)
	GetMany(ctx context.Context, limit, offset int) ([]domain.Book, error)
}
