package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"git.hifx.in/concurrency/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	csvfile, err := os.Open("../localfile.csv")
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
		go model.WriteOp(record)
	}
	time.Sleep(100 * time.Millisecond)
}
