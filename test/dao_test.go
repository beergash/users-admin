package test

import (
	"came-users/app/db"
	"came-users/app/model"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var DB *sql.DB

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initDB() {
	fmt.Println("Initialize Database")
	db, err_conn := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		"postgres", "postgres", "cametest", "localhost"))
	DB = db

	check(err_conn)
	pwd, _ := os.Getwd()

	dat, err_file := ioutil.ReadFile(pwd + "/scripts/create_schema.sql")
	check(err_file)
	res, err_exec := db.Exec(string(dat))
	_ = res
	check(err_exec)
}

func TestAllUsersQry(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	users := dao.GetAllUsers(DB)
	assert.Equal(t, 5, len(users))
}

func TestUserByIdQry(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	u, found := dao.FindUserById(DB, 1)
	assert.True(t, found)
	assert.Equal(t, "Abc123", u.Username)
	assert.Equal(t, "Andrea", u.Name)
	assert.Equal(t, "Rossi", u.Surname)
}

func TestSaveUser(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	u := model.User{Username: "user_test", Name: "Kevin", Surname: "Reds", Password: "pwd", BirthDate: time.Now()}
	user_saved, err := dao.SaveUser(DB, &u)
	assert.Equal(t, "user_test", user_saved.Username)
	assert.Nil(t, err)
}

func TestSaveAlreadyExistingUser(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	u := model.User{Username: "Abc123", Name: "Kevin", Surname: "Reds", Password: "pwd", BirthDate: time.Now()}
	user_saved, err := dao.SaveUser(DB, &u)
	assert.NotNil(t, err)
	assert.Nil(t, user_saved)
}

func TestUpdateUser(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	u := model.User{Username: "Abc123", Name: "Michele", Surname: "Rossi", Password: "pwd", BirthDate: time.Now()}
	user_upd, err := dao.UpdateUser(DB, &u, 1)
	assert.Nil(t, err)
	assert.Equal(t, "Abc123", user_upd.Username)
	assert.Equal(t, "Michele", user_upd.Name)
}

func TestDeleteUser(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	err := dao.DeleteUser(DB, 2)
	assert.Nil(t, err)
}

func TestSearchUsers(t *testing.T) {
	initDB()
	dao := &db.Dao{}
	filter := model.UserSearchRequest{Name: "Andrea", Surname: "Rossi"}
	users, err := dao.GetAllUsersByFilters(DB, &filter)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Abc123", users[0].Username)
}
