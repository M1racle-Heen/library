package bookmongorepo

import (
	"time"

	"github.com/M1racle-Heen/library/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson: "name"`
	Description string             `bson: "description"`
	ReleaseDate time.Time          `bson: "releaseDate"`
	Author      string             `bson: "author"`
	Genre       string             `bson: "genre"`
}

func toBookScheme(book domain.Book) Book {
	return Book{
		ID:          primitive.NewObjectID(),
		Name:        book.Name,
		Description: book.Description,
		ReleaseDate: book.ReleaseDate,
		Author:      book.Author,
		Genre:       book.Genre,
	}
}

func toBook(bookSchema Book) domain.Book {
	return domain.Book{
		ID:          bookSchema.ID.Hex(),
		Name:        bookSchema.Name,
		Description: bookSchema.Description,
		ReleaseDate: bookSchema.ReleaseDate,
		Author:      bookSchema.Author,
		Genre:       bookSchema.Genre,
	}
}
