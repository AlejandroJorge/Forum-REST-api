package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repo domain.UserRepository
}

// Returns the ID of the created user, can return ErrIncorrectParameters, ErrPasswordUnableToHash, ErrExistingEmail
func (serv userServiceImpl) Create(email, password string) (uint, error) {
	if !util.IsEmailFormat(email) || password == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return 0, ErrIncorrectParameters
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logging.LogDomainError(ErrPasswordUnableToHash)
		return 0, ErrPasswordUnableToHash
	}

	newID, err := serv.repo.Create(email, string(hashedPassword))
	if err == repository.ErrRepeatedEntity {
		logging.LogDomainError(ErrExistingEmail)
		return 0, ErrExistingEmail
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return 0, ErrUnknown
	}

	return newID, nil
}

// Can return ErrNotExistingEntity
func (serv userServiceImpl) Delete(id uint) error {
	err := serv.repo.Delete(id)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Returns a valid user, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv userServiceImpl) GetByEmail(email string) (domain.User, error) {
	if !util.IsEmailFormat(email) {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.User{}, ErrIncorrectParameters
	}

	user, err := serv.repo.GetByEmail(email)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.User{}, ErrNotExistingEntity
	}
	if err != nil {
		return domain.User{}, ErrUnknown
	}

	return user, nil
}

// Returns a valid user, can return ErrIncorrectParameters, ErrNotExistingEntity
func (serv userServiceImpl) GetByID(id uint) (domain.User, error) {
	if id == 0 {
		logging.LogDomainError(ErrIncorrectParameters)
		return domain.User{}, ErrIncorrectParameters
	}

	user, err := serv.repo.GetByID(id)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return domain.User{}, ErrNotExistingEntity
	}
	if err != nil {
		return domain.User{}, ErrUnknown
	}

	return user, nil
}

// Can return ErrNotExistingEntity, ErrIncorrectParameters
func (serv userServiceImpl) UpdateEmail(id uint, email string) error {
	if id == 0 || !util.IsEmailFormat(email) {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	err := serv.repo.UpdateEmail(id, email)
	if err == repository.ErrNoRowsAffected {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		return ErrUnknown
	}

	return nil
}

// Can return ErrNotExistingEntity, ErrIncorrectParameters, ErrPasswordUnableToHash
func (serv userServiceImpl) UpdatePassword(id uint, password string) error {
	if id == 0 || password == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logging.LogDomainError(ErrPasswordUnableToHash)
		return ErrPasswordUnableToHash
	}

	err = serv.repo.UpdateHashedPassword(id, string(hashed))
	if err == repository.ErrNoRowsAffected {
		logging.LogUnexpectedDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogUnexpectedDomainError(err)
		return ErrUnknown
	}

	return nil
}

// Returns nil if credentials are OK, can return ErrIncorrectParameters, ErrNotValidCredentials, ErrNotExistingEntity
func (serv userServiceImpl) CheckCredentials(email, password string) error {
	if !util.IsEmailFormat(email) || password == "" {
		logging.LogDomainError(ErrIncorrectParameters)
		return ErrIncorrectParameters
	}

	user, err := serv.repo.GetByEmail(email)
	if err == repository.ErrEmptySelection {
		logging.LogDomainError(ErrNotExistingEntity)
		return ErrNotExistingEntity
	}
	if err != nil {
		logging.LogDomainError(err)
		return ErrUnknown
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		logging.LogDomainError(ErrNotValidCredentials)
		return ErrNotValidCredentials
	}

	return nil
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return userServiceImpl{repo: repo}
}
