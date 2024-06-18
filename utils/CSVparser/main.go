package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("pgx", "postgres://zxcuser:zxc@127.0.0.1:5432/zxcdatabase?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	file, err := os.Open("binlist-data.csv")
	if err != nil {
		fmt.Println("Error on file open", err)
		return
	}
	defer file.Close()

	lineCount := count(file)

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error on file seek:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) >= 4 {
			_, err := db.Exec("INSERT INTO bin (bin, bank) VALUES ($1, $2)", fields[0], fields[4])
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("\rProgress: %d/%d", i, lineCount)
		}
	}
}

func count(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}
