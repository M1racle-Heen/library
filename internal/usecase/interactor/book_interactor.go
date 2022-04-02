package interactor

import (
	"context"

	"github.com/M1racle-Heen/library/internal/domain"
	"github.com/M1racle-Heen/library/internal/repository"
	"github.com/M1racle-Heen/library/internal/usecase"
)

type bookInteractor struct {
	bookRepo repository.BookRepo
}

func NewBookInteractor(bookRepo repository.BookRepo) usecase.BookUsecase {
	if bookRepo == nil {
		panic("bookRepo argument is empty")
	}

	return bookInteractor{
		bookRepo: bookRepo,
	}

}

func (interactor bookInteractor) Create(ctx context.Context, book domain.Book) (*domain.Book, error) {

	createdBook, err := interactor.bookRepo.Store(ctx, book)

	if err != nil {
		return nil, err
	}

	return createdBook, nil

}

func (interactor bookInteractor) Get(ctx context.Context, id string) (*domain.Book, error) {
	book, err := interactor.bookRepo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return book, nil

}

func (interactor bookInteractor) Delete(ctx context.Context, id string) error {
	return interactor.bookRepo.Remove(ctx, id)
}
