package repository

import (
	"testing"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestPostCreateAndRead(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "adsasas@asdasd.com", HashedPassword: "1dw8a15s",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "A65S1D6AS1D",
	})
	util.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Some title", Description: "Stuff", Content: "516ASD165ASD",
	}
	postID, err := postRepo.CreateNew(newPost)
	util.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(newPost.Title, retrievedPost.Title, t)
	util.AssertEqu(newPost.Description, retrievedPost.Description, t)
	util.AssertEqu(newPost.Content, retrievedPost.Content, t)
	util.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
}

func TestPostCreateForNoProfile(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "1a56sd1a6sd1@unexistent.com", HashedPassword: "1dw8a15s",
	})
	util.EndTestIfError(err, t)

	_, err = postRepo.CreateNew(domain.Post{
		Title: "Some random title", Content: "Something", Description: "asdasd", OwnerID: userID,
	})
	util.AssertEqu(util.ErrNoCorrespondingProfile, err, t)
}

func TestPostGetMultiplePosts(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "9aw8d45as41d98@a.com", HashedPassword: "1dw8a15s",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "w8a1d96a531",
	})
	util.EndTestIfError(err, t)

	newPosts := []domain.Post{
		{OwnerID: userID, Title: "A title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Another one", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Some title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
		{OwnerID: userID, Title: "Just title", Description: "5AS16D", Content: "5ADS16D51A3AS1D3A"},
	}

	for _, newPost := range newPosts {
		newPostID, err := postRepo.CreateNew(newPost)
		util.EndTestIfError(err, t)

		retrievedPost, err := postRepo.GetByID(newPostID)
		util.EndTestIfError(err, t)

		util.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
		util.AssertEqu(newPost.Title, retrievedPost.Title, t)
		util.AssertEqu(newPost.Description, retrievedPost.Description, t)
		util.AssertEqu(newPost.Content, retrievedPost.Content, t)
	}

	retrievedPosts, err := postRepo.GetByUser(userID)
	util.EndTestIfError(err, t)

	util.AssertEqu(len(newPosts), len(retrievedPosts), t)

	for i := range retrievedPosts {
		newPost := newPosts[i]
		retrievedPost := retrievedPosts[i]

		util.AssertEqu(newPost.OwnerID, retrievedPost.OwnerID, t)
		util.AssertEqu(newPost.Title, retrievedPost.Title, t)
		util.AssertEqu(newPost.Description, retrievedPost.Description, t)
		util.AssertEqu(newPost.Content, retrievedPost.Content, t)
	}
}

func TestPostGetPopularAfterFuture(t *testing.T) {
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	retrievedPosts, err := postRepo.GetPopularAfter(time.Now().Add(time.Minute), 10)
	util.AssertEqu(util.ErrEmptySelection, err, t)
	util.AssertEqu(0, len(retrievedPosts), t)
}

func TestPostUpdate(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "A84S8D6416@asdasd.com", HashedPassword: "8796A5S41D536A4D",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "as1486asd648as5dsa51",
	})
	util.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Some title", Description: "Stuff", Content: "516ASD165ASD",
	}
	postID, err := postRepo.CreateNew(newPost)
	util.EndTestIfError(err, t)

	updatedTitle := "A48S6D54ASD"
	err = postRepo.UpdateTitle(postID, updatedTitle)
	util.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(updatedTitle, retrievedPost.Title, t)

	updatedDescription := "Some description #a5sd"
	err = postRepo.UpdateDescription(postID, updatedDescription)
	util.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(updatedDescription, retrievedPost.Description, t)

	updatedContent := "Content a6s54da6s5d asda Some more content"
	err = postRepo.UpdateContent(postID, updatedContent)
	util.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(updatedContent, retrievedPost.Content, t)
}

func TestPostDelete(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "random@email4654.com", HashedPassword: "8796A5S41D536A4D",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "uniqueN4me",
	})
	util.EndTestIfError(err, t)

	newPost := domain.Post{
		OwnerID: userID, Title: "Another title", Description: "d3scription", Content: "6845DAWS4D6A",
	}
	postID, err := postRepo.CreateNew(newPost)
	util.EndTestIfError(err, t)

	err = postRepo.Delete(postID)
	util.EndTestIfError(err, t)

	_, err = postRepo.GetByID(postID)
	util.AssertEqu(util.ErrEmptySelection, err, t)
}

func TestPostLikes(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "raskld@54a.com", HashedPassword: "8796A5S41D",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, DisplayName: "Some name", TagName: "15asd1ADCB",
	})
	util.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Another title", Description: "d3scription", Content: "6845DAWS4D6A",
	})
	util.EndTestIfError(err, t)

	err = postRepo.AddLike(userID, postID)
	util.EndTestIfError(err, t)

	retrievedPost, err := postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(1, retrievedPost.Likes, t)

	err = postRepo.DeleteLike(userID, postID)
	util.EndTestIfError(err, t)

	retrievedPost, err = postRepo.GetByID(postID)
	util.EndTestIfError(err, t)

	util.AssertEqu(0, retrievedPost.Likes, t)
}
