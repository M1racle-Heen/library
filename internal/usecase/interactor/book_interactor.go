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

func (interactor bookInteractor) Create(book domain.Book) (*domain.Book, error) {

	createdBook, err := interactor.bookRepo.Store(book)

	if err != nil {
		return nil, err
	}

	return createdBook, nil

}

func (interactor bookInteractor) Get(id string) (*domain.Book, error) {
	book, err := interactor.bookRepo.Get(id)

	if err != nil {
		return nil, err
	}

	return book, nil

}

func (interactor bookInteractor) Delete(id string) error {
	return interactor.bookRepo.Remove(id)
}

// func TestBookInteractor_Get(t *testing.T) {
// 	type arguments struct {
// 		id string
// 	}
// 	type expectations struct {
// 		getBook *domain.Book
// 		err     error
// 	}
// 	type dependencies struct {
// 		bookRepo *mocks.BookRepo
// 	}

// 	type test struct {
// 		arguments    arguments
// 		expectations expectations
// 		dependencies dependencies
// 		prepare      func(bookRepo *mocks.BookRepo, getBook *domain.Book, err error)
// 	}

// 	id := "123"
// 	book := domain.Book{}
// 	testCases := []test{
// 		{
// 			arguments:    arguments{id: id},
// 			expectations: expectations{getBook: &book, err: nil},
// 			dependencies: dependencies{bookRepo: &mocks.BookRepo{}},
// 			prepare: func(bookRepo *mocks.BookRepo, getBook *domain.Book, err error) {
// 				bookRepo.On("Get", mock.Anything).Return(getBook, err)
// 			},
// 		},
// 		{
// 			arguments:    arguments{id: id},
// 			expectations: expectations{getBook: nil, err: errors.New("Some error")},
// 			dependencies: dependencies{bookRepo: &mocks.BookRepo{}},
// 			prepare: func(bookRepo *mocks.BookRepo, getBook *domain.Book, err error) {
// 				bookRepo.On("Get", mock.Anything).Return(getBook, err)
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {

// 		tc.prepare(tc.dependencies.bookRepo, tc.expectations.getBook, tc.expectations.err)

// 		interactor := bookInteractor{bookRepo: tc.dependencies.bookRepo}

// 		getBook, err := interactor.Get(tc.arguments.id)

// 		assert.Equal(t, tc.expectations.getBook, getBook)
// 		assert.Equal(t, tc.expectations.err, err)
// 	}
// }
