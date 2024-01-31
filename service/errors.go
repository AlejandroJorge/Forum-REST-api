package service

import "errors"

var ErrUnknown = errors.New("Unregistered service error")

var ErrIncorrectParameters = errors.New("Parameters provided aren't valid")

var ErrPasswordUnableToHash = errors.New("Couldn't hash password")

var ErrExistingEmail = errors.New("The email provided is already registered")

var ErrAlreadyExisting = errors.New("The exact same entity already exists")

var ErrNotExistingEntity = errors.New("The entity doesn't exist")

var ErrNotValidCredentials = errors.New("Credentials are invalid")

var ErrDependencyNotSatisfied = errors.New("Dependency couldn't be satisfied")

var ErrProfileExistsOrTagNameIsRepeated = errors.New("Profile for this user already exists or tagname is already registered")
