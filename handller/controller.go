package handller

import (
	"my_lib/repository"
	"my_lib/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {
	Router := gin.Default()
	bookHandler := NewHandlerInterface(usecase.NewUsecaseInterface(repository.NewRepoInterface()))
	Router.POST("/", Homepage)
	Router.POST("/addbook", bookHandler.AddNewBook)
	Router.DELETE("/book/del/:name", bookHandler.DeleteBook)
	/*Router.GET("/book/getbook/:name", GetOneBook)
	Router.GET("/book/getoncat/:category", GetBookOnCategory)
	Router.GET("/getall", GetAllBooks)
	*/
	Router.Run(":4141")
}
