// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldZipCode holds the string denoting the zip_code field in the database.
	FieldZipCode = "zip_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "user_roles"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "roles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldName,
	FieldAvatar,
	FieldPhone,
	FieldCity,
	FieldState,
	FieldCountry,
	FieldAddress,
	FieldZipCode,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"user_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/efectn/library-management/pkg/database/ent/runtime"
//
var (
	Hooks [2]ent.Hook
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
