package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func read_from_db(num_lines int) {
	db, err := sql.Open("sqlite3", "./air_quality.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM Observations LIMIT ?", num_lines)
	if err != nil {
		log.Fatal(err)
	}

	var obs_time_int int64
	var obs_time time.Time
	var co2 int
	var temp float32
	var humidity float32

	for rows.Next() {
		err = rows.Scan(&obs_time_int, &co2, &temp, &humidity)
		if err != nil {
			log.Fatal(err)
		}

		obs_time = time.Unix(obs_time_int, 0)

		fmt.Printf("%v, %d, %v, %v\n", obs_time, co2, temp, humidity)
	}
}

func read_lines(file_path string, num_lines int) [][]string {
	// open file
	f, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	lines := make([][]string, num_lines)
	num_lines_read := 0
	csvReader := csv.NewReader(f)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		date_string := line[0]

		cutoff_date, err := time.Parse("01/02/06 15:04:05", "12/20/21 23:39:33")
		if err != nil {
			panic(err)
		}
		obs_date, err := time.Parse("01/02/06 15:04:05", date_string)
		if err != nil {
			panic(err)
		}

		fmt.Print(obs_date.Format("01/02/06 15:04:05"))
		if obs_date.Before(cutoff_date) {
			fmt.Print(" - Before\n")
		} else {
			fmt.Print(" - After\n")
			break
		}

		lines[num_lines_read] = line
		num_lines_read++

		if num_lines_read >= num_lines {
			break
		}
	}

	fmt.Printf("Number of lines read: %d\n", num_lines_read)

	return lines
}

func main() {
	fmt.Println("Hello, World!")

	// lines := read_lines("air_quality.csv", 12)
	// fmt.Println("2d: ", lines)

	read_from_db(40)
}
