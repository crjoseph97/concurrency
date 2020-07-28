package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func writeOp(c []string) {
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
func main() {
	csvfile, err := os.Open("localfile.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	var procmax int
	fmt.Printf("Enter the number logical processors required - ")
	fmt.Scanln(&procmax)
	runtime.GOMAXPROCS(procmax)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		go writeOp(record)
	}
	time.Sleep(100 * time.Millisecond)
}
