package repository

import (
	"os"
	"path"
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestMain(t *testing.M) {
	fixWorkingDir()
	config.Initialize()
	t.Run()
}

func TestCreateAndRead(t *testing.T) {
	repo := NewSQLiteUserRepository(config.SQLiteDatabase())

	newUsers := []domain.User{
		{Email: "asdasda@gmail.com", HashedPassword: "1AS56D1AS6D51ASD6"},
		{Email: "awoidjaojd@asd.com", HashedPassword: "A84SD1A35SD1"},
		{Email: "email@5a6sd.pe", HashedPassword: "48AW74965S4D16A5"},
	}

	for _, newUser := range newUsers {
		email := newUser.Email
		id, err := repo.CreateNew(newUser)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		retrievedUser, err := repo.GetByEmail(email)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if retrievedUser.Email != newUser.Email {
			t.Errorf("Expected '%s', got '%s'", newUser.Email, retrievedUser.Email)
		}
		if retrievedUser.HashedPassword != newUser.HashedPassword {
			t.Errorf("Expected '%s', got '%s'", newUser.HashedPassword, retrievedUser.HashedPassword)
		}

		retrievedUser, err = repo.GetByID(id)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if retrievedUser.Email != newUser.Email {
			t.Errorf("Expected '%s', got '%s'", newUser.Email, retrievedUser.Email)
		}
		if retrievedUser.HashedPassword != newUser.HashedPassword {
			t.Errorf("Expected '%s', got '%s'", newUser.HashedPassword, retrievedUser.HashedPassword)
		}

	}

	_, err := repo.GetByEmail("unexistentemail@gmail.com")
	if err != util.ErrEmptySelection {
		t.Errorf("Expected '%s', got '%s'", util.ErrEmptySelection.Error(), err)
	}
}

func TestCreateDuplicated(t *testing.T) {
	repo := NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "sameemail@gmail.com", HashedPassword: "A1S5DA1"}

	_, err := repo.CreateNew(newUser)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	_, err = repo.CreateNew(newUser)
	if err != util.ErrRepeatedEntity {
		t.Errorf("Expected '%s', got '%s'", util.ErrRepeatedEntity, err)
	}

}

func TestUpdate(t *testing.T) {
	repo := NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "myfirstemail@gmail.com", HashedPassword: "ASD51A6S165ASD"}

	id, err := repo.CreateNew(newUser)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	updatedEmail := "mysecondemail@gmail.com"
	err = repo.UpdateEmail(id, updatedEmail)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	retrievedUser, err := repo.GetByID(id)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if retrievedUser.Email != updatedEmail {
		t.Errorf("Error updating user email with id %d, expected '%s', got '%s'", id, updatedEmail, retrievedUser.Email)
		t.Errorf("Expected '%s', got '%s'", updatedEmail, retrievedUser.Email)
	}

	updatedHashedPassword := "AS1D56AS1D65AS1D"
	err = repo.UpdateHashedPassword(id, updatedHashedPassword)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	retrievedUser, err = repo.GetByID(id)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if retrievedUser.HashedPassword != updatedHashedPassword {
		t.Errorf("Expected '%s', got '%s'", updatedHashedPassword, retrievedUser.HashedPassword)
	}
}

func TestDelete(t *testing.T) {
	repo := NewSQLiteUserRepository(config.SQLiteDatabase())

	newUser := domain.User{Email: "myfirstemail@gmail.com", HashedPassword: "ASD51A6S165ASD"}

	id, err := repo.CreateNew(newUser)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = repo.Delete(id)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	_, err = repo.GetByID(id)
	if err != util.ErrEmptySelection {
		t.Errorf("Expected '%s', got ''%s''", util.ErrEmptySelection, err)
	}
}

func fixWorkingDir() {
	currentDir, err := os.Getwd()
	util.PanicIfError(err)
	workingDir := path.Dir(currentDir)
	config.SetWorkingDir(workingDir)
}
