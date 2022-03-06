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

/*book.Id = uuid.New().String()
	book.AddTime = time.Now()

	_, err = repository.DBP.Exec(constants.ADD_BOOK, book.Name, book.Id, book.Category, book.Author, book.Translator, book.Owner, book.AddTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "the book didnt add")
		return
	}

	_, err = repository.DBP.Exec(constants.ADD_BOOK_CATEGORY, book.Category, book.Name, book.Author, book.Translator)
	if err != nil {
		log.Println("error on saving the data in category table its been saved in book table")
		c.JSON(http.StatusInternalServerError, "the book didnot save correctlly")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": " the book successfully added",
		"data":    book,
	})
}

func DeleteBook(c *gin.Context) {
	log.Println("deleting a book")
	name := c.Param("name")
	_, err := repository.DBP.Exec(constants.DELETE_BOOK, name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "the book didnt delete")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "the book deleted",
	})
}

func GetOneBook(c *gin.Context) {
	log.Println("get one book")
	var book models.Book
	name := c.Param("name")
	rows, err := repository.DBP.Query(constants.GET_ONE_BOOK, name)
	log.Println(rows)
	for rows.Next() {
		err = rows.Scan(&book.Name, &book.Id, &book.Category, &book.Author, &book.Translator, &book.Owner, &book.AddTime)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, "error in scanning")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "this is the book",
		"data":    book,
	})
}

func GetBookOnCategory(c *gin.Context) {
	log.Println("get book on its category")
	var book models.Book
	category := c.Param("category")
	rows, err := repository.DBP.Query(constants.GET_BOOK_ON_CATEGORY, category)
	if err != nil {
		log.Println("error on reading query")
		return
	}
	for rows.Next() {
		err := rows.Scan(&book.Category, &book.Name, &book.Author, &book.Translator)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, "error in scanning ")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   book,
	})
}

func GetAllBooks(c *gin.Context) {
	log.Println("geting all records")
	var book models.Book
	var books []models.Book
	rows, err := repository.DBP.Query(constants.GET_ALL_BOOKS)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "the books are disaccessable")
		return
	}
	for rows.Next() {
		err = rows.Scan(&book.Name)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, "error on scanning the records")
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{
		"stauts": http.StatusOK,
		"data":   books,
	})
}
*/
