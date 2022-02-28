package storage

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
)

var (
	ErrMissingFile     = errors.NewErrors(fiber.StatusBadRequest, "You must enter file.")
	ErrUnvalidMIMEType = errors.NewErrors(fiber.StatusForbidden, "You must enter valid file. Please check extension!")
	ErrMaxUploadSize   = errors.NewErrors(fiber.StatusForbidden, "You've reached max file size!")
)

type FileOpts struct {
	FormName string
	SavePath string
	Width    int
	Height   int
	Required bool
	DoFunc   func()
}

func UploadFile(c *fiber.Ctx, opts FileOpts) error {
	// Get Multipart file
	mf, err := c.MultipartForm()
	if err != nil {
		return err
	}
	if mf.File == nil {
		return err
	}
	fileheader := mf.File[opts.FormName]
	if fileheader == nil && opts.Required {
		return ErrMissingFile
	} else if fileheader == nil {
		return nil
	}

	// Check MIME type
	if mime := IsValidMIME(filepath.Ext(fileheader[0].Filename)); !mime {
		return ErrUnvalidMIMEType
	}

	// Check size
	maxUploadSize := api.App.Config.App.Files.MaxSize * 1024 * 1024
	if fileheader[0].Size > maxUploadSize {
		return ErrMaxUploadSize
	}

	// Convert WEBP & Resize Image
	file, err := fileheader[0].Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	buffer, err = bimg.NewImage(buffer).Process(bimg.Options{
		Width:    opts.Width,
		Height:   opts.Height,
		Type:     bimg.WEBP,
		Lossless: true,
		Quality:  80,
	})
	if err != nil {
		return err
	}

	// Save file to storage
	if err := api.App.DB.S3.Set(opts.SavePath+".webp", buffer, 0); err != nil {
		return err
	}

	// Do custom user functions
	opts.DoFunc()

	return nil
}

func UpdateAvatar(c *fiber.Ctx, uu *ent.UserUpdateOne, name string, removeAvatar bool) error {
	time := fmt.Sprint(time.Now().Unix())
	if removeAvatar {
		removeOldAvatar(uu)
		uu.ClearAvatar()
	} else {
		removeOldAvatar(uu)
		err := UploadFile(c, FileOpts{
			FormName: "avatar",
			SavePath: "avatars/" + name + "-" + time + "-avatar",
			Width:    256,
			Height:   256,
			DoFunc: func() {
				uu.SetAvatar(name + "-" + time + "-avatar.webp")
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func removeOldAvatar(uu *ent.UserUpdateOne) error {
	avatar, err := uu.Mutation().OldAvatar(context.Background())
	if err != nil {
		return err
	}

	_ = api.App.DB.S3.Delete("avatars/" + avatar)
	return nil
}
