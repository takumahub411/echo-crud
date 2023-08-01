package controllers

import (
	"echo-crud/initializers"
	"echo-crud/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddMovie(e echo.Context) error {
	var body models.Movie

	if err := e.Bind(&body); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}

	}

	result := initializers.DB.Create(&body)
	if result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: result.Error,
		}
	}
	return e.JSON(http.StatusOK, body)
}

func GetMovies(e echo.Context) error {
	var movies []models.Movie

	result := initializers.DB.Find(&movies)

	if result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: result.Error,
		}
		//c.AbortWithError(http.StatusNotFound, result.Error)
	}
	return e.JSON(http.StatusOK, &movies)

}

func GetMovie(e echo.Context) error {
	var movies models.Movie

	id := e.Param("id")

	result := initializers.DB.Find(&movies, id)

	if result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: result.Error,
		}
	}
	return e.JSON(http.StatusOK, &movies)

}

func UpdateMovie(e echo.Context) error {
	var movies models.Movie

	// update request
	var updateInfo models.RequestMovie

	id := e.Param("id")

	if err := e.Bind(&updateInfo); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	if result := initializers.DB.First(&movies, id); result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: result.Error,
		}
	}

	movies.Title = updateInfo.Title
	movies.LeadingRole = updateInfo.LeadingRole
	movies.Description = updateInfo.Description
	movies.Stars = updateInfo.Stars

	// update

	if updateResult := initializers.DB.Save(&movies); updateResult.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: updateResult.Error,
		}
	}

	return e.JSON(http.StatusOK, &movies)
}

func DeleteMovie(e echo.Context) error {
	var movies models.Movie

	id := e.Param("id")

	result := initializers.DB.First(&movies, id)

	if result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: result.Error,
		}
	}

	if deleteResult := initializers.DB.Delete(&movies); deleteResult.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: deleteResult.Error,
		}
	}

	return e.JSON(http.StatusOK, &movies)
}
