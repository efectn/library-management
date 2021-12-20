// Modified version of https://github.com/harranali/authority. Special thanks to @harranali.
// TODO: Add tests.
package utils

import (
	"errors"

	"github.com/efectn/library-management/pkg/database/models"
	"github.com/efectn/library-management/pkg/globals/api"
	"gorm.io/gorm"
)

// Authority helps deal with permissions
type Authority struct{}

var (
	ErrPermissionInUse     = errors.New("cannot delete assigned permission")
	ErrPermissionNotFound  = errors.New("permission not found")
	ErrRoleAlreadyAssigned = errors.New("this role is already assigned to the user")
	ErrRoleInUse           = errors.New("cannot delete assigned role")
	ErrRoleNotFound        = errors.New("role not found")
)

// CreateRole stores a role in the databaseADV
// it accepts the role name. it returns an error
// in case of any
func (Authority) CreateRole(roleName string) (models.Role, error) {
	var dbRole models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&dbRole)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// create
			api.App.DB.Gorm.Create(&models.Role{Name: roleName}).First(&dbRole)
			return dbRole, nil
		}
	}

	return dbRole, res.Error
}

// CreatePermission stores a permission in the database
// it accepts the permission name. it returns an error
// in case of any
func (Authority) CreatePermission(permName string) (models.Permission, error) {
	var dbPerm models.Permission
	res := api.App.DB.Gorm.Where("name = ?", permName).First(&dbPerm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// create
			api.App.DB.Gorm.Create(&models.Permission{Name: permName}).First(&dbPerm)
			return dbPerm, nil
		}
	}

	return dbPerm, res.Error
}

// DeleteRole deletes a given role
// if the role is assigned to a user it returns an error
func (Authority) DeleteRole(roleName string) error {
	// find the role
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// delete the role
	api.App.DB.Gorm.Where("name = ?", roleName).Select("Users").Delete(models.Role{})

	return nil
}

// DeletePermission deletes a given permission
// if the permission is assigned to a role it returns an error
func (Authority) DeletePermission(permName string) error {
	// find the permission
	var perm models.Permission
	res := api.App.DB.Gorm.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	// delete the permission
	if err := api.App.DB.Gorm.Delete(&perm); err.Error != nil {
		return err.Error
	}

	return nil
}

// AssignPermissions assigns a group of permissions to a given role
// it accepts in the first parameter the role name, it returns an error if there is not matching record
// of the role name in the database.
// the second parameter is a slice of strings which represents a group of permissions to be assigned to the role
// if any of these permissions doesn't have a matching record in the database the operations stops, changes reverted
// and error is returned
// in case of success nothing is returned
func (Authority) AssignPermissions(roleName string, permNames ...string) error {
	// get the role id
	var role models.Role
	rRes := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if rRes.Error != nil {
		if errors.Is(rRes.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	var perms []models.Permission
	// get the permissions ids
	for _, permName := range permNames {
		var perm models.Permission
		pRes := api.App.DB.Gorm.Where("name = ?", permName).First(&perm)
		if pRes.Error != nil {
			if errors.Is(pRes.Error, gorm.ErrRecordNotFound) {
				return ErrPermissionNotFound
			}

		}

		perms = append(perms, perm)
	}

	api.App.DB.Gorm.Model(&role).Association("Permissions").Append(&perms)

	return nil
}

// AssignRole assigns a given role to a user
// the first parameter is the user id, the second parameter is the role name
// if the role name doesn't have a matching record in the data base an error is returned
// if the user have already a role assigned to him an error is returned
func (Authority) AssignRole(userID uint, roleName string) error {
	// make sure the role exist
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}
	}

	// check if the role is already assigned
	var user models.Users
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return err.Error
	}

	if err := api.App.DB.Gorm.Model(&user).Association("Roles").Find(&role); err != nil {
		return ErrRoleAlreadyAssigned
	}

	// assign the role
	api.App.DB.Gorm.Model(&user).Association("Roles").Append(&role)

	return nil
}

// CheckRole checks if a role is assigned to a user
// it accepts the user id as the first parameter
// the role as the second parameter
// it returns an error if the role is not present in database
func (Authority) CheckRole(userID uint, roleName string) (bool, error) {
	// find the role
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrRoleNotFound
		}

	}

	// check if the role is already assigned
	var user models.Users
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return false, err.Error
	}

	var roles []models.Role
	if err := api.App.DB.Gorm.Model(&user).Association("Roles").Find(&roles); err != nil {
		return false, ErrRoleAlreadyAssigned
	}

	for _, v := range roles {
		if v.ID == role.ID {
			return true, nil
		}
	}

	return false, nil
}

