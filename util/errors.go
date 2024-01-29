package util

import "errors"

var ErrRepeatedEntity = errors.New("There's already an equal entity registered")

var ErrEmptySelection = errors.New("The query retrieved nothing")

var ErrNoCorrespondingUser = errors.New("There's no user corresponding to this action")

var ErrNoCorrespondingProfile = errors.New("There's no profile corresponding to this action")

var ErrNoCorrespondingProfileOrPost = errors.New("There's no profile or post corresponding to this action")

var ErrNoCorrespondingProfileOrComment = errors.New("There's no profile or comment corresponding to this action")

var ErrIncorrectParameters = errors.New("The parameters provided were incorrect or incomplete")

var ErrPasswordNotGenerated = errors.New("The password couldn't be hashed")
