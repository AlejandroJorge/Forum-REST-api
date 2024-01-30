package repository

import (
	"database/sql"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/mattn/go-sqlite3"
)

type sqliteCommentRepository struct {
	db *sql.DB
}

func (repo sqliteCommentRepository) AddLike(userId uint, commentId uint) error {
	db := repo.db

	query := `
	INSERT INTO Comment_Likings(Liker_ID, Comment_ID)
	VALUES (?,?)
	`
	_, err := db.Exec(query, userId, commentId)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingProfileOrComment
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteCommentRepository) CreateNew(comment domain.Comment) (uint, error) {
	db := repo.db

	query := `
	INSERT INTO Comment(Post_ID, User_ID, Content)
	VALUES (?,?,?)
	`
	res, err := db.Exec(query, comment.PostID, comment.UserID, comment.Content)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingProfileOrPost
		}
	}
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (repo sqliteCommentRepository) Delete(id uint) error {
	db := repo.db

	query := `
	DELETE FROM Comment
	WHERE Comment_ID = ?
	`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteCommentRepository) DeleteLike(userId uint, commentId uint) error {
	db := repo.db

	query := `
	DELETE FROM Comment_Likings
	WHERE Liker_ID = ? AND Comment_ID = ?
	`
	_, err := db.Exec(query, userId, commentId)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingProfileOrComment
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteCommentRepository) GetByID(id uint) (domain.Comment, error) {
	db := repo.db

	var comment domain.Comment
	query := `
	SELECT c.Comment_ID, c.Post_ID, c.User_ID, c.Content, COUNT(l.Liker_ID) AS Like_Count
	FROM Comment c
	LEFT JOIN Comment_Likings l ON c.Comment_ID = l.Comment_ID
	WHERE c.Comment_ID = ?
	GROUP BY c.Comment_ID
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.Comment{}, err
	}

	return comment, nil
}

func (repo sqliteCommentRepository) GetByPost(postID uint) ([]domain.Comment, error) {
	db := repo.db

	var comments []domain.Comment
	query := `
	SELECT c.Comment_ID, c.Post_ID, c.User_ID, c.Content, COUNT(l.Liker_ID) AS Like_Count
	FROM Comment c
	LEFT JOIN Comment_Likings l ON c.Comment_ID = l.Comment_ID
	WHERE c.Comment_ID IN(
		SELECT Comment_ID FROM Comment WHERE Post_ID = ?
	)
	GROUP BY c.Comment_ID
	`
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment domain.Comment
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		return nil, util.ErrEmptySelection
	}

	return comments, nil
}

func (repo sqliteCommentRepository) GetByUser(userID uint) ([]domain.Comment, error) {
	db := repo.db

	var comments []domain.Comment
	query := `
	SELECT c.Comment_ID, c.Post_ID, c.User_ID, c.Content, COUNT(l.Liker_ID) AS Like_Count
	FROM Comment c
	LEFT JOIN Comment_Likings l ON c.Comment_ID = l.Comment_ID
	WHERE c.Comment_ID IN(
		SELECT Comment_ID FROM Comment WHERE User_ID = ?
	)
	GROUP BY c.Comment_ID
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment domain.Comment
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		return nil, util.ErrEmptySelection
	}

	return comments, nil
}

func (repo sqliteCommentRepository) UpdateContent(id uint, newContent string) error {
	db := repo.db

	query := `
	UPDATE Comment
	SET Content = ?
	WHERE Comment_ID = ?
	`
	_, err := db.Exec(query, newContent, id)
	if err != nil {
		return err
	}

	return nil
}

func NewSQLiteCommentRepository(db *sql.DB) domain.CommentRepository {
	return sqliteCommentRepository{db: db}
}
