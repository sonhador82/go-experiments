package tutorial

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	id   int
	name string
)

// FetchData - query for data
func FetchData(db *sql.DB) {
	rows, err := db.Query("select `id`, `name` from `testdata1` where `id`=?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

// PrepareQuery preare query
func PrepareQuery(db *sql.DB) {
	stmt, err := db.Prepare("select `id`, `name` from testdata1 where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// GetOne - get single row
func GetOne(db *sql.DB) {
	var name string
	err := db.QueryRow("select `name` from testdata1 where `id`= ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

// UpdateQuery example
func UpdateQuery(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO `testdata1`(`name`) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Don")
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %d, Rows: %d\n", lastID, rowCnt)
}

// UpdateInTransaction - transaction example
func UpdateInTransaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO `testdata1`(`name`) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Bob Jones")
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		log.Println("rolling back")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last id in transaction: %d", lastID)
}
