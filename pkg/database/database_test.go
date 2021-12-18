package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key   = "foo"
	value = []byte("bar")

	testStore = Init()
)

func init() {
	testStore.SetupRedis("redis://127.0.0.1:6379/", false)
}

func Test_Redis_Set(t *testing.T) {
        t.Parallel()
	assert := assert.New(t)

	err := testStore.Redis.Set(key, value, 0)
	assert.Equal(err, nil)
}

func Test_Redis_Get(t *testing.T) {
        t.Parallel()
	assert := assert.New(t)

	err := testStore.Redis.Set(key, value, 0)
	assert.Equal(err, nil)

	data, err := testStore.Redis.Get(key)
	assert.Equal(err, nil)
	assert.Equal(data, value)
}

func Test_Redis_Delete(t *testing.T) {
        t.Parallel()
	assert := assert.New(t)

	err := testStore.Redis.Set(key, value, 0)
	assert.Equal(err, nil)

	err = testStore.Redis.Delete(key)
	assert.Equal(err, nil)

	data, err := testStore.Redis.Get(key)
	assert.Equal(err, nil)
	assert.Equal(true, len(data) == 0)
}

func Test_Redis_NotExists(t *testing.T) {
        t.Parallel()
	assert := assert.New(t)

	data, err := testStore.Redis.Get(key)
	assert.Equal(err, nil)
	assert.Equal(true, len(data) == 0)
}

func Test_Redis_Close(t *testing.T) {
        t.Parallel()
	assert := assert.New(t)

	assert.Equal(testStore.Redis.Close(), nil)
}
