package routes

import (
	"log"
	"os"
	"unit_testing/configs"
	c "unit_testing/controllers"
	m "unit_testing/middlewares"
	r "unit_testing/repositories"
	s "unit_testing/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB = configs.InitDB()

	JWT = m.NewJWTS()

	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS, JWT)

	bookR = r.NewBookRepository(DB)
	bookS = s.NewBookService(bookR)
	bookC = c.NewBookController(bookS)
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	m.LoggerMiddleware(e)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/users", userC.GetUsersController)
	auth.GET("/users/:id", userC.GetUserController)
	e.POST("/users", userC.CreateController)
	auth.DELETE("/users/:id", userC.DeleteController)
	auth.PUT("/users/:id", userC.UpdateController)

	e.GET("/books", bookC.GetBooksController)
	e.GET("/books/:id", bookC.GetBookController)
	auth.POST("/books", bookC.CreateController)
	auth.DELETE("/books/:id", bookC.DeleteController)
	auth.PUT("/books/:id", bookC.UpdateController)

	return e
}
