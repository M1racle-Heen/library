package interactor

import (
	"context"

	"github.com/M1racle-Heen/library/internal/domain"
	"github.com/M1racle-Heen/library/internal/repository"
	"github.com/M1racle-Heen/library/internal/usecase"
)

type bookInteractor struct {
	bookRepo repository.BookRepository
}

func NewBookInteractor(bookRepo repository.BookRepository) usecase.BookUsecase {
	return bookInteractor{bookRepo: bookRepo}
}

func (interactor bookInteractor) CreateBook(ctx context.Context, book domain.Book) (*domain.Book, error) {

	createBook, err := interactor.bookRepo.CreateBook(ctx, book)

	if err != nil {
		return nil, err
	}
	return createBook, nil

}

func (interactor bookInteractor) GetByID(ctx context.Context, id string) (*domain.Book, error) {
	book, err := interactor.bookRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (interactor bookInteractor) GetMany(ctx context.Context, limit, offset int) ([]domain.Book, error) {
	books, err := interactor.bookRepo.GetMany(ctx, limit, offset)
	if err != nil {
		return books, err
	}
	return books, nil
}
