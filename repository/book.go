package repository

import (
	"context"
	"log"
	"my_lib/constants"
	"my_lib/errorhandler"
	"my_lib/models"
)

type RepoInterface interface {
	AddBook(ctx context.Context, bookData models.Book) error
	DeleteBook(ctx context.Context, name string) error
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

func (b *book) DeleteBook(ctx context.Context, name string) error {
	_, err := DBP.Exec(constants.DELETE_BOOK, name)
	if err != nil {
		log.Println(errorhandler.LayerConnectionError + "repository")
		return err
	}
	return nil
}
