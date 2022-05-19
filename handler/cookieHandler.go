package handler

import (
	"CompanyAPI/ent"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type cookieHandler struct {
	db *ent.Client
}

func NewCookieHandler(db *ent.Client) *cookieHandler {
	return &cookieHandler{db}
}

func (cookHandler *cookieHandler) CreateCookie(c echo.Context) error {
	author := new(http.Cookie)

	author.Name = "name"
	author.Value = "Ghema Allan Ferdiansyah"

	c.SetCookie(author)

	return c.JSON(http.StatusOK, author)
}

func (cookHandler *cookieHandler) ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("name")

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, cookie)
}

func (cookHandler *cookieHandler) DeleteCookie(c echo.Context) error {
	cookie := &http.Cookie{
		Name:    "name",
		Value:   "",
		Expires: time.Unix(0, 0),
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, cookie)
}
