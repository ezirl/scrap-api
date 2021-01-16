package sqldb

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root@/scrap")
	if err != nil {
		panic(err)
	}

	return db
}
