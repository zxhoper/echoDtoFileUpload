package controllers

import (
	"echoDtoFileUpload/dtos"
	"echoDtoFileUpload/models"
	"echoDtoFileUpload/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FileUpload(c echo.Context) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	return c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}

// RemoteUpload function does the same thing as the FileUpload function.
// However, we created url variable and validate it using the Echo’s Bind method.
// We also passed the variable to the RemoteUpload service as an argument
//    and returned the appropriate response.
func RemoteUpload(c echo.Context) error {
	var url models.Url

	//validate the request body
	if err := c.Bind(&url); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Error uploading file"},
			})
	}

	return c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
