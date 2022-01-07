// Package utils: Modified version of https://github.com/harranali/authority. Special thanks to @harranali.
// TODO: Add tests.
package utils

import (
	"context"

	"github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/database/ent/permission"
	"github.com/efectn/library-management/pkg/database/ent/role"
	"github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
)

// Authority helps deal with permissions
type Authority struct{}

// CreateRole stores a role in the databaseADV
// it accepts the role name. it returns an error
// in case of any
func (Authority) CreateRole(name string) (*ent.Role, error) {
	return api.App.DB.Ent.Role.Create().SetName(name).Save(context.Background())
}

// CreatePermission stores a permission in the database
// it accepts the permission name. it returns an error
// in case of any
func (Authority) CreatePermission(name string) (*ent.Permission, error) {
	return api.App.DB.Ent.Permission.Create().SetName(name).Save(context.Background())
}

// DeleteRole deletes a given role
// if the role is assigned to a user it returns an error
func (Authority) DeleteRole(name string) error {
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(name)).First(context.Background())
	if err != nil {
		return err
	}

	return api.App.DB.Ent.Role.DeleteOne(role).Exec(context.Background())
}

// DeletePermission deletes a given permission
// if the permission is assigned to a role it returns an error
func (Authority) DeletePermission(name string) error {
	perm, err := api.App.DB.Ent.Permission.Query().Where(permission.NameEQ(name)).First(context.Background())
	if err != nil {
		return err
	}

	return api.App.DB.Ent.Permission.DeleteOne(perm).Exec(context.Background())
}

// AssignPermissions assigns a group of permissions to a given role
// it accepts in the first parameter the role name, it returns an error if there is not matching record
// of the role name in the database.
// the second parameter is a slice of strings which represents a group of permissions to be assigned to the role
// if any of these permissions doesn't have a matching record in the database the operations stops, changes reverted
// and error is returned
// in case of success nothing is returned
func (Authority) AssignPermissions(roleName string, permNames ...string) error {
	// get the role
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(roleName)).First(context.Background())
	if err != nil {
		return err
	}

	// get the permission
	perms, err := api.App.DB.Ent.Permission.Query().Where(permission.NameIn(permNames...)).All(context.Background())
	if err != nil {
		return err
	}

	// assign permissions
	err = role.Update().AddPermissions(perms...).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// AssignRole assigns a given role to a user
// the first parameter is the user id, the second parameter is the role name
// if the role name doesn't have a matching record in the database an error is returned
// if the user have already a role assigned to him an error is returned
func (Authority) AssignRole(userID int, roleName string) error {
	// get the user
	user, err := api.App.DB.Ent.User.Query().Where(user.IDEQ(userID)).First(context.Background())
	if err != nil {
		return err
	}

	// make sure the role exist
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(roleName)).First(context.Background())
	if err != nil {
		return err
	}

	// assign the role
	err = user.Update().AddRoles(role).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// CheckRole checks if a role is assigned to a user
// it accepts the user id as the first parameter
// the role as the second parameter
// it returns an error if the role is not present in database
func (Authority) CheckRole(userID int, roleName string) (bool, error) {
	// get the user
	user, err := api.App.DB.Ent.User.Query().Where(user.IDEQ(userID)).First(context.Background())
	if err != nil {
		return false, err
	}

	return user.QueryRoles().Where(role.NameEQ(roleName)).Exist(context.Background())
}

// CheckPermission checks if a permission is assigned to the role that's assigned to the user.
// it accepts the user id as the first parameter
// the permission as the second parameter
// it returns an error if the permission is not present in the database
func (Authority) CheckPermission(userID int, permName string) (bool, error) {
	// get the user
	user, err := api.App.DB.Ent.User.Query().Where(user.IDEQ(userID)).WithRoles().First(context.Background())
	if err != nil {
		return false, err
	}

	return user.QueryRoles().QueryPermissions().Where(permission.NameEQ(permName)).Exist(context.Background())
}

// CheckRolePermission checks if a role has the permission assigned
// it accepts the role as the first parameter
// it accepts the permission as the second parameter
// it returns an error if the role is not present in database
// it returns an error if the permission is not present in database
func (Authority) CheckRolePermission(roleName string, permName string) (bool, error) {
	// get the user
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(roleName)).First(context.Background())
	if err != nil {
		return false, err
	}

	return role.QueryPermissions().Where(permission.NameEQ(permName)).Exist(context.Background())
}

// RevokeRole revokes a user's role
// it returns an error in case of any
func (Authority) RevokeRole(userID int, roleName string) error {
	// get the user
	user, err := api.App.DB.Ent.User.Query().Where(user.IDEQ(userID)).First(context.Background())
	if err != nil {
		return err
	}

	// make sure the role exist
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(roleName)).First(context.Background())
	if err != nil {
		return err
	}

	// assign the role
	err = user.Update().RemoveRoles(role).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// RevokePermission revokes a permission from the role
// it returns an error in case of any
func (Authority) RevokePermission(roleName string, permNames ...string) error {
	// get the role
	role, err := api.App.DB.Ent.Role.Query().Where(role.NameEQ(roleName)).First(context.Background())
	if err != nil {
		return err
	}

	// get the permission
	perms, err := api.App.DB.Ent.Permission.Query().Where(permission.NameIn(permNames...)).All(context.Background())
	if err != nil {
		return err
	}

	// assign permissions
	err = role.Update().RemovePermissions(perms...).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// GetRoles returns all stored roles
func (Authority) GetRoles() ([]*ent.Role, error) {
	return api.App.DB.Ent.Role.Query().All(context.Background())
}

// GetUserRoles returns all user assigned roles
func (Authority) GetUserRoles(userID int) ([]*ent.Role, error) {
	// get the user
	user, err := api.App.DB.Ent.User.Query().Where(user.IDEQ(userID)).First(context.Background())
	if err != nil {
		return []*ent.Role{}, err
	}

	return user.QueryRoles().All(context.Background())
}

// GetPermissions returns all stored permissions
func (Authority) GetPermissions() ([]*ent.Permission, error) {
	return api.App.DB.Ent.Permission.Query().All(context.Background())
}
