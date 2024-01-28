package repository

import (
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestCommentCreateWithNoPostNoUser(t *testing.T) {
	commentRepo := NewSQLiteCommentRepository(config.SQLiteDatabase())
	_, err := commentRepo.CreateNew(domain.Comment{UserID: 0, PostID: 0, Content: "Something"})
	util.AssertEqu(util.ErrNoCorrespondingProfileOrPost, err, t)
}

func TestCommentGetMultiple(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())
	commentRepo := NewSQLiteCommentRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "A15SD61@commentrepo.com", HashedPassword: "5A6SD1",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, TagName: "Some name for comment repo testing",
	})
	util.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Smth", Description: "asd", Content: "5a61sd",
	})
	util.EndTestIfError(err, t)

	newComments := []domain.Comment{
		{PostID: postID, UserID: userID, Content: "Some comment"},
		{PostID: postID, UserID: userID, Content: "Another comment"},
		{PostID: postID, UserID: userID, Content: "Some different comment"},
	}
	for _, newComment := range newComments {
		commentID, err := commentRepo.CreateNew(newComment)
		util.EndTestIfError(err, t)

		retrievedComment, err := commentRepo.GetByID(commentID)
		util.EndTestIfError(err, t)

		util.AssertEqu(newComment.PostID, retrievedComment.PostID, t)
		util.AssertEqu(newComment.UserID, retrievedComment.UserID, t)
		util.AssertEqu(newComment.Content, retrievedComment.Content, t)
	}

	retrievedComments, err := commentRepo.GetByPost(postID)
	util.EndTestIfError(err, t)

	for i := range newComments {
		util.AssertEqu(newComments[i].PostID, retrievedComments[i].PostID, t)
		util.AssertEqu(newComments[i].UserID, retrievedComments[i].UserID, t)
		util.AssertEqu(newComments[i].Content, retrievedComments[i].Content, t)
	}

	retrievedComments, err = commentRepo.GetByUser(userID)
	util.EndTestIfError(err, t)

	for i := range newComments {
		util.AssertEqu(newComments[i].PostID, retrievedComments[i].PostID, t)
		util.AssertEqu(newComments[i].UserID, retrievedComments[i].UserID, t)
		util.AssertEqu(newComments[i].Content, retrievedComments[i].Content, t)
	}
}

func TestCommentLikes(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())
	commentRepo := NewSQLiteCommentRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "5A1S561@commentrepo.com", HashedPassword: "5A6SD1",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, TagName: "Some other name for comment repo testing",
	})
	util.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Smth", Description: "asd", Content: "5a61sd",
	})
	util.EndTestIfError(err, t)

	commentID, err := commentRepo.CreateNew(domain.Comment{
		PostID: postID, UserID: userID, Content: "Some content",
	})

	err = commentRepo.AddLike(userID, commentID)
	util.EndTestIfError(err, t)

	retrievedComment, err := commentRepo.GetByID(commentID)
	util.EndTestIfError(err, t)

	util.AssertEqu(uint(1), retrievedComment.Likes, t)

	err = commentRepo.DeleteLike(userID, commentID)
	util.EndTestIfError(err, t)

	retrievedComment, err = commentRepo.GetByID(commentID)
	util.EndTestIfError(err, t)

	util.AssertEqu(uint(0), retrievedComment.Likes, t)
}

func TestCommentUpdate(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())
	commentRepo := NewSQLiteCommentRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "84s45@commentrepo.com", HashedPassword: "5A6SD1",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, TagName: "Made for updating a comment",
	})
	util.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Smth", Description: "asd", Content: "5a61sd",
	})
	util.EndTestIfError(err, t)

	commentID, err := commentRepo.CreateNew(domain.Comment{
		PostID: postID, UserID: userID, Content: "AAAAAAA",
	})

	updatedContent := "BBBBBBBBBB"
	err = commentRepo.UpdateContent(commentID, updatedContent)
	util.EndTestIfError(err, t)

	retrievedComment, err := commentRepo.GetByID(commentID)
	util.EndTestIfError(err, t)

	util.AssertEqu(updatedContent, retrievedComment.Content, t)
}

func TestCommentDelete(t *testing.T) {
	userRepo := NewSQLiteUserRepository(config.SQLiteDatabase())
	profileRepo := NewSQLiteProfileRepository(config.SQLiteDatabase())
	postRepo := NewSQLitePostRepository(config.SQLiteDatabase())
	commentRepo := NewSQLiteCommentRepository(config.SQLiteDatabase())

	userID, err := userRepo.CreateNew(domain.User{
		Email: "a68s4d86as4d@commentrepo.com", HashedPassword: "5A6SD1",
	})
	util.EndTestIfError(err, t)

	_, err = profileRepo.CreateNew(domain.Profile{
		UserID: userID, TagName: "Made for deleting a comment",
	})
	util.EndTestIfError(err, t)

	postID, err := postRepo.CreateNew(domain.Post{
		OwnerID: userID, Title: "Smth", Description: "asd", Content: "5a61sd",
	})
	util.EndTestIfError(err, t)

	commentID, err := commentRepo.CreateNew(domain.Comment{
		PostID: postID, UserID: userID, Content: "Some content",
	})

	err = commentRepo.Delete(commentID)
	util.EndTestIfError(err, t)

	_, err = commentRepo.GetByID(commentID)
	util.AssertEqu(util.ErrEmptySelection, err, t)
}
