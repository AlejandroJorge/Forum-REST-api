package repository

import (
	"database/sql"
	"time"

	"github.com/AlejandroJorge/forum-rest-api/domain"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteUserRepository struct {
	db *sql.DB
}

// Returns the ID of the created user and can return ErrRepeatedEntity
func (repo sqliteUserRepository) Create(email, hashedPassword string) (uint, error) {
	db := repo.db

	query := `
  INSERT INTO User(Email, Hashed_Password, Registration_Date)
  VALUES (?,?,?)
  `
	res, err := db.Exec(query, email, hashedPassword, time.Now().Unix())
	if sqliteErr, ok := err.(sqlite3.Error); ok {
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
func (repo sqliteUserRepository) Delete(id uint) error {
	db := repo.db

	query := `
  DELETE FROM User
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

// Returns a valid user and can return ErrEmptySelection
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
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.User{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.User{}, ErrUnknown
	}

	user.RegistrationDate = time.Unix(unixSeconds, 0)

	return user, nil
}

// Returns a valid user and can return ErrEmptySelection
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
	if err == sql.ErrNoRows {
		logging.LogRepositoryError(ErrEmptySelection)
		return domain.User{}, ErrEmptySelection
	}
	if err != nil {
		logging.LogUnexpectedRepositoryError(err)
		return domain.User{}, ErrUnknown
	}

	user.RegistrationDate = time.Unix(unixSeconds, 0)

	return user, nil
}

// Can return ErrNoRowsAffected
func (repo sqliteUserRepository) UpdateEmail(id uint, newEmail string) error {
	db := repo.db

	query := `
  UPDATE User
  SET Email = ?
  WHERE User_ID = ?
  `
	res, err := db.Exec(query, newEmail, id)
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
func (repo sqliteUserRepository) UpdateHashedPassword(id uint, newHashedPassword string) error {
	db := repo.db

	query := `
  UPDATE User
  SET Hashed_Password = ?
  WHERE User_ID = ?
  `
	res, err := db.Exec(query, newHashedPassword, id)
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

func NewSQLiteUserRepository(db *sql.DB) domain.UserRepository {
	return sqliteUserRepository{
		db: db,
	}
}
