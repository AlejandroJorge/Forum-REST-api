package service

import "errors"

var ErrUnknown = errors.New("Unregistered service error")

var ErrIncorrectParameters = errors.New("Parameters provided aren't valid")

var ErrPasswordUnableToHash = errors.New("Couldn't hash password")

var ErrExistingEmail = errors.New("The email provided is already registered")

var ErrNotExistingEntity = errors.New("The entity doesn't exist")

var ErrNotValidCredentials = errors.New("Credentials are invalid")
