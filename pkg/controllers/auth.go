package controllers

import (
	"errors"
	"time"

	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/globals"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	res := globals.App.DB.Gorm.Create(&models.Users{Email: u.Email,
		Password: string(password),
		Name:     u.Name,
		Phone:    u.Phone,
		City:     u.City,
		State:    u.State,
		Country:  u.Country,
		ZipCode:  u.ZipCode,
		Address:  u.Adress,
	})

	if res.Error != nil {
		return utils.ReturnErrorMessage(c, res.Error)
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully!",
		"user":    u,
	})
}

func (AuthController) Login(c *fiber.Ctx) error {
	var user models.Users
	u := new(LoginRequest)
	utils.ParseBody(c, u)

	validate := utils.ValidateStruct(*u)
	if validate != nil {
		return c.Status(fiber.StatusForbidden).JSON(validate)
	}

	// Check exists
	res := globals.App.DB.Gorm.Where("email = ?", u.Email).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return utils.ReturnErrorMessage(c, errors.New("user not found"), fiber.StatusNotFound)
		}
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err == nil {
		claims := jwt.MapClaims{
			"fields": user,
			"exp":    time.Now().Add(time.Hour * globals.App.Config.Middleware.Jwt.Hours).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(globals.App.Config.Middleware.Jwt.Key))
		if err != nil {
			return utils.ReturnErrorMessage(c, err)
		}

		return c.JSON(fiber.Map{"message": "User logged in successfully!",
			"token": t,
		})

	}

	return utils.ReturnErrorMessage(c, errors.New("check password"), fiber.StatusUnauthorized)
}
