package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteUserRepository struct {
	db *sql.DB
}

func (repo sqliteUserRepository) CreateNew(user domain.User) (uint, error) {
	db := repo.db

	query := `
  INSERT INTO User(Email, Hashed_Password, Registration_Date)
  VALUES (?,?,?)
  `
	res, err := db.Exec(query, user.Email, user.HashedPassword, time.Now().Unix())
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return 0, util.ErrRepeatedEntity
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

func (repo sqliteUserRepository) Delete(id uint) error {
	db := repo.db

	query := `
  DELETE FROM User
  WHERE User_ID = ?
  `
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteUserRepository) GetByEmail(email string) (domain.User, error) {
	db := repo.db

	var user domain.User
	var unixSeconds int64
	query := `
  SELECT User_ID, Email, Hashed_Password, Registration_Date
  FROM User
  WHERE Email = ?
  `
	row := db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &unixSeconds)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.User{}, err
	}
	user.RegistrationDate = time.Unix(unixSeconds, 0)

	return user, nil
}

func (repo sqliteUserRepository) GetByID(id uint) (domain.User, error) {
	db := repo.db

	var user domain.User
	var unixSeconds int64
	query := `
  SELECT User_ID, Email, Hashed_Password, Registration_Date
  FROM User
  WHERE User_ID = ?
  `
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &unixSeconds)
	if err != nil {
		if err == sql.ErrNoRows {
			err = util.ErrEmptySelection
		}
		return domain.User{}, err
	}
	user.RegistrationDate = time.Unix(unixSeconds, 0)

	return user, nil
}

func (repo sqliteUserRepository) UpdateEmail(id uint, newEmail string) error {
	db := repo.db

	query := `
  UPDATE User
  SET Email = ?
  WHERE User_ID = ?
  `
	_, err := db.Exec(query, newEmail, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo sqliteUserRepository) UpdateHashedPassword(id uint, newHashedPassword string) error {
	db := repo.db

	query := `
  UPDATE User
  SET Hashed_Password = ?
  WHERE User_ID = ?
  `
	_, err := db.Exec(query, newHashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func NewSQLiteUserRepository(db *sql.DB) domain.UserRepository {
	return sqliteUserRepository{
		db: db,
	}
}
