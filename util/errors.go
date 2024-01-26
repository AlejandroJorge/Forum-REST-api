package util

import "errors"

var ErrRepeatedEntity = errors.New("There's already an equal entity registered")

var ErrEmptySelection = errors.New("The query retrieved nothing")
