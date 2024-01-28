package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/mattn/go-sqlite3"
)

type sqlitePostRepository struct {
	db *sql.DB
}

func (repo sqlitePostRepository) AddLike(userId uint, postId uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO Post_Likings(Liker_ID, Post_ID)
	VALUES (?,?)
	`
	_, err = tx.Exec(query, userId, postId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo sqlitePostRepository) DeleteLike(userId uint, postId uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	DELETE FROM Post_Likings
	WHERE Liker_ID = ? AND Post_ID = ?
	`
	_, err = tx.Exec(query, userId, postId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo sqlitePostRepository) CreateNew(post domain.Post) (uint, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
  INSERT INTO Post(Title, Description, Content, Creation_Date, Owner_ID)
  VALUES (?,?,?,?,?)
  `
	res, err := tx.Exec(query, post.Title, post.Description, post.Content, time.Now().Unix(), post.OwnerID)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingProfile
		}
	}
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint(newId), nil
}

func (repo sqlitePostRepository) Delete(id uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	DELETE FROM Post
	WHERE Post_ID = ?
	`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo sqlitePostRepository) GetByID(id uint) (domain.Post, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return domain.Post{}, err
	}

	var post domain.Post
	var creationDate int64
	query := `
	SELECT p.Post_ID, p.Owner_ID, p.Title, p.Description, p.Content, p.Creation_Date, COUNT(l.Liker_ID)
	FROM Post p
	LEFT JOIN Post_Likings l ON p.Post_ID = l.Post_ID
	WHERE p.Post_ID = ?
	GROUP BY p.Post_ID
	`
	row := tx.QueryRow(query, id)
	err = row.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		tx.Rollback()
		return domain.Post{}, err
	}
	post.CreationDate = time.Unix(creationDate, 0)

	err = tx.Commit()
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (repo sqlitePostRepository) GetByUser(userId uint) ([]domain.Post, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}

	var posts []domain.Post
	query := `
	SELECT p.Post_ID, p.Owner_ID, p.Title, p.Description, p.Content, p.Creation_Date, COUNT(l.Liker_ID)
	FROM Post p
	LEFT JOIN Post_Likings l ON p.Post_ID = l.Post_ID
	WHERE p.Owner_ID = ?
	GROUP BY p.Post_ID
	`
	rows, err := tx.Query(query, userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		var post domain.Post
		var creationDate int64

		err = rows.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
		if err != nil {
			if err == sql.ErrNoRows {
				err = util.ErrEmptySelection
			}
			tx.Rollback()
			return nil, err
		}

		post.CreationDate = time.Unix(creationDate, 0)
		posts = append(posts, post)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo sqlitePostRepository) GetPopularAfter(moment time.Time, amount uint) ([]domain.Post, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}

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
	rows, err := tx.Query(query, momentInteger, amount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		var post domain.Post
		var creationDate int64
		err = rows.Scan(&post.PostID, &post.OwnerID, &post.Title, &post.Description, &post.Content, &creationDate, &post.Likes)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		post.CreationDate = time.Unix(creationDate, 0)
		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, util.ErrEmptySelection
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repo sqlitePostRepository) UpdateContent(id uint, newContent string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	UPDATE Post
	SET	Content = ?
	WHERE Post_ID = ?
	`
	_, err = tx.Exec(query, newContent, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo sqlitePostRepository) UpdateDescription(id uint, newDescription string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	UPDATE Post
	SET	Description = ?
	WHERE Post_ID = ?
	`
	_, err = tx.Exec(query, newDescription, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo sqlitePostRepository) UpdateTitle(id uint, newTitle string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
	UPDATE Post
	SET	Title = ?
	WHERE Post_ID = ?
	`
	_, err = tx.Exec(query, newTitle, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewSQLitePostRepository(db *sql.DB) domain.PostRepository {
	return sqlitePostRepository{db: db}
}
