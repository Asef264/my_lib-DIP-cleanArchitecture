package repository

import (
	"context"
	"log"
	"my_lib/constants"
	"my_lib/models"
)

type RepoInterface interface {
	AddBook(ctx context.Context, bookData models.Book) error
}

type book struct{}

func NewRepoInterface() RepoInterface {
	return &book{}
}

func (b *book) AddBook(ctx context.Context, bookData models.Book) error {
	_, err := DBP.Exec(constants.ADD_BOOK, bookData.Name, bookData.Id, bookData.Category, bookData.Author, bookData.Translator, bookData.Owner, bookData.AddTime)
	if err != nil {
		log.Println("error on adding data to the table")
		return err
	}
	log.Println("the book successfully added")
	return nil
}
