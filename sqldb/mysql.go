package sqldb

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "max@/scrap")
	if err != nil {
		panic(err)
	}

	return db
}
