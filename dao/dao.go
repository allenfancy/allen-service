package dao

import (
	"allen-service/model"
	"database/sql"
	"log"
)

type Dao struct {
	db *sql.DB
}

func New() (d *Dao) {
	db, err := sql.Open("mysql", "root:123@/test")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	d = &Dao{
		db: db,
	}
	return
}

func (d *Dao) Insert(user *model.UserInfo) (affected int64, err error) {
	stmt, err := d.db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("guotie", "guotie")
	stmt.Exec("testuser", "123123")
	return
}

func (d *Dao) Update() (affected int64, err error) {
	return
}

func (d *Dao) Delete() (affected int64, err error) {
	return
}

func (d *Dao) Query() (affected int64, err error) {

	return
}
