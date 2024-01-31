package repository

import "errors"

var ErrUnknown = errors.New("Unregistered repository error")

var ErrRepeatedEntity = errors.New("There's already an entity with a unique attribute")

var ErrNoRowsAffected = errors.New("The action affected no rows")

var ErrEmptySelection = errors.New("The query retrieved no rows")

var ErrNoMatchingDependency = errors.New("The entity required doens't exist")
