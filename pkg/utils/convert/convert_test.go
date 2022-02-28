package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// From https://github.com/gofiber/fiber/blob/master/utils/convert_test.go#L9
func Test_UnsafeString(t *testing.T) {
	t.Parallel()

	out := UnsafeString([]byte("Hello, World!"))
	assert.Equal(t, "Hello, World!", out)
}

// From https://github.com/gofiber/fiber/blob/master/utils/convert_test.go#L17
func Benchmark_UnsafeString(b *testing.B) {
	hello := []byte("Hello, World!")
	var res string

	b.Run("unsafe", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = UnsafeString(hello)
		}
		assert.Equal(b, "Hello, World!", res)
	})

	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = string(hello)
		}
		assert.Equal(b, "Hello, World!", res)
	})
}

func Test_UnsafeBytes(t *testing.T) {
	t.Parallel()
	res := UnsafeBytes("Hello, World!")
	assert.Equal(t, []byte("Hello, World!"), res)
}

// go test -v -run=^$ -bench=UnsafeBytes -benchmem -count=4

func Benchmark_UnsafeBytes(b *testing.B) {
	hello := "Hello, World!"
	var res []byte
	b.Run("unsafe", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = UnsafeBytes(hello)
		}
		assert.Equal(b, []byte("Hello, World!"), res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = []byte(hello)
		}
		assert.Equal(b, []byte("Hello, World!"), res)
	})
}
