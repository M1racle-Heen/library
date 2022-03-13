package interactor

import (
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

func (interactor bookInteractor) Create(
	book domain.Book,
) (*domain.Book, error) {

	createdBook, err := interactor.bookRepo.Store(book)

	if err != nil {
		return nil, err
	}

	return createdBook, nil

}

func (interactor bookInteractor) Get(
	id string,
) (*domain.Book, error) {
	book, err := interactor.bookRepo.Get(id)

	if err != nil {
		return nil, err
	}

	return book, nil

}

func (interactor bookInteractor) Delete(id string) error {
	return interactor.bookRepo.Remove(id)
}
