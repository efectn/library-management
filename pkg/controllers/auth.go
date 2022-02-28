package controllers

import (
	"context"
	"time"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/efectn/library-management/pkg/utils/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

type RegisterRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
	Name     string `validate:"required,min=3,max=32" json:"name"`
	Phone    string `validate:"omitempty,e164" json:"phone,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
	ZipCode  int    `validate:"number" form:"zip_code" json:"zip_code"`
	Address  string `json:"address,omitempty"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func (AuthController) Register(c *fiber.Ctx) error {
	u := new(RegisterRequest)
	if err := utils.ParseAndValidate(c, u); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	_, err := api.App.DB.Ent.User.Create().SetEmail(u.Email).
		SetPassword(u.Password).
		SetName(u.Name).
		SetPhone(u.Phone).
		SetCity(u.City).
		SetState(u.State).
		SetCountry(u.Country).
		SetZipCode(u.ZipCode).
		SetAddress(u.Address).
		Save(context.Background())

	if ent.IsConstraintError(err) {
		return errors.NewErrors(fiber.StatusForbidden, "This email address is not available for sign up, please try something else")
	} else if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully!",
		"user":    u,
	})
}

func (AuthController) Login(c *fiber.Ctx) error {
	u := new(LoginRequest)
	if err := utils.ParseAndValidate(c, u); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	// Check exists
	user, err := api.App.DB.Ent.User.Query().Where(user.EmailEQ(u.Email)).First(context.Background())
	if ent.IsNotFound(err) {
		return errors.NewErrors(fiber.StatusNotFound, "User not found!")
	} else if err != nil {
		return err
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
			return err
		}

		return c.JSON(fiber.Map{"message": "User logged in successfully!",
			"token": t,
		})

	}

	return errors.NewErrors(fiber.StatusUnauthorized, "Check password!")
}
