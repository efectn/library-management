package utils

import (
	"testing"

	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/webserver"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type exampleRequest struct {
	City string `validate:"required,min=3"`
	Age  uint   `validate:"min=18"`
}

type demo struct {
	Name string `form:"name"`
}

func init() {
	api.App = new(webserver.AppSkel)
	api.App.Validator = validator.New()
}

func Test_ValidateStruct(t *testing.T) {
	t.Parallel()

	req := new(exampleRequest)
	req.Age = 18
	req.City = "Karabük"

	resp := ValidateStruct(req)
	assert.Empty(t, resp)

	req = new(exampleRequest)
	req.Age = 16

	resp = ValidateStruct(req)
	assert.NotEmpty(t, resp)
}

func Benchmark_ValidateStruct(b *testing.B) {
	var resp []*errorResponse

	req := new(exampleRequest)
	req.Age = 18

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resp = ValidateStruct(*req)
	}

	assert.NotEmpty(b, resp)
}

func Test_ParseBody(t *testing.T) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	body := []byte("name=john")

	c.Request().SetBody(body)
	c.Request().Header.SetContentType(fiber.MIMEApplicationForm)
	c.Request().Header.SetContentLength(len(body))

	d := new(demo)
	err := ParseBody(c, d)

	assert.Empty(t, err)
	assert.Equal(t, "john", d.Name)
}

func Benchmark_ParseBody(b *testing.B) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	body := []byte("name=john")

	c.Request().SetBody(body)
	c.Request().Header.SetContentType(fiber.MIMEApplicationForm)
	c.Request().Header.SetContentLength(len(body))

	d := new(demo)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ParseBody(c, d)
	}

	assert.Empty(b, ParseBody(c, d))
	assert.Equal(b, "john", d.Name)
}
