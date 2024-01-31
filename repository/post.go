package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/mattn/go-sqlite3"
)

type sqlitePostRepository struct {
	db *sql.DB
}

// Can return ErrRepeatedEntity, ErrNoMatchingDependency
func (repo sqlitePostRepository) AddLike(userId uint, postId uint) error {
	db := repo.db

	query := `
	INSERT INTO Post_Likings(Liker_ID, Post_ID)
	VALUES (?,?)
	`
	_, err := db.Exec(query, userId, postId)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			logging.LogRepositoryError(ErrNoMatchingDependency)
			return ErrNoMatchingDependency
		}
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			logging.LogRepositoryError(ErrRepeatedEntity)
			return ErrRepeatedEntity
		}
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return ErrUnknown
	}

	return nil
}

// Can return ErrNoRowsAffected
func (repo sqlitePostRepository) DeleteLike(userId uint, postId uint) error {
	db := repo.db

	query := `
	DELETE FROM Post_Likings
	WHERE Liker_ID = ? AND Post_ID = ?
	`
	res, err := db.Exec(query, userId, postId)
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

// Returns the id of the created post and can return ErrNoMatchingDependency, ErrRepeatedEntity
func (repo sqlitePostRepository) Create(ownerID uint, title, description, content string) (uint, error) {
	db := repo.db

	query := `
  INSERT INTO Post(Title, Description, Content, Creation_Date, Owner_ID)
  VALUES (?,?,?,?,?)
  `
	res, err := db.Exec(query, title, description, content, time.Now().Unix(), ownerID)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			logging.LogRepositoryError(ErrNoMatchingDependency)
			return 0, ErrNoMatchingDependency
		}
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			logging.LogRepositoryError(ErrRepeatedEntity)
			return 0, ErrRepeatedEntity
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
func (repo sqlitePostRepository) Delete(id uint) error {
	db := repo.db

	query := `
	DELETE FROM Post
	WHERE Post_ID = ?
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

// Returns a valid profile and can return ErrEmptySelection
func (repo sqlitePostRepository) GetByID(id uint) (domain.Post, error) {
	db := repo.db

	var post domain.Post
	var creationDate int64
	query := `
	SELECT p.Post_ID, p.Owner_ID, p.Title, p.Description, p.Content, p.Creation_Date, COUNT(l.Liker_ID)
	FROM Post p
	LEFT JOIN Post_Likings l ON p.Post_ID = l.Post_ID
	WHERE p.Post_ID = ?
	GROUP BY p.Post_ID
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.Post{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.Post{}, ErrUnknown
	}

	return post, nil
}

// Can return ErrEmptySelection
func (repo sqlitePostRepository) GetByUser(userId uint) ([]domain.Post, error) {
	db := repo.db

	var posts []domain.Post
	query := `
	SELECT p.Post_ID, p.Owner_ID, p.Title, p.Description, p.Content, p.Creation_Date, COUNT(l.Liker_ID)
	FROM Post p
	LEFT JOIN Post_Likings l ON p.Post_ID = l.Post_ID
	WHERE p.Owner_ID = ?
	GROUP BY p.Post_ID
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var post domain.Post
		var creationDate int64

		err = rows.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		post.CreationDate = time.Unix(creationDate, 0)
		posts = append(posts, post)
	}

	if len(posts) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return posts, nil
}

// Can return ErrEmptySelection
func (repo sqlitePostRepository) GetPopularAfter(moment time.Time, amount uint) ([]domain.Post, error) {
	db := repo.db

	var posts []domain.Post
	momentInteger := moment.Unix()
	query := `
	SELECT p.Post_ID, p.Owner_ID, p.Title, p.Description, p.Content, p.Creation_Date, COUNT(l.Liker_ID) AS Like_Count
	FROM Post p
	LEFT JOIN Post_Likings l ON p.Post_ID = l.Post_ID
	WHERE p.Creation_Date >= ?
	GROUP BY p.Post_ID
	ORDER BY Like_Count DESC
	LIMIT ?
	`
	rows, err := db.Query(query, momentInteger, amount)
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var post domain.Post
		var creationDate int64
		err = rows.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		post.CreationDate = time.Unix(creationDate, 0)
		posts = append(posts, post)
	}

	if len(posts) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return posts, nil
}

// Can return ErrNoRowsAffected
func (repo sqlitePostRepository) UpdateContent(id uint, newContent string) error {
	db := repo.db

	query := `
	UPDATE Post
	SET	Content = ?
	WHERE Post_ID = ?
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

// Can return ErrNoRowsAffected
func (repo sqlitePostRepository) UpdateDescription(id uint, newDescription string) error {
	db := repo.db

	query := `
	UPDATE Post
	SET	Description = ?
	WHERE Post_ID = ?
	`
	res, err := db.Exec(query, newDescription, id)
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
func (repo sqlitePostRepository) UpdateTitle(id uint, newTitle string) error {
	db := repo.db

	query := `
	UPDATE Post
	SET	Title = ?
	WHERE Post_ID = ?
	`
	res, err := db.Exec(query, newTitle, id)
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

func NewSQLitePostRepository(db *sql.DB) domain.PostRepository {
	return sqlitePostRepository{db: db}
}
