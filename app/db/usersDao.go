package db

import (
	"came-users/app/model"
	"database/sql"
	_ "errors"
	"log"

	_ "github.com/lib/pq"
)

const selectAllUsersQry string = "SELECT ID, USERNAME, PASSWORD, NAME, SURNAME, BIRTH_DATE FROM USERS"
const selectUsersWithFiltersQry string = "SELECT ID, USERNAME, PASSWORD, NAME, SURNAME, BIRTH_DATE FROM USERS WHERE NAME = $1 AND SURNAME = $2"
const insertUserStatement string = "INSERT INTO USERS (USERNAME, PASSWORD, NAME, SURNAME, BIRTH_DATE) VALUES ($1, $2, $3, $4, $5)"
const selectUserByIdQry string = "SELECT ID, USERNAME, PASSWORD, NAME, SURNAME, BIRTH_DATE FROM USERS WHERE id=$1"
const deleteUserStmt string = "DELETE FROM USERS WHERE ID = $1"
const updateUserStmt string = "UPDATE USERS SET USERNAME = $2 , PASSWORD = $3 , NAME = $4 , SURNAME = $5, BIRTH_DATE = $6 WHERE ID = $1"

type Dao struct{}

func (a *Dao) GetAllUsers(db *sql.DB) []model.User {
	rows, err := db.Query(selectAllUsersQry)
	if err != nil {
		log.Fatal(err)
	}
	users := []model.User{}

	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Name, &u.Surname, &u.BirthDate); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	defer rows.Close()
	return users
}

func (a *Dao) FindUserById(db *sql.DB, id int64) (model.User, bool) {
	log.Printf("searching user with id %d ", id)
	u := model.User{}
	row := db.QueryRow(selectUserByIdQry, id)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Name, &u.Surname, &u.BirthDate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No found rows for id %d", id)
			return u, false
		} else {
			panic(err)
		}
	}
	return u, true
}

// insert new user
func (a *Dao) SaveUser(db *sql.DB, u *model.User) (*model.User, error) {
	_, err := db.Exec(insertUserStatement, u.Username, u.Password, u.Name, u.Surname, u.BirthDate)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %s correctly added", u.Username)
	return u, nil

}

// update user
func (a *Dao) UpdateUser(db *sql.DB, u *model.User, id int64) (*model.User, error) {
	_, err := db.Exec(updateUserStmt, id, u.Username, u.Password, u.Name, u.Surname, u.BirthDate)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %s correctly updated", u.Username)
	return u, nil

}

// delete user
func (a *Dao) DeleteUser(db *sql.DB, id int64) error {
	_, err := db.Exec(deleteUserStmt, id)
	if err != nil {
		return err
	}
	log.Printf("userid: %d correctly deleted", id)
	return nil

}

func (a *Dao) GetAllUsersByFilters(db *sql.DB, filter *model.UserSearchRequest) ([]model.User, error) {
	rows, err := db.Query(selectUsersWithFiltersQry, filter.Name, filter.Surname)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	users := []model.User{}

	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Name, &u.Surname, &u.BirthDate); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, u)
	}
	defer rows.Close()
	return users, nil
}
