package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	Structure "servers/Structure"
	Constants "servers/internal"
)

func dbConn() (db *sql.DB) {
	dbDriver := Constants.DB_DRIVER
	dbUser := Constants.DB_USER
	dbPass := Constants.DB_PASS
	dbName := Constants.DB_NAME
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func Show(username string) Structure.Profile {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Profile WHERE Nickname=?", username)
	if err != nil {
		fmt.Println(err)
	}
	profile := Structure.Profile{}
	for selDB.Next() {
		var id int64
		var username string
		var nickname string
		var password []byte
		var picture sql.NullString
		err = selDB.Scan(&id, &username, &nickname, &password, &picture)
		if err != nil {
			fmt.Println(err)
		}
		profile.ID = id
		profile.Username = username
		profile.Nickname = nickname
		profile.Password = password
		if picture.Valid {
			profile.ImageRef = picture.String
		} else {
			profile.ImageRef = ""
		}
		profile.Valid = true
	}
	defer db.Close()
	return profile
}

func UpdateUserProfile(profile *Structure.Profile) (bool, error) {
	name := profile.Username
	nickname := profile.Nickname
	password := profile.Password
	imageRef := profile.ImageRef
	db := dbConn()
	insForm, err := db.Prepare("UPDATE Profile SET Name=?, Password=?, ImageRef=? WHERE Nickname=?")
	if err != nil {
		return false, err
	}
	insForm.Exec(name, password, imageRef, nickname)
	log.Println("UPDATED: Profile of user: " + nickname)
	defer db.Close()
	return true, nil
}

func UpdateProfile(username string, name string) (bool, error) {
	db := dbConn()
	insForm, err := db.Prepare("UPDATE Profile SET Name=? WHERE Nickname=?")
	if err != nil {
		return false, err
	}
	insForm.Exec(name, username)
	defer db.Close()
	return true, nil
}

func UpdatePassword(password []byte, username string) (bool, error) {
	db := dbConn()
	insForm, err := db.Prepare("UPDATE Profile SET Password=? WHERE Nickname=?")
	if err != nil {
		return false, err
	}
	insForm.Exec(password, username)
	log.Println("UPDATED: Password of user: " + username)
	defer db.Close()
	return true, nil
}

func UpdateImageRef(imageRef string, username string) (bool, error) {
	db := dbConn()
	insForm, err := db.Prepare("UPDATE Profile SET ImageRef=? WHERE Nickname=?")
	if err != nil {
		return false, err
	}
	insForm.Exec(imageRef, username)
	log.Println("UPDATED: image ref of user: " + username)
	defer db.Close()
	return true, nil
}
