package handller

import (
	"fmt"
	"log"
	"my_lib/errorhandler"
	"my_lib/models"
	"my_lib/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddNewBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type book struct {
	logic usecase.UsecaseInterface
}

func NewHandlerInterface(logic usecase.UsecaseInterface) HandlerInterface {
	return &book{
		logic: logic,
	}
}
func Homepage(c *gin.Context) {
	fmt.Sprintf("welcome to your library")
	c.JSON(http.StatusOK, "you can now manage your library")
}

func (b *book) AddNewBook(c *gin.Context) {
	log.Println("adding new book")
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		log.Println(errorhandler.BindError)
		c.JSON(http.StatusInternalServerError, "error in binding")
		return
	}
	newbook, err := b.logic.AddBook(c, book)
	if err != nil {
		log.Println(errorhandler.LayerConnectionError + "in handler")
		c.JSON(http.StatusInternalServerError, "error on connecting to lower layer")
		return
	}
	c.JSON(http.StatusOK, newbook)
}

func (b *book) DeleteBook(c *gin.Context) {
	log.Println("deleting a record")
	name := c.Param("name")
	err := b.logic.DeleteBook(c, name)
	if err != nil {
		log.Println(errorhandler.LayerConnectionError + "in handler")
		c.JSON(http.StatusInternalServerError, "error on sending data to logic")
		return
	}
	c.JSON(http.StatusOK, "deleted successfully")
}
