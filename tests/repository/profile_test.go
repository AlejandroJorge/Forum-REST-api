package repository

import (
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/tests"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestProfileCreateWithSameID(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "someemail@gmail.com", HashedPassword: "A5S1D6"},
	)
	tests.EndTestIfError(err, t)

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "SomeUser", TagName: "someone"})
	tests.EndTestIfError(err, t)

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "AnotherUser", TagName: "anotherone"})
	tests.AssertEqu(util.ErrRepeatedEntity, err, t)
}

func TestProfileCreateWithNullID(t *testing.T) {
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	_, err := profileRepo.CreateNew(domain.Profile{UserID: 0, DisplayName: "NullUser", TagName: "nullone"})
	tests.AssertEqu(util.ErrNoCorrespondingUser, err, t)
}

func TestProfileCreateWithSameTagName(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	idFirst, err := userRepo.CreateNew(domain.User{
		Email: "a1s86d5a1sd@gmail.com", HashedPassword: "A5S1D6"},
	)
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: idFirst, DisplayName: "Some name", TagName: "RepeatedTagName",
	})
	tests.EndTestIfError(err, t)

	idSecond, err := userRepo.CreateNew(domain.User{
		Email: "78aw5a61dw@gmail.com", HashedPassword: "A5S1D6"},
	)
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: idSecond, DisplayName: "Random stuff", TagName: "RepeatedTagName",
	})
	tests.AssertEqu(util.ErrRepeatedEntity, err, t)
}

func TestProfileUpdateDisplayName(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "randomemail@gmail.com", HashedPassword: "A5W4da15S361AD56WD1"},
	)
	tests.EndTestIfError(err, t)

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "RandomUser", TagName: "randomone"})
	tests.EndTestIfError(err, t)

	updatedDisplayName := "newDisplayName"
	err = profileRepo.UpdateDisplayName(id, updatedDisplayName)
	tests.EndTestIfError(err, t)

	retrievedProfile, err := profileRepo.GetByUserID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedDisplayName, retrievedProfile.DisplayName, t)

	updatedTagName := "newTagName"
	err = profileRepo.UpdateTagName(id, updatedTagName)
	tests.EndTestIfError(err, t)

	retrievedProfile, err = profileRepo.GetByUserID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedTagName, retrievedProfile.TagName, t)

	updatedPicturePath := "https://somepage.images.com/5668"
	err = profileRepo.UpdatePicturePath(id, updatedPicturePath)
	tests.EndTestIfError(err, t)

	retrievedProfile, err = profileRepo.GetByUserID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedPicturePath, retrievedProfile.PicturePath, t)

	updatedBackgroundPath := "https://somepage.images.com/87983"
	err = profileRepo.UpdateBackgroundPath(id, updatedBackgroundPath)
	tests.EndTestIfError(err, t)

	retrievedProfile, err = profileRepo.GetByUserID(id)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedBackgroundPath, retrievedProfile.BackgroundPath, t)
}

func TestProfileDelete(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	id, err := userRepo.CreateNew(domain.User{
		Email: "fordeleting@gmail.com", HashedPassword: "A5W4da15S361"},
	)
	tests.EndTestIfError(err, t)

	id, err = profileRepo.CreateNew(domain.Profile{UserID: id, DisplayName: "DyingUser", TagName: "ForDeleting"})
	tests.EndTestIfError(err, t)

	err = profileRepo.Delete(id)
	tests.EndTestIfError(err, t)

	_, err = profileRepo.GetByUserID(id)
	tests.AssertEqu(util.ErrEmptySelection, err, t)
}

func TestProfileAddDeleteFollowers(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	firstID, err := userRepo.CreateNew(domain.User{
		Email: "followed@somemail.com", HashedPassword: "A5W4da15S361"},
	)
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{UserID: firstID, DisplayName: "DyingUser", TagName: "SA51S6D51"})
	tests.EndTestIfError(err, t)

	secondID, err := userRepo.CreateNew(domain.User{
		Email: "follower@somemail.com", HashedPassword: "A5W4da15S361"},
	)
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{UserID: secondID, DisplayName: "DyingUser", TagName: "1A8W5D61AW2D135A"})
	tests.EndTestIfError(err, t)

	err = profileRepo.AddFollow(secondID, firstID)
	tests.EndTestIfError(err, t)

	retrievedFollowed, err := profileRepo.GetByUserID(firstID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(1), retrievedFollowed.Followers, t)

	retrievedFollower, err := profileRepo.GetByUserID(secondID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(1), retrievedFollower.Follows, t)

	profileRepo.DeleteFollow(secondID, firstID)

	retrievedFollowed, err = profileRepo.GetByUserID(firstID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(0), retrievedFollowed.Followers, t)

	retrievedFollower, err = profileRepo.GetByUserID(secondID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(0), retrievedFollower.Follows, t)
}

func TestProfileGetFollowingsAndFollows(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())

	firstID, err := userRepo.CreateNew(domain.User{
		Email: "followed@followed.com", HashedPassword: "A5W4da15S361"},
	)
	tests.EndTestIfError(err, t)

	firstTagName := "Followed"
	_, err = profileRepo.CreateNew(domain.Profile{UserID: firstID, DisplayName: "DyingUser", TagName: firstTagName})
	tests.EndTestIfError(err, t)

	secondID, err := userRepo.CreateNew(domain.User{
		Email: "follower@follower.com", HashedPassword: "A5W4da15S361"},
	)
	tests.EndTestIfError(err, t)

	secondTagName := "Follower"
	_, err = profileRepo.CreateNew(domain.Profile{UserID: secondID, DisplayName: "DyingUser", TagName: secondTagName})
	tests.EndTestIfError(err, t)

	err = profileRepo.AddFollow(secondID, firstID)
	tests.EndTestIfError(err, t)

	retrievedPosts, err := profileRepo.GetFollowersByID(firstID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(secondID, retrievedPosts[0].UserID, t)

	retrievedPosts, err = profileRepo.GetFollowersByTagName(firstTagName)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(secondID, retrievedPosts[0].UserID, t)

	retrievedPosts, err = profileRepo.GetFollowsByID(secondID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(firstID, retrievedPosts[0].UserID, t)

	retrievedPosts, err = profileRepo.GetFollowsByTagName(secondTagName)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(firstID, retrievedPosts[0].UserID, t)
}
