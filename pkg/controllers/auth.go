package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

type RegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
	Name     string `validate:"required,min=3,max=32" json:"name"`
	Phone    string `validate:"e164" json:"phone"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	ZipCode  int    `validate:"number" form:"zip_code" json:"zip_code"`
	Adress   string `json:"address"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func (AuthController) Register(c *fiber.Ctx) error {
	u := new(RegisterRequest)
	utils.ParseBody(c, u)

	validate := utils.ValidateStruct(*u)
	if validate != nil {
		return c.Status(fiber.StatusForbidden).JSON(validate)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return utils.ReturnErrorMessage(c, err)
	}

	_, err = api.App.DB.Ent.User.Create().SetEmail(u.Email).
		SetPassword(string(password)).
		SetName(u.Name).
		SetPhone(u.Phone).
		SetCity(u.City).
		SetState(u.State).
		SetCountry(u.Country).
		SetZipCode(u.ZipCode).
		SetAddress(u.Adress).
		Save(context.Background())

	if err != nil {
		return utils.ReturnErrorMessage(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully!",
		"user":    u,
	})
}

func (AuthController) Login(c *fiber.Ctx) error {
	u := new(LoginRequest)
	utils.ParseBody(c, u)

	validate := utils.ValidateStruct(*u)
	if validate != nil {
		return c.Status(fiber.StatusForbidden).JSON(validate)
	}

	// Check exists
	user, err := api.App.DB.Ent.User.Query().Where(user.EmailEQ(u.Email)).First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return utils.ReturnErrorMessage(c, errors.New("user not found"), fiber.StatusNotFound)
		}
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err == nil {
		claims := jwt.MapClaims{
			"fields": user,
			"exp":    time.Now().Add(time.Hour * api.App.Config.Middleware.Jwt.Hours).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(api.App.Config.Middleware.Jwt.Key))
		if err != nil {
			return utils.ReturnErrorMessage(c, err)
		}

		return c.JSON(fiber.Map{"message": "User logged in successfully!",
			"token": t,
		})

	}

	return utils.ReturnErrorMessage(c, errors.New("check password"), fiber.StatusUnauthorized)
}
