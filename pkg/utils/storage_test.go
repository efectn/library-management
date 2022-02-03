package utils

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory"
	"github.com/stretchr/testify/assert"
)

func Test_Ctx_SaveFileToStorage(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	storage := memory.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		fh, err := c.FormFile("file")
		assert.Equal(t, nil, err)

		err = SaveFileToStorage(fh, "test", storage)
		assert.Equal(t, nil, err)

		file, err := storage.Get("test")
		assert.Equal(t, []byte("hello world"), file)
		assert.Equal(t, nil, err)

		err = storage.Delete("test")
		assert.Equal(t, nil, err)

		return nil
	})

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	ioWriter, err := writer.CreateFormFile("file", "test")
	assert.Equal(t, nil, err)

	_, err = ioWriter.Write([]byte("hello world"))
	assert.Equal(t, nil, err)
	writer.Close()

	req := httptest.NewRequest(fiber.MethodPost, "/test", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Length", strconv.Itoa(len(body.Bytes())))

	resp, err := app.Test(req)
	assert.Equal(t, nil, err, "app.Test(req)")
	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Status code")
}
