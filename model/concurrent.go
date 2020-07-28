package model

import (
	"database/sql"
	"fmt"
	"log"
)

func WriteOp(c []string) {
	// fmt.Println(c)
	db, err := sql.Open("mysql", "root:1234@/test")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("insert into csv(first, second) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(c[0], c[1])
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert ID", id)
	defer db.Close()
}
