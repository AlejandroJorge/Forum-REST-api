package repository

import (
	"database/sql"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/mattn/go-sqlite3"
)

type sqliteProfileRepository struct {
	db *sql.DB
}

func (repo sqliteProfileRepository) CreateNew(profile domain.Profile) (uint, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
  INSERT INTO Profile(User_ID, Display_Name, Tag_Name, Picture_Path, Background_Path)
  VALUES (?,?,?,?,?)
  `
	res, err := tx.Exec(query, profile.UserID, profile.DisplayName, profile.TagName, profile.PicturePath, profile.BackgroundPath)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
			err = util.ErrRepeatedEntity
		}
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			err = util.ErrNoCorrespondingUser
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

func (repo sqliteProfileRepository) Delete(id uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  DELETE FROM Profile 
  WHERE User_ID = ?
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

func (repo sqliteProfileRepository) GetByTagName(tagName string) (domain.Profile, error) {
	var profile domain.Profile
	query := `
  SELECT User_ID, Display_Name, Tag_Name, Picture_Path, Background_Path 
  FROM Profile
  WHERE Tag_Name = ?
  `
	row := repo.db.QueryRow(query, tagName)
	err := row.Scan(&profile.UserID, &profile.DisplayName, &profile.TagName, &profile.PicturePath, &profile.BackgroundPath)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.Profile{}, err
	}

	query = `
	SELECT COUNT(*)
	FROM Following
	WHERE Followed_ID = ?
	`
	row = repo.db.QueryRow(query, profile.UserID)
	err = row.Scan(&profile.Followers)
	if err != nil {
		return domain.Profile{}, err
	}

	query = `
	SELECT COUNT(*)
	FROM Following
	WHERE Follower_ID = ?
	`
	row = repo.db.QueryRow(query, profile.UserID)
	err = row.Scan(&profile.Follows)
	if err != nil {
		return domain.Profile{}, err
	}

	return profile, nil
}

func (repo sqliteProfileRepository) GetByUserID(userId uint) (domain.Profile, error) {
	var profile domain.Profile
	query := `
  SELECT User_ID, Display_Name, Tag_Name, Picture_Path, Background_Path 
  FROM Profile
  WHERE User_ID = ?
  `
	row := repo.db.QueryRow(query, userId)
	err := row.Scan(&profile.UserID, &profile.DisplayName, &profile.TagName, &profile.PicturePath, &profile.BackgroundPath)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.Profile{}, err
	}

	query = `
	SELECT COUNT(*)
	FROM Following
	WHERE Followed_ID = ?
	`
	row = repo.db.QueryRow(query, profile.UserID)
	err = row.Scan(&profile.Followers)
	if err != nil {
		return domain.Profile{}, err
	}

	query = `
	SELECT COUNT(*)
	FROM Following
	WHERE Follower_ID = ?
	`
	row = repo.db.QueryRow(query, profile.UserID)
	err = row.Scan(&profile.Follows)
	if err != nil {
		return domain.Profile{}, err
	}

	return profile, nil
}

func (repo sqliteProfileRepository) UpdateBackgroundPath(id uint, newBackgroundPath string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE Profile
	SET Background_Path = ?
	WHERE User_ID = ?
	`
	_, err = tx.Exec(query, newBackgroundPath, id)
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

func (repo sqliteProfileRepository) UpdateDisplayName(id uint, newDisplayName string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE Profile
	SET Display_Name = ?
	WHERE User_ID = ?
	`
	_, err = tx.Exec(query, newDisplayName, id)
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

func (repo sqliteProfileRepository) UpdatePicturePath(id uint, newPicturePath string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE Profile
	SET Picture_Path = ?
	WHERE User_ID = ?
	`
	_, err = tx.Exec(query, newPicturePath, id)
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

func (repo sqliteProfileRepository) UpdateTagName(id uint, newTagName string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE Profile
	SET Tag_Name = ?
	WHERE User_ID = ?
	`
	_, err = tx.Exec(query, newTagName, id)
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

func NewSQLiteProfileRepository(db *sql.DB) domain.ProfileRepository {
	return sqliteProfileRepository{db: db}
}
