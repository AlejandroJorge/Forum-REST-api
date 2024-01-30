package repository

import (
	"testing"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/tests"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestPostCreateAndRead(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "adsasas@asdasd.com", HashedPassword: "1dw8a15s",
	})
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "A65S1D6AS1D",
	})
	tests.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Some title", Description: "Stuff", Content: "516ASD165ASD",
	}
	postID, err := postRepo.CreateNew(newPost)
	tests.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(newPost.Title, retrievedPost.Title, t)
	tests.AssertEqu(newPost.Description, retrievedPost.Description, t)
	tests.AssertEqu(newPost.Content, retrievedPost.Content, t)
	tests.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
}

func TestPostCreateForNoProfile(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "1a56sd1a6sd1@unexistent.com", HashedPassword: "1dw8a15s",
	})
	tests.EndTestIfError(err, t)

	_, err = postRepo.CreateNew(domain.Post{
		Title: "Some random title", Content: "Something", Description: "asdasd", OwnerID: userID,
	})
	tests.AssertEqu(util.ErrNoCorrespondingProfile, err, t)
}

func TestPostGetMultiplePosts(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "9aw8d45as41d98@a.com", HashedPassword: "1dw8a15s",
	})
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "w8a1d96a531",
	})
	tests.EndTestIfError(err, t)

	newPosts := []domain.Post{
		{OwnerID: userID, Title: "A title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Another one", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Some title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Just title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
	}

	for _, newPost := range newPosts {
		newPostID, err := postRepo.CreateNew(newPost)
		tests.EndTestIfError(err, t)

		retrievedPost, err := postRepo.GetByID(newPostID)
		tests.EndTestIfError(err, t)

		tests.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
		tests.AssertEqu(newPost.Title, retrievedPost.Title, t)
		tests.AssertEqu(newPost.Description, retrievedPost.Description, t)
		tests.AssertEqu(newPost.Content, retrievedPost.Content, t)
	}

	retrievedPosts, err := postRepo.GetByUser(userID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(len(newPosts), len(retrievedPosts), t)

	for i := range retrievedPosts {
		newPost := newPosts[i]
		retrievedPost := retrievedPosts[i]

		tests.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
		tests.AssertEqu(newPost.Title, retrievedPost.Title, t)
		tests.AssertEqu(newPost.Description, retrievedPost.Description, t)
		tests.AssertEqu(newPost.Content, retrievedPost.Content, t)
	}
}

func TestPostGetPopularAfterFuture(t *testing.T) {
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	retrievedPosts, err := postRepo.GetPopularAfter(time.Now().Add(time.Minute), 10)
	tests.AssertEqu(util.ErrEmptySelection, err, t)
	tests.AssertEqu(0, len(retrievedPosts), t)
}

func TestPostUpdate(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "A84S8D6416@asdasd.com", HashedPassword: "8796A5S41D536A4D",
	})
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "as1486asd648as5dsa51",
	})
	tests.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Some title", Description: "Stuff", Content: "516ASD165ASD",
	}
	postID, err := postRepo.CreateNew(newPost)
	tests.EndTestIfError(err, t)

	updatedTitle := "A48S6D54ASD"
	err = postRepo.UpdateTitle(postID, updatedTitle)
	tests.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedTitle, retrievedPost.Title, t)

	updatedDescription := "Some description #a5sd"
	err = postRepo.UpdateDescription(postID, updatedDescription)
	tests.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedDescription, retrievedPost.Description, t)

	updatedContent := "Content a6s54da6s5d asda Some more content"
	err = postRepo.UpdateContent(postID, updatedContent)
	tests.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(updatedContent, retrievedPost.Content, t)
}

func TestPostDelete(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "random@email4654.com", HashedPassword: "8796A5S41D536A4D",
	})
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "uniqueN4me",
	})
	tests.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Another title", Description: "d3scription", Content: "6845DAWS4D6A",
	}
	postID, err := postRepo.CreateNew(newPost)
	tests.EndTestIfError(err, t)

	err = postRepo.Delete(postID)
	tests.EndTestIfError(err, t)

	_, err = postRepo.GetByID(postID)
	tests.AssertEqu(util.ErrEmptySelection, err, t)
}

func TestPostLikes(t *testing.T) {
	userRepo := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := repository.NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "raskld@54a.com", HashedPassword: "8796A5S41D",
	})
	tests.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "15asd1ADCB",
	})
	tests.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Another title", Description: "d3scription", Content: "6845DAWS4D6A",
	})
	tests.EndTestIfError(err, t)

	err = postRepo.AddLike(userID, postID)
	tests.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(1), retrievedPost.Likes, t)

	err = postRepo.DeleteLike(userID, postID)
	tests.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	tests.EndTestIfError(err, t)

	tests.AssertEqu(uint(0), retrievedPost.Likes, t)
}
