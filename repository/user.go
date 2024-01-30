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
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
  INSERT INTO User(Email, Hashed_Password, Registration_Date)
  VALUES (?,?,?)
  `
	res, err := tx.Exec(query, user.Email, user.HashedPassword, time.Now().Unix())
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			tx.Rollback()
			return 0, util.ErrRepeatedEntity
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

func (repo sqliteUserRepository) Delete(id uint) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  DELETE FROM User
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

func (repo sqliteUserRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	var unixSeconds int64
	query := `
  SELECT User_ID, Email, Hashed_Password, Registration_Date
  FROM User
  WHERE Email = ?
  `
	row := repo.db.QueryRow(query, email)
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
	var user domain.User
	var unixSeconds int64
	query := `
  SELECT User_ID, Email, Hashed_Password, Registration_Date
  FROM User
  WHERE User_ID = ?
  `
	row := repo.db.QueryRow(query, id)
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
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE User
  SET Email = ?
  WHERE User_ID = ?
  `
	_, err = tx.Exec(query, newEmail, id)
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

func (repo sqliteUserRepository) UpdateHashedPassword(id uint, newHashedPassword string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	query := `
  UPDATE User
  SET Hashed_Password = ?
  WHERE User_ID = ?
  `
	_, err = tx.Exec(query, newHashedPassword, id)
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

func NewSQLiteUserRepository(db *sql.DB) domain.UserRepository {
	return sqliteUserRepository{
		db: db,
	}
}
