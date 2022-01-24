package utils

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/efectn/library-management/pkg/utils/convert"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func Test_ReturnErrorMessage(t *testing.T) {
	t.Parallel()

	app := fiber.New()

	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	// Check default status code
	app.Get("/", func(c *fiber.Ctx) error {
		return ReturnErrorMessage(c, "test")
	})

	body := &bytes.Buffer{}
	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/", body))

	assert.Equal(t, err, nil)
	assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)

	// Check custom status code
	app.Get("/c", func(c *fiber.Ctx) error {
		return ReturnErrorMessage(c, "test", fiber.StatusUnauthorized)
	})

	resp, err = app.Test(httptest.NewRequest(fiber.MethodGet, "/c", body))

	assert.Equal(t, err, nil)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	// Check error message
	ReturnErrorMessage(c, "test")
	assert.Equal(t, "{\"message\":\"test\"}", convert.UnsafeString(c.Response().Body()))
}
