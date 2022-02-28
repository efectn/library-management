package admin

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/efectn/library-management/pkg/database/ent"
	euser "github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils"
	"github.com/efectn/library-management/pkg/utils/convert"
	"github.com/efectn/library-management/pkg/utils/database"
	"github.com/efectn/library-management/pkg/utils/errors"
	"github.com/efectn/library-management/pkg/utils/storage"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

type CreateUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
	Name     string `validate:"required,min=3,max=32" json:"name"`
	Phone    string `validate:"omitempty,e164" json:"phone,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
	ZipCode  int    `validate:"number" form:"zip_code" json:"zip_code"`
	Address  string `json:"address,omitempty"`
	RoleIDs  []int  `form:"role_id" json:"role_id,omitempty"`
}

type UpdateUserRequest struct {
	Email        string `validate:"omitempty,email" json:"email"`
	Password     string `validate:"omitempty,min=8" json:"password"`
	Name         string `validate:"omitempty,min=3,max=32" json:"name"`
	Phone        string `validate:"omitempty,e164" json:"phone,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
	ZipCode      int    `validate:"number" form:"zip_code" json:"zip_code"`
	Address      string `json:"address,omitempty"`
	RoleIDs      []int  `form:"role_id" json:"role_id,omitempty"`
	RemoveAvatar bool   `form:"remove_avatar"`
}

func (UserController) Index(c *fiber.Ctx) error {
	users, err := api.App.DB.Ent.User.Query().
		WithRoles().
		Order(ent.Asc(euser.FieldID)).
		All(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User list has retrieved successfully!",
		"users":   users,
	})
}

func (UserController) Store(c *fiber.Ctx) error {
	u := new(CreateUserRequest)
	if err := utils.ParseAndValidate(c, u); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), api.App.Config.App.Hash.BcryptCost)
	if err != nil {
		return err
	}

	uc := api.App.DB.Ent.User.Create().SetEmail(u.Email).
		SetPassword(convert.UnsafeString(password)).
		SetName(u.Name).
		SetPhone(u.Phone).
		SetCity(u.City).
		SetState(u.State).
		SetCountry(u.Country).
		SetZipCode(u.ZipCode).
		SetAddress(u.Address).
		AddRoleIDs(u.RoleIDs...)

	time := fmt.Sprint(time.Now().Unix())
	err = storage.UploadFile(c, storage.FileOpts{
		FormName: "avatar",
		SavePath: "avatars/" + u.Name + "-" + time + "-avatar",
		Width:    256,
		Height:   256,
		DoFunc: func() {
			uc.SetAvatar(u.Name + "-" + time + "-avatar.webp")
		},
	})
	if err != nil {
		return err
	}

	// Create user
	user, err := uc.Save(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		// Remove created avatar
		_ = api.App.DB.S3.Delete("avatars/" + u.Name + "-" + time + "-avatar.webp")

		return err
	}

	return c.JSON(fiber.Map{
		"message": "The user created successfully!",
		"user":    user,
	})
}

func (UserController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	user, err := api.App.DB.Ent.User.Query().
		Where(euser.IDEQ(id)).
		WithRoles().
		First(context.Background())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The user retrieved successfully!",
		"role":    user,
	})

}

func (UserController) Update(c *fiber.Ctx) error {
	u := new(UpdateUserRequest)
	if err := utils.ParseAndValidate(c, u); err != nil {
		return errors.NewErrors(fiber.StatusForbidden, err)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	uu := api.App.DB.Ent.User.UpdateOneID(id)

	// Update fields if given
	if u.Email != "" {
		uu.SetEmail(u.Email)
	}

	if u.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(u.Password), api.App.Config.App.Hash.BcryptCost)
		if err != nil {
			return err
		}

		uu.SetPassword(convert.UnsafeString(password))
	}

	if u.Name != "" {
		uu.SetName(u.Name)
	}

	// Update optional fields
	uu.SetPhone(u.Phone).
		SetCity(u.City).
		SetState(u.State).
		SetCountry(u.Country).
		SetZipCode(u.ZipCode).
		SetAddress(u.Address)

	// Update roles
	if u.RoleIDs != nil {
		uu.ClearRoles().AddRoleIDs(u.RoleIDs...)
	}

	// Update & remove avatar
	time := fmt.Sprint(time.Now().Unix())
	if u.RemoveAvatar {
		avatar, err := uu.Mutation().OldAvatar(context.Background())
		if err != nil {
			return err
		}

		if err := api.App.DB.S3.Delete("avatars/" + avatar); err != nil {
			return err
		}

		uu.ClearAvatar()
	} else {
		err = storage.UploadFile(c, storage.FileOpts{
			FormName: "avatar",
			SavePath: "avatars/" + u.Name + "-" + time + "-avatar",
			Width:    256,
			Height:   256,
			DoFunc: func() {
				uu.SetAvatar(u.Name + "-" + time + "-avatar.webp")
			},
		})
		if err != nil {
			return err
		}
	}

	user, err := uu.Save(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The user updated successfully!",
		"role":    user,
	})
}

func (UserController) Destroy(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	tx, err := api.App.DB.Ent.Tx(context.Background())
	if err != nil {
		return err
	}

	err = tx.User.DeleteOneID(id).Exec(context.Background())
	if err = errors.HandleEntErrors(err); err != nil {
		return database.EntRollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "The user deleted successfully!",
	})
}
