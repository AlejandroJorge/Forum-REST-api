package repository

import (
	"database/sql"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/mattn/go-sqlite3"
)

type sqliteCommentRepository struct {
	db *sql.DB
}

// Can return ErrNoMatchingDependency
func (repo sqliteCommentRepository) AddLike(userId uint, commentId uint) error {
	db := repo.db

	query := `
	INSERT INTO Comment_Likings(Liker_ID, Comment_ID)
	VALUES (?,?)
	`
	_, err := db.Exec(query, userId, commentId)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			logging.LogRepositoryError(ErrNoMatchingDependency)
			return ErrNoMatchingDependency
		}
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	return nil
}

// Returns the id of the created comment, can return ErrNoMatchingDependency
func (repo sqliteCommentRepository) Create(postID, userID uint, content string) (uint, error) {
	db := repo.db

	query := `
	INSERT INTO Comment(Post_ID, User_ID, Content)
	VALUES (?,?,?)
	`
	res, err := db.Exec(query, postID, userID, content)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			logging.LogRepositoryError(ErrNoMatchingDependency)
			return 0, ErrNoMatchingDependency
		}
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return 0, ErrUnknown
	}

	newId, err := res.LastInsertId()
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return 0, ErrUnknown
	}

	return uint(newId), nil
}

// Can return ErrNoRowsAffected
func (repo sqliteCommentRepository) Delete(id uint) error {
	db := repo.db

	query := `
	DELETE FROM Comment
	WHERE Comment_ID = ?
	`
	res, err := db.Exec(query, id)
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	amountAffected, err := res.RowsAffected()
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	if amountAffected == 0 {
		logging.LogRepositoryError(ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	return nil
}

// Can return ErrNoRowsAffected
func (repo sqliteCommentRepository) DeleteLike(userId uint, commentId uint) error {
	db := repo.db

	query := `
	DELETE FROM Comment_Likings
	WHERE Liker_ID = ? AND Comment_ID = ?
	`
	res, err := db.Exec(query, userId, commentId)
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	amountAffected, err := res.RowsAffected()
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	if amountAffected == 0 {
		logging.LogRepositoryError(ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	return nil
}

// Returns a valid comment and can return ErrEmptySelection
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
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.Comment{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.Comment{}, ErrUnknown
	}

	return comment, nil
}

// Returns an slice of valid comments, can return ErrEmptySelection
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var comment domain.Comment
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return comments, nil
}

// Returns an slice of valid comments, can return ErrEmptySelection
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var comment domain.Comment
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Likes)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return comments, nil
}

// Can return ErrNoRowsAffected
func (repo sqliteCommentRepository) UpdateContent(id uint, newContent string) error {
	db := repo.db

	query := `
	UPDATE Comment
	SET Content = ?
	WHERE Comment_ID = ?
	`
	res, err := db.Exec(query, newContent, id)
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	amountAffected, err := res.RowsAffected()
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	if amountAffected == 0 {
		logging.LogRepositoryError(ErrNoRowsAffected)
		return ErrNoRowsAffected
	}

	return nil
}

func NewSQLiteCommentRepository(db *sql.DB) domain.CommentRepository {
	return sqliteCommentRepository{db: db}
}
