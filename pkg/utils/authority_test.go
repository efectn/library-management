package utils

import (
	"context"
	"testing"

	"github.com/efectn/library-management/pkg/database"
	"github.com/efectn/library-management/pkg/database/ent/enttest"
	"github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/webserver"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var authority = Authority{}
var connString = "host=localhost port=5432 user=postgres dbname=library_management password=postgres sslmode=disable"

func init() {
	api.App = new(webserver.AppSkel)
	api.App.DB = new(database.Database)
}

func Test_CreateRole(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeleteRole("test")

	_, err := authority.CreateRole("test")
	assert.NoError(t, err)

	_, err = authority.CreateRole("test")
	assert.Equal(t, err.Error(), "authority: the role has created already")

	authority.DeleteRole("test")
}

func Test_CreatePermission(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeletePermission("test")

	_, err := authority.CreatePermission("test")
	assert.NoError(t, err)

	_, err = authority.CreatePermission("test")
	assert.Equal(t, err.Error(), "authority: the permission has created already")

	authority.DeletePermission("test")
}

func Test_DeleteRole(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeleteRole("test")

	_, err := authority.CreateRole("test")
	assert.NoError(t, err)

	err = authority.DeleteRole("test")
	assert.NoError(t, err)

	err = authority.DeleteRole("test123456")
	assert.Equal(t, err.Error(), "authority: the role not found")
}

func Test_DeletePermission(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeletePermission("test")

	_, err := authority.CreatePermission("test")
	assert.NoError(t, err)

	err = authority.DeletePermission("test")
	assert.NoError(t, err)

	err = authority.DeletePermission("test123456")
	assert.Equal(t, err.Error(), "authority: the permission(s) not found")
}

func Test_AssignPermission(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeletePermission("test")
	authority.DeleteRole("test")

	_, err := authority.CreatePermission("test")
	assert.NoError(t, err)

	err = authority.AssignPermissions("test", "test")
	assert.Equal(t, err.Error(), "authority: the role not found")

	_, err = authority.CreateRole("test")
	assert.NoError(t, err)

	err = authority.AssignPermissions("test", "test")
	assert.NoError(t, err)

	authority.DeletePermission("test")
	authority.DeleteRole("test")
}

func Test_AssignRole_CheckRole(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())
	authority.DeleteRole("test")
	authority.DeleteRole("test2")

	_, err := authority.CreateRole("test")
	assert.NoError(t, err)

	_, err = authority.CreateRole("test2")
	assert.NoError(t, err)

	err = authority.AssignRole(1200000, "test")
	assert.Equal(t, err.Error(), "authority: the user not found")

	u, err := api.App.DB.Ent.User.Create().SetName("test").SetPassword("123").SetEmail("-").Save(context.Background())
	assert.NoError(t, err)

	err = authority.AssignRole(u.ID, "test")
	assert.NoError(t, err)

	exists, err := authority.CheckRole(u.ID, "test")
	assert.Equal(t, exists, true)
	assert.NoError(t, err)

	exists, err = authority.CheckRole(u.ID, "test2")
	assert.Equal(t, exists, false)
	assert.NoError(t, err)

	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())

	exists, err = authority.CheckRole(u.ID, "test")
	assert.Equal(t, exists, false)
	assert.Equal(t, err.Error(), "authority: the user not found")

	authority.DeleteRole("test")
}

func Test_CheckPermission_CheckRolePermission(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeletePermission("test")
	authority.DeletePermission("test2")
	authority.DeleteRole("test")
	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())

	u, _ := api.App.DB.Ent.User.Create().SetName("test").SetPassword("123").SetEmail("-").Save(context.Background())
	authority.CreateRole("test")
	authority.CreatePermission("test")
	authority.CreatePermission("test2")

	authority.AssignPermissions("test", "test")
	authority.AssignRole(u.ID, "test")

	exists, err := authority.CheckPermission(u.ID, "test")
	assert.Equal(t, exists, true)
	assert.NoError(t, err)

	exists, err = authority.CheckPermission(u.ID, "test2")
	assert.Equal(t, exists, false)
	assert.NoError(t, err)

	exists, err = authority.CheckRolePermission("test", "test")
	assert.Equal(t, exists, true)
	assert.NoError(t, err)

	exists, err = authority.CheckRolePermission("test", "test2")
	assert.Equal(t, exists, false)
	assert.NoError(t, err)

	authority.DeletePermission("test")
	authority.DeletePermission("test2")
	authority.DeleteRole("test")
	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())
}

func Test_RevokeRole_RevokePermission(t *testing.T) {
	api.App.DB.Ent = enttest.Open(t, "postgres", connString)

	authority.DeletePermission("test")
	authority.DeleteRole("test")
	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())

	u, _ := api.App.DB.Ent.User.Create().SetName("test").SetPassword("123").SetEmail("-").Save(context.Background())
	authority.CreateRole("test")
	authority.CreatePermission("test")

	authority.AssignPermissions("test", "test")
	authority.AssignRole(u.ID, "test")

	exists, err := authority.CheckRole(u.ID, "test")
	assert.Equal(t, exists, true)
	assert.NoError(t, err)

	err = authority.RevokeRole(u.ID, "test")
	assert.NoError(t, err)

	exists, err = authority.CheckRole(u.ID, "test")
	assert.Equal(t, exists, false)
	assert.NoError(t, err)

	exists, err = authority.CheckRolePermission("test", "test")
	assert.Equal(t, exists, true)
	assert.NoError(t, err)

	err = authority.RevokePermission("test", "test")
	assert.NoError(t, err)

	exists, err = authority.CheckRolePermission("test", "test")
	assert.Equal(t, exists, false)
	assert.NoError(t, err)

	authority.DeletePermission("test")
	authority.DeleteRole("test")
	api.App.DB.Ent.User.Delete().Where(user.NameEQ("test")).Exec(context.Background())
}
