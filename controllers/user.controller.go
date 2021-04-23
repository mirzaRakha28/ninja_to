package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mirzaRakha28/ninja_to/models"
)

func Register(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")
	jenjang := c.FormValue("jenjang")

	result, err := models.Register(email, username, password, jenjang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
