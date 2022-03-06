package usecase

import (
	"context"
	"log"
	"my_lib/errorhandler"
	"my_lib/models"
	"my_lib/repository"
	"time"

	"github.com/google/uuid"
)

type UsecaseInterface interface {
	AddBook(ctx context.Context, bookDate models.Book) (newbook *models.Book, err error)
	DeleteBook(ctx context.Context, name string) error
}

type book struct {
	repo repository.RepoInterface
}

func NewUsecaseInterface(repo repository.RepoInterface) UsecaseInterface {
	return &book{
		repo: repo,
	}
}

func (b *book) AddBook(ctx context.Context, bookData models.Book) (newbook *models.Book, err error) {
	bookData.AddTime = time.Now()
	bookData.Id = uuid.New().String()
	err = b.repo.AddBook(ctx, models.Book{
		Name:       bookData.Name,
		Id:         bookData.Id,
		Category:   bookData.Category,
		Author:     bookData.Author,
		Translator: bookData.Translator,
		Owner:      bookData.Owner,
		AddTime:    bookData.AddTime,
	})
	if err != nil {
		log.Println(errorhandler.LayerConnectionError + "usecase")
		log.Println("error on setting data in object in usecase")
		return nil, err
	}
	return &bookData, nil
}

func (b *book) DeleteBook(ctx context.Context, name string) error {
	err := b.repo.DeleteBook(ctx, name)
	if err != nil {
		log.Println(errorhandler.LayerConnectionError + "usecase")
		return err
	}
	return nil
}
