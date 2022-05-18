package handler

import (
	"CompanyAPI/ent"
	"CompanyAPI/ent/user"
	"CompanyAPI/request"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

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

	duplicateUser := authHandler.db.User.Query().Where(user.Username(req.Username)).ExistX(ctx)
	duplicateEmail := authHandler.db.User.Query().Where(user.Email(req.Email)).ExistX(ctx)

	if duplicateUser == true && duplicateEmail == true {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"username": "Username already exist!",
			"email":    "Email already exist!",
		})
	}
	if duplicateUser == false && duplicateEmail == true {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"email": "Email already exist!",
		})
	}
	if duplicateUser == true && duplicateEmail == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"username": "Username already exist!",
		})
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

	claims := &jwtCustomClaim{
		authHandler.db.User.Query().Where(user.Username(req.Username)).Select(user.FieldName).StringX(ctx),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (authHandler *authenticationHandler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaim)

	return c.JSON(http.StatusOK, echo.Map{
		"name": "This user name is " + claims.Name,
	})
}

func Config() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaim{},
		SigningKey: []byte("secret"),
	}
}
