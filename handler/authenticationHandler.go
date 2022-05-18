package handler

import (
	"CompanyAPI/ent"
	"CompanyAPI/ent/user"
	"CompanyAPI/request"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type authenticationHandler struct {
	db *ent.Client
}

func NewAuthenticationHandler(db *ent.Client) *authenticationHandler {
	return &authenticationHandler{db}
}

func (authHandler *authenticationHandler) Register(c echo.Context) error {
	ctx := context.Background()
	req := new(request.UserRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user, wrong := authHandler.db.User.Create().SetEmail(req.Email).SetName(req.Name).SetPassword(string(password)).SetUsername(req.Username).Save(ctx)

	if wrong != nil {
		return c.JSON(http.StatusBadRequest, wrong.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (authHandler *authenticationHandler) Login(c echo.Context) error {
	ctx := context.Background()
	req := new(request.LoginRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	if checkPassword := authHandler.db.User.Query().Where(user.Username(req.Username)).ExistX(ctx); checkPassword != true {
		return c.JSON(http.StatusBadRequest, "Username Not Found!")
	}

	password := authHandler.db.User.Query().Where(user.Username(req.Username)).Select(user.FieldPassword).StringX(ctx)

	if check := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); check != nil {
		return c.JSON(http.StatusBadRequest, "Wrong Username/Password!")
	}

	return c.JSON(http.StatusOK, "Login!")
}
