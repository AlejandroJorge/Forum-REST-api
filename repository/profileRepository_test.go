package repository

import (
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestProfileCreateWithSameID(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "someemail@gmail.com", HashedPassword: "A5S1D6"},
	)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "SomeUser", TagName: "someone"})
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "AnotherUser", TagName: "anotherone"})
	if err != util.ErrRepeatedEntity {
		t.Errorf("Expected '%s', got '%s'", util.ErrRepeatedEntity, err)
	}
}

func TestProfileCreateWithNullID(t *testing.T) {
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())

	_, err := profileRepo.CreateNew(domain.Profile{UserID: 0, DisplayName: "NullUser", TagName: "nullone"})
	if err != util.ErrNoCorrespondingUser {
		t.Errorf("Expected '%s', got '%s'", util.ErrNoCorrespondingUser, err)
	}
}

func TestProfileCreateWithSameTagName(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())

	idFirst, err := userRepo.CreateNew(domain.User{
		Email: "a1s86d5a1sd@gmail.com", HashedPassword: "A5S1D6"},
	)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: idFirst, DisplayName: "Some name", TagName: "RepeatedTagName",
	})
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	idSecond, err := userRepo.CreateNew(domain.User{
		Email: "78aw5a61dw@gmail.com", HashedPassword: "A5S1D6"},
	)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: idSecond, DisplayName: "Random stuff", TagName: "RepeatedTagName",
	})
	if err != util.ErrRepeatedEntity {
		t.Errorf("Expected '%s', got '%s'", util.ErrRepeatedEntity, err)
	}
}

func TestProfileUpdateDisplayName(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "randomemail@gmail.com", HashedPassword: "A5W4da15S361AD56WD1"},
	)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "RandomUser", TagName: "randomone"})
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	updatedDisplayName := "newDisplayName"
	err = profileRepo.UpdateDisplayName(id, updatedDisplayName)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	retrievedProfile, err := profileRepo.GetByUserID(id)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	if retrievedProfile.DisplayName != updatedDisplayName {
		t.Errorf("Expected '%s', got '%s'", updatedDisplayName, retrievedProfile.DisplayName)
	}

	updatedTagName := "newTagName"
	err = profileRepo.UpdateTagName(id, updatedTagName)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	retrievedProfile, err = profileRepo.GetByUserID(id)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	if retrievedProfile.TagName != updatedTagName {
		t.Errorf("Expected '%s', got '%s'", updatedTagName, retrievedProfile.TagName)
	}

	updatedPicturePath := "https://somepage.images.com/5668"
	err = profileRepo.UpdatePicturePath(id, updatedPicturePath)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	retrievedProfile, err = profileRepo.GetByUserID(id)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	if retrievedProfile.PicturePath != updatedPicturePath {
		t.Errorf("Expected '%s', got '%s'", updatedPicturePath, retrievedProfile.PicturePath)
	}

	updatedBackgroundPath := "https://somepage.images.com/87983"
	err = profileRepo.UpdateBackgroundPath(id, updatedBackgroundPath)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	retrievedProfile, err = profileRepo.GetByUserID(id)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	if retrievedProfile.BackgroundPath != updatedBackgroundPath {
		t.Errorf("Expected '%s', got '%s'", updatedBackgroundPath, retrievedProfile.BackgroundPath)
	}
}

func TestProfileDelete(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "fordeleting@gmail.com", HashedPassword: "A5W4da15S361"},
	)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "DyingUser", TagName: "ForDeleting"})
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	err = profileRepo.Delete(id)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	_, err = profileRepo.GetByUserID(id)
	if err != util.ErrEmptySelection {
		t.Errorf("Expected '%s', got '%s'", util.ErrEmptySelection, err)
	}
}
