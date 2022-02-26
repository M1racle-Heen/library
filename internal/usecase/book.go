package usecase

type BookUsecase interface {
	CreateBook() (*domain.Book, error)
	GetBook() (*domain.Book, error)
}
