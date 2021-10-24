package dbase

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"
)

type Updater struct {
	db         *sql.DB
	scriptPath string
}

func (u *Updater) Db(db *sql.DB) *Updater {
	u.db = db
	return u
}

func (u *Updater) ScriptPath(scriptPath string) *Updater {
	u.scriptPath = scriptPath
	return u
}

func (u *Updater) Update() error {
	files, err := ioutil.ReadDir(u.scriptPath)
	if err != nil {
		log.Fatal(err)
	}

	var lastAppliedScript string
	var appliedScriptsCount int
	sqlStatement := `SELECT COUNT(scriptName) FROM applied_scripts`
	err = u.db.QueryRow(sqlStatement).Scan(&appliedScriptsCount)
	if err != nil {
		log.Fatal(err)
	}
	if appliedScriptsCount > 0 {
		sqlStatement = `SELECT MAX(scriptName) FROM applied_scripts`
		err = u.db.QueryRow(sqlStatement).Scan(&lastAppliedScript)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range files {
		if lastAppliedScript == "" || lastAppliedScript < file.Name() {
			data, err := ioutil.ReadFile(u.scriptPath + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			scripts := strings.Split(string(data), ";")
			for _, script := range scripts {
				if strings.Trim(script, " \n\t\r") != "" {
					_, err = u.db.Query(script)

					if err != nil {
						log.Fatal(err)
					}
				}
			}

			sqlStatement := `insert into applied_scripts values(?)`
			_, err = u.db.Exec(sqlStatement, file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
