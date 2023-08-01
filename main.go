package main

import (
	"echo-crud/controllers"
	"echo-crud/initializers"
	"github.com/labstack/echo/v4"
	"os"
)

func init() {
	initializers.LoadEnvfile()
	initializers.ConnectDb()
	initializers.SyncDatabase()

}

func main() {
	e := echo.New()
	e.POST("/signup", controllers.CreateUser)
	e.POST("/login", controllers.LoginUser)

	e.POST("/movies", controllers.AddMovie)
	e.GET("/movies", controllers.GetMovies)
	e.GET("/movies/:id", controllers.GetMovie)
	e.PUT("/movies/:id", controllers.UpdateMovie)
	e.DELETE("/movies/:id", controllers.DeleteMovie)
	e.Logger.Fatal(e.Start(os.Getenv("port")))
}
