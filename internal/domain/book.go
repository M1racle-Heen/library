package domain

import "time"

type Book struct {
	ID          string
	Name        string
	Description string
	ReleaseDate time.Time
	Author      string
	Genre       string
}
