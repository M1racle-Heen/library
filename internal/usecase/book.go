package usecase

type BookUsecase interface {
	CreateBook() (*domain.Book, error)
	GetBook(id string) (*domain.Book, error)
	GetMany(limit, offset int) ([]domain.Book, error)
}
