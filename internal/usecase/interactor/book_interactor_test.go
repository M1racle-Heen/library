package interactor

import (
	"testing"

	"github.com/M1racle-Heen/library/internal/repository"
	"github.com/M1racle-Heen/library/internal/repository/mocks"
	"github.com/M1racle-Heen/library/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewBookInteractor(t *testing.T) {
	type arguments struct {
		BookRepo repository.BookRepository
	}

	type expectations struct {
		uc usecase.BookUsecase
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}

	bookRepo := &mocks.BookRepository{}
	uc := bookInteractor{bookRepo: bookRepo}

	testCases := []test{
		{
			arguments:    arguments{BookRepo: bookRepo},
			expectations: expectations{uc: uc},
		},
		{
			arguments: arguments{BookRepo: nil},
		},
	}

	for _, tc := range testCases {
		defer func() {
			recover()
		}()

		uc := NewBookInteractor(tc.arguments.BookRepo)
		assert.Equal(t, tc.expectations.uc, uc)
	}
}
