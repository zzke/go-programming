package main

import (
	"database/sql"
	"log"
)

func main() {
	// db是一个sql.DB类型的对象，该对象线程安全，且内部已包含了一个连接池
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from where id = ?", 1)
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
