package service

import (
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repo domain.UserRepository
}

func (serv userServiceImpl) CreateNew(createInfo struct {
	NewEmail    string
	NewPassword string
}) (uint, error) {
	if !util.IsEmailFormat(createInfo.NewEmail) || createInfo.NewPassword == "" {
		return 0, util.ErrIncorrectParameters
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createInfo.NewPassword), 10)
	if err != nil {
		return 0, util.ErrPasswordNotGenerated
	}

	newID, err := serv.repo.CreateNew(domain.User{
		Email:          createInfo.NewEmail,
		HashedPassword: string(hashedPassword),
	})
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (serv userServiceImpl) Delete(id uint) error {
	return serv.repo.Delete(id)
}

func (serv userServiceImpl) GetByEmail(email string) (domain.User, error) {
	if !util.IsEmailFormat(email) {
		return domain.User{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByEmail(email)
}

func (serv userServiceImpl) GetByID(id uint) (domain.User, error) {
	if id == 0 {
		return domain.User{}, util.ErrIncorrectParameters
	}

	return serv.repo.GetByID(id)
}

func (serv userServiceImpl) Update(id uint, updateInfo struct {
	UpdatedEmail    string
	UpdatedPassword string
}) error {
	if id == 0 {
		return util.ErrIncorrectParameters
	}

	if util.IsEmailFormat(updateInfo.UpdatedEmail) {
		err := serv.repo.UpdateEmail(id, updateInfo.UpdatedEmail)
		if err != nil {
			return err
		}
	}

	if updateInfo.UpdatedPassword != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(updateInfo.UpdatedPassword), 10)
		if err != nil {
			return util.ErrPasswordNotGenerated
		}

		err = serv.repo.UpdateHashedPassword(id, string(hashed))
		if err != nil {
			return err
		}
	}

	return nil
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return userServiceImpl{repo: repo}
}
