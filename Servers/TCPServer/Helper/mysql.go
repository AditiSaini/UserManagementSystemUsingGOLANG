package helper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	Structure "../Structure"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "users"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Show(username string) Structure.Profile {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Profile WHERE Nickname=?", username)
	if err != nil {
		panic(err.Error())
	}
	profile := Structure.Profile{}
	for selDB.Next() {
		var id int64
		var username string
		var nickname string
		var password string
		var picture sql.NullString
		err = selDB.Scan(&id, &username, &nickname, &password, &picture)
		if err != nil {
			panic(err.Error())
		}
		profile.ID = id
		profile.Username = username
		profile.Nickname = nickname
		if picture.Valid {
			profile.ProfilePicture = picture.String
		} else {
			profile.ProfilePicture = ""
		}

		profile.Password = password
		profile.Valid = true
	}
	defer db.Close()
	return profile
}

func Update(username string, name string) {
	db := dbConn()
	insForm, err := db.Prepare("UPDATE Profile SET Name=? WHERE Nickname=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(name, username)
	log.Println("UPDATE: Name: " + name + " of User: " + username)
	defer db.Close()
}
