package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/mattn/go-sqlite3"
)

type sqliteProfileRepository struct {
	db *sql.DB
}

func (repo sqliteProfileRepository) AddFollow(followerId uint, followedId uint) error {
	db := repo.db

	query := `
	INSERT INTO Following(Follower_ID,Followed_ID,Following_Date)
	VALUES (?,?,?)
	`
	_, err := db.Exec(query, followerId, followedId, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) DeleteFollow(followerId uint, followedId uint) error {
	db := repo.db

	query := `
	DELETE FROM Following
	WHERE Follower_ID = ? AND Followed_ID = ?
	`
	_, err := db.Exec(query, followerId, followedId)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) GetFollowersByID(userId uint) ([]domain.Profile, error) {
	db := repo.db

	var posts []domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.User_ID IN (
		SELECT Follower_ID FROM Following WHERE Followed_ID = ?
	)
	GROUP BY p.User_ID
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post domain.Profile
		err = rows.Scan(&post.UserID, &post.DisplayName, &post.TagName, &post.PicturePath, &post.BackgroundPath, &post.Follows, &post.Followers)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, util.ErrEmptySelection
	}

	return posts, nil
}

func (repo sqliteProfileRepository) GetFollowersByTagName(tagName string) ([]domain.Profile, error) {
	db := repo.db

	var posts []domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.User_ID IN (
		SELECT f.Follower_ID FROM Following f, Profile p 
		WHERE f.Followed_ID = p.User_ID AND p.Tag_Name = ?
	)
	GROUP BY p.User_ID
	`
	rows, err := db.Query(query, tagName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post domain.Profile
		err = rows.Scan(&post.UserID, &post.DisplayName, &post.TagName, &post.PicturePath, &post.BackgroundPath, &post.Follows, &post.Followers)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, util.ErrEmptySelection
	}

	return posts, nil
}

func (repo sqliteProfileRepository) GetFollowsByID(userId uint) ([]domain.Profile, error) {
	db := repo.db

	var posts []domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.User_ID IN (
		SELECT Followed_ID FROM Following WHERE Follower_ID = ?
	)
	GROUP BY p.User_ID
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post domain.Profile
		err = rows.Scan(&post.UserID, &post.DisplayName, &post.TagName, &post.PicturePath, &post.BackgroundPath, &post.Follows, &post.Followers)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, util.ErrEmptySelection
	}

	return posts, nil
}

func (repo sqliteProfileRepository) GetFollowsByTagName(tagName string) ([]domain.Profile, error) {
	db := repo.db

	var posts []domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.User_ID IN (
		SELECT f.Followed_ID FROM Following f, Profile p 
		WHERE f.Follower_ID = p.User_ID AND p.Tag_Name = ?
	)
	GROUP BY p.User_ID
	`
	rows, err := db.Query(query, tagName)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post domain.Profile
		err = rows.Scan(&post.UserID, &post.DisplayName, &post.TagName, &post.PicturePath, &post.BackgroundPath, &post.Follows, &post.Followers)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, util.ErrEmptySelection
	}

	return posts, nil
}

func (repo sqliteProfileRepository) CreateNew(profile domain.Profile) (uint, error) {
	db := repo.db

	query := `
  INSERT INTO Profile(User_ID, Display_Name, Tag_Name, Picture_Path, Background_Path)
  VALUES (?,?,?,?,?)
  `
	res, err := db.Exec(query, profile.UserID, profile.DisplayName, profile.TagName, profile.PicturePath, profile.BackgroundPath)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			err = util.ErrRepeatedEntity
		}
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingUser
		}
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			err = util.ErrRepeatedEntity
		}
	}
	if err != nil {
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(newId), nil
}

func (repo sqliteProfileRepository) Delete(id uint) error {
	db := repo.db

	query := `
  DELETE FROM Profile 
  WHERE User_ID = ?
  `
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) GetByTagName(tagName string) (domain.Profile, error) {
	db := repo.db

	var profile domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.Tag_Name = ?
	GROUP BY p.User_ID
  `
	row := db.QueryRow(query, tagName)
	err := row.Scan(&profile.UserID, &profile.DisplayName, &profile.TagName, &profile.PicturePath, &profile.BackgroundPath, &profile.Followers, &profile.Follows)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.Profile{}, err
	}

	return profile, nil
}

func (repo sqliteProfileRepository) GetByUserID(userId uint) (domain.Profile, error) {
	db := repo.db

	var profile domain.Profile
	query := `
  SELECT p.User_ID, p.Display_Name, p.Tag_Name, p.Picture_Path, p.Background_Path, COUNT(f1.Follower_ID), COUNT(f2.Followed_ID)
  FROM Profile p
	LEFT JOIN Following f1 ON p.User_ID = f1.Followed_ID
	LEFT JOIN Following f2 ON p.User_ID = f2.Follower_ID
  WHERE p.User_ID = ?
	GROUP BY p.User_ID
  `
	row := db.QueryRow(query, userId)
	err := row.Scan(&profile.UserID, &profile.DisplayName, &profile.TagName, &profile.PicturePath, &profile.BackgroundPath, &profile.Followers, &profile.Follows)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.Profile{}, err
	}

	return profile, nil
}

func (repo sqliteProfileRepository) UpdateBackgroundPath(id uint, newBackgroundPath string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Background_Path = ?
	WHERE User_ID = ?
	`
	_, err := db.Exec(query, newBackgroundPath, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) UpdateDisplayName(id uint, newDisplayName string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Display_Name = ?
	WHERE User_ID = ?
	`
	_, err := db.Exec(query, newDisplayName, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) UpdatePicturePath(id uint, newPicturePath string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Picture_Path = ?
	WHERE User_ID = ?
	`
	_, err := db.Exec(query, newPicturePath, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteProfileRepository) UpdateTagName(id uint, newTagName string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Tag_Name = ?
	WHERE User_ID = ?
	`
	_, err := db.Exec(query, newTagName, id)
	if err != nil {
		return err
	}

	return nil
}

func NewSQLiteProfileRepository(db *sql.DB) domain.ProfileRepository {
	return sqliteProfileRepository{db: db}
}
