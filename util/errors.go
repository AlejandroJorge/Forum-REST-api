package util

import "errors"

var ErrRepeatedEntity = errors.New("There's already an equal entity registered")

var ErrEmptySelection = errors.New("The query retrieved nothing")

var ErrNoCorrespondingUser = errors.New("There's no user corresponding to this action")
