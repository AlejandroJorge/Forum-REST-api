package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/mattn/go-sqlite3"
)

type sqliteProfileRepository struct {
	db *sql.DB
}

// Can return ErrRepeatedEntity, ErrNoMatchingDependency
func (repo sqliteProfileRepository) AddFollow(followerId uint, followedId uint) error {
	db := repo.db

	query := `
	INSERT INTO Following(Follower_ID,Followed_ID,Following_Date)
	VALUES (?,?,?)
	`
	_, err := db.Exec(query, followerId, followedId, time.Now().Unix())
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
func (repo sqliteProfileRepository) DeleteFollow(followerId uint, followedId uint) error {
	db := repo.db

	query := `
	DELETE FROM Following
	WHERE Follower_ID = ? AND Followed_ID = ?
	`
	res, err := db.Exec(query, followerId, followedId)
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

// Returns an slice of valid profiles, can return ErrEmptySelection
func (repo sqliteProfileRepository) GetFollowersByID(userId uint) ([]domain.Profile, error) {
	db := repo.db

	var profiles []domain.Profile
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var p domain.Profile
		err = rows.Scan(&p.UserID, &p.DisplayName, &p.TagName, &p.PicturePath, &p.BackgroundPath, &p.Follows, &p.Followers)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		profiles = append(profiles, p)
	}

	if len(profiles) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return profiles, nil
}

// Returns an slice of valid profiles, can return ErrEmptySelection
func (repo sqliteProfileRepository) GetFollowersByTagName(tagName string) ([]domain.Profile, error) {
	db := repo.db

	var profiles []domain.Profile
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var p domain.Profile
		err = rows.Scan(&p.UserID, &p.DisplayName, &p.TagName, &p.PicturePath, &p.BackgroundPath, &p.Follows, &p.Followers)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		profiles = append(profiles, p)
	}

	if len(profiles) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return profiles, nil
}

// Returns an slice of valid profiles, can return ErrEmptySelection
func (repo sqliteProfileRepository) GetFollowsByID(userId uint) ([]domain.Profile, error) {
	db := repo.db

	var profiles []domain.Profile
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var p domain.Profile
		err = rows.Scan(&p.UserID, &p.DisplayName, &p.TagName, &p.PicturePath, &p.BackgroundPath, &p.Follows, &p.Followers)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		profiles = append(profiles, p)
	}

	if len(profiles) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return profiles, nil
}

// Returns an slice of valid profiles, can return ErrEmptySelection
func (repo sqliteProfileRepository) GetFollowsByTagName(tagName string) ([]domain.Profile, error) {
	db := repo.db

	var profiles []domain.Profile
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
		logging.LogUnexpectedRepositoryError(err)
		return nil, ErrUnknown
	}

	for rows.Next() {
		var p domain.Profile
		err = rows.Scan(&p.UserID, &p.DisplayName, &p.TagName, &p.PicturePath, &p.BackgroundPath, &p.Follows, &p.Followers)
		if err != nil {
			logging.LogUnexpectedRepositoryError(err)
			return nil, ErrUnknown
		}

		profiles = append(profiles, p)
	}

	if len(profiles) == 0 {
		logging.LogRepositoryError(ErrEmptySelection)
		return nil, ErrEmptySelection
	}

	return profiles, nil
}

// Returns the id of the created profile, can return ErrNoMatchingDependency, ErrRepeatedEntity
func (repo sqliteProfileRepository) Create(userID uint, tagName, displayName string) (uint, error) {
	db := repo.db

	query := `
  INSERT INTO Profile(User_ID, Display_Name, Tag_Name)
  VALUES (?,?,?,?,?)
  `
	res, err := db.Exec(query, userID, displayName, tagName)
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
func (repo sqliteProfileRepository) Delete(id uint) error {
	db := repo.db

	query := `
  DELETE FROM Profile 
  WHERE User_ID = ?
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
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.Profile{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.Profile{}, ErrUnknown
	}

	return profile, nil
}

// Returns a valid profile and can return ErrEmptySelection
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
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.Profile{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.Profile{}, ErrUnknown
	}

	return profile, nil
}

// Can return ErrNoRowsAffected
func (repo sqliteProfileRepository) UpdateBackgroundPath(id uint, newBackgroundPath string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Background_Path = ?
	WHERE User_ID = ?
	`
	res, err := db.Exec(query, newBackgroundPath, id)
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
func (repo sqliteProfileRepository) UpdateDisplayName(id uint, newDisplayName string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Display_Name = ?
	WHERE User_ID = ?
	`
	res, err := db.Exec(query, newDisplayName, id)
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
func (repo sqliteProfileRepository) UpdatePicturePath(id uint, newPicturePath string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Picture_Path = ?
	WHERE User_ID = ?
	`
	res, err := db.Exec(query, newPicturePath, id)
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
func (repo sqliteProfileRepository) UpdateTagName(id uint, newTagName string) error {
	db := repo.db

	query := `
  UPDATE Profile
	SET Tag_Name = ?
	WHERE User_ID = ?
	`
	res, err := db.Exec(query, newTagName, id)
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

func NewSQLiteProfileRepository(db *sql.DB) domain.ProfileRepository {
	return sqliteProfileRepository{db: db}
}
