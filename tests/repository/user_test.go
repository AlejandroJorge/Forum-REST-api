package repository

import (
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/tests"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestUserCreate(t *testing.T) {
	repo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())

	newUsers := []domain.User{
		{Email: "asdasda@gmail.com", HashedPassword: "1AS56D1AS6D51ASD6"},
		{Email: "awoidjaojd@asd.com", HashedPassword: "A84SD1A35SD1"},
		{Email: "email@5a6sd.pe", HashedPassword: "48AW74965S4D16A5"},
	}

	for _, newUser := range newUsers {
		email := newUser.Email
		id, err := repo.CreateNew(newUser)
		tests.EndTestIfError(err, t)

		retrievedUser, err := repo.GetByEmail(email)
		tests.EndTestIfError(err, t)

		tests.AssertEqu(newUser.Email, retrievedUser.Email, t)
		tests.AssertEqu(newUser.HashedPassword, retrievedUser.HashedPassword, t)

		retrievedUser, err = repo.GetByID(id)
		tests.EndTestIfError(err, t)

		tests.AssertEqu(newUser.Email, retrievedUser.Email, t)
		tests.AssertEqu(newUser.HashedPassword, retrievedUser.HashedPassword, t)
	}

	_, err := repo.GetByEmail("unexistentemail@gmail.com")
	tests.AssertEqu(util.ErrEmptySelection, err, t)
}

func TestUserCreateDuplicated(t *testing.T) {
	repo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "sameemail@gmail.com", HashedPassword: "A1S5DA1"}

	_, err := repo.CreateNew(newUser)
	tests.EndTestIfError(err, t)

	_, err = repo.CreateNew(newUser)
	tests.AssertEqu(util.ErrRepeatedEntity, err, t)
}

func TestUserUpdate(t *testing.T) {
	repo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "myfirstemail@gmail.com", HashedPassword: "ASD51A6S165ASD"}

	id, err := repo.CreateNew(newUser)
	tests.EndTestIfError(err, t)

	updatedEmail := "mysecondemail@gmail.com"
	err = repo.UpdateEmail(id, updatedEmail)
	tests.EndTestIfError(err, t)

	retrievedUser, err := repo.GetByID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedEmail, retrievedUser.Email, t)

	updatedHashedPassword := "AS1D56AS1D65AS1D"
	err = repo.UpdateHashedPassword(id, updatedHashedPassword)
	tests.EndTestIfError(err, t)

	retrievedUser, err = repo.GetByID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedHashedPassword, retrievedUser.HashedPassword, t)
}

func TestUserDelete(t *testing.T) {
	repo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "myfirstemail@gmail.com", HashedPassword: "ASD51A6S165ASD"}

	id, err := repo.CreateNew(newUser)
	tests.EndTestIfError(err, t)

	err = repo.Delete(id)
	tests.EndTestIfError(err, t)

	_, err = repo.GetByID(id)
	tests.AssertEqu(util.ErrEmptySelection, err, t)
}