// CheckPermission checks if a permission is assigned to the role that's assigned to the user.
// it accepts the user id as the first parameter
// the permission as the second parameter
// it returns an error if the permission is not present in the database
func (Authority) CheckPermission(userID uint, permName string) (bool, error) {
	// the user role
	var user models.Users
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return false, err.Error
	}

	// the permission
	var perm models.Permission
	if err := api.App.DB.Gorm.Where("name = ?", permName).Find(&perm); err.Error != nil {
		return false, err.Error
	}

	// Get relations
	if err := api.App.DB.Gorm.Preload("Roles.Permissions").Find(&user); err.Error != nil {
		return false, err.Error
	}
	var perms []models.Permission
	for _, v := range user.Roles {
		perms = append(perms, v.Permissions...)
	}

	// Check exists
	for _, v := range perms {
		if v.Name == perm.Name {
			return true, nil
		}
	}

	return false, nil
}

// CheckRolePermission checks if a role has the permission assigned
// it accepts the role as the first parameter
// it accepts the permission as the second parameter
// it returns an error if the role is not present in database
// it returns an error if the permission is not present in database
func (Authority) CheckRolePermission(roleName string, permName string) (bool, error) {
	// find the role
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrRoleNotFound
		}

	}

	// find the permission
	var perm models.Permission
	res = api.App.DB.Gorm.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, ErrPermissionNotFound
		}

	}

	// find the rolePermission
	var perms []models.Permission
	err := api.App.DB.Gorm.Model(&role).Association("Permissions").Find(&perms)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

	}

	for _, v := range perms {
		if v.ID == perm.ID {
			return true, nil
		}
	}

	return false, nil
}

// RevokeRole revokes a user's role
// it returns a error in case of any
func (Authority) RevokeRole(userID uint, roleName string) error {
	// find the role
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// revoke the role
	var user models.Users
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return err.Error
	}

	api.App.DB.Gorm.Model(&user).Association("Roles").Delete(&role)

	return nil
}

// RevokePermission revokes a permission from the user's assigned role
// it returns an error in case of any
func (Authority) RevokePermission(userID uint, permName string) error {
	// revoke the permission from all roles of the user
	// find the user roles
	var user models.Users
	var roles []models.Role
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return err.Error
	}

	if err := api.App.DB.Gorm.Model(&user).Association("Roles").Find(&roles); err != nil {
		return ErrRoleAlreadyAssigned
	}

	// find the permission
	var perm models.Permission
	res := api.App.DB.Gorm.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	api.App.DB.Gorm.Model(&roles).Association("Permissions").Delete(&perm)

	return nil
}

// RevokeRolePermission revokes a permission from a given role
// it returns an error in case of any
func (Authority) RevokeRolePermission(roleName string, permName string) error {
	// find the role
	var role models.Role
	res := api.App.DB.Gorm.Where("name = ?", roleName).First(&role)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}

	}

	// find the permission
	var perm models.Permission
	res = api.App.DB.Gorm.Where("name = ?", permName).First(&perm)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}

	}

	// revoke the permission
	api.App.DB.Gorm.Model(&role).Association("Permissions").Delete(&perm)

	return nil
}

// GetRoles returns all stored roles
func (Authority) GetRoles() ([]string, error) {
	var result []string
	var roles []models.Role
	api.App.DB.Gorm.Find(&roles)

	for _, role := range roles {
		result = append(result, role.Name)
	}

	return result, nil
}

// GetUserRoles returns all user assigned roles
func (Authority) GetUserRoles(userID uint) ([]models.Role, error) {
	var user models.Users
	if err := api.App.DB.Gorm.First(&user, userID); err.Error != nil {
		return []models.Role{}, err.Error
	}

	var roles []models.Role
	if err := api.App.DB.Gorm.Model(&user).Association("Roles").Find(&roles); err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}

// GetPermissions returns all stored permissions
func (Authority) GetPermissions() ([]string, error) {
	var result []string
	var perms []models.Permission
	api.App.DB.Gorm.Find(&perms)

	for _, perm := range perms {
		result = append(result, perm.Name)
	}

	return result, nil
}
