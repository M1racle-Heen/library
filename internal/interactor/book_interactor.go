package interactor

import (
	"context"

	"github.com/M1racle-Heen/library/internal/domain"
)

type BookInteractor struct {
}

func (b BookInteractor) CreateBook(ctx context.Context, book domain.Book) (*domain.Book, error) {
	return nil, nil
}
