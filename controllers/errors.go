package controllers

import (
	"errors"
)

var (
	ErrFieldNotFound = errors.New("field not found in Secret")
)
