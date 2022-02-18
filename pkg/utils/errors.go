package utils

import (
	"errors"
)

// Public errors for Authority
var (
	ErrRoleCreatedAlready = errors.New("authority: the role has created already")
	ErrPermCreatedAlready = errors.New("authority: the permission has created already")
	ErrRoleNotFound       = errors.New("authority: the role not found")
	ErrUserNotFound       = errors.New("authority: the user not found")
	ErrPermNotFound       = errors.New("authority: the permission(s) not found")
)
