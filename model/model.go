package model

import (
	"errors"
)

const ApplicationName = "stuff"

const DateFormat = "2006-01-02"
const DateTimeFormat = "2000-01-15T00:00:00Z"

var ErrNoContent = errors.New("no rows in result set")
