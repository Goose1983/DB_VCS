package main

import (
	dbase "DB_VCS/updater"
	"database/sql"
	"log"
)

func main() {
	u := dbase.Updater{}
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/urms")
	if err != nil {
		log.Println(err)
	}
	err = u.Db(db).ScriptPath("./migrations/prod").Update()
	if err != nil {
		log.Println(err)
	}
}
