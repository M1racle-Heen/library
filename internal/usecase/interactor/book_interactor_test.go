package interactor

import (
	"errors"
	"testing"

	"github.com/M1racle-Heen/library/internal/domain"
	"github.com/M1racle-Heen/library/internal/repository"
	"github.com/M1racle-Heen/library/internal/repository/mocks"
	"github.com/M1racle-Heen/library/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBookInteractor(t *testing.T) {
	type arguments struct {
		bookRepo repository.BookRepo
	}

	type expectations struct {
		uc usecase.BookUsecase
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}

	bookRepo := &mocks.BookRepo{}
	uc := bookInteractor{bookRepo: bookRepo}

	testCases := []test{
		{
			arguments: arguments{
				bookRepo: bookRepo,
			},
			expectations: expectations{
				uc: uc,
			},
		},
		{
			arguments: arguments{
				bookRepo: nil,
			},
		},
	}

	for _, tc := range testCases {
		defer func() {
			recover()
		}()

		uc := NewBookInteractor(tc.arguments.bookRepo)
		assert.Equal(t, tc.expectations.uc, uc)

	}

}

func TestBookInteractor_Create(t *testing.T) {
	type arguments struct {
		book domain.Book
	}

	type expectations struct {
		createdBook *domain.Book
		err         error
	}

	type dependencies struct {
		bookRepo *mocks.BookRepo
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(bookRepo *mocks.BookRepo, createdBook *domain.Book, err error)
	}

	book := domain.Book{}

	testCases := []test{
		{
			arguments: arguments{
				book: book,
			},
			expectations: expectations{
				createdBook: &book,
				err:         nil,
			},
			dependencies: dependencies{
				bookRepo: &mocks.BookRepo{},
			},
			prepare: func(bookRepo *mocks.BookRepo, createdBook *domain.Book, err error) {
				bookRepo.On("Store", mock.Anything).Return(createdBook, err)
			},
		},
		{
			arguments: arguments{
				book: book,
			},
			expectations: expectations{
				createdBook: nil,
				err:         errors.New("random error"),
			},
			dependencies: dependencies{
				bookRepo: &mocks.BookRepo{},
			},
			prepare: func(bookRepo *mocks.BookRepo, createdBook *domain.Book, err error) {
				bookRepo.On("Store", mock.Anything).Return(createdBook, err)
			},
		},
	}

	for _, tc := range testCases {
		tc.prepare(tc.dependencies.bookRepo, tc.expectations.createdBook, tc.expectations.err)

		interactor := bookInteractor{
			bookRepo: tc.dependencies.bookRepo,
		}

		createdBook, err := interactor.Create(tc.arguments.book)

		assert.Equal(t, tc.expectations.createdBook, createdBook)
		assert.Equal(t, tc.expectations.err, err)

	}

}
