package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonhador82/go-sql-tut/tutorial"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/sandbox1") // это только абстракция, в этот момент соединение не откроется
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	tutorial.FetchData(db)
	tutorial.PrepareQuery(db)
	tutorial.GetOne(db)
	tutorial.UpdateQuery(db)
	tutorial.UpdateInTransaction(db)

	defer db.Close()
}
