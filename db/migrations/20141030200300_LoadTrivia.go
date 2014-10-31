package main

import (
	"bufio"
	"database/sql"
	"os"
	"strings"
)

// Up is executed when this migration is applied
func Up_20141030200300(txn *sql.Tx) {
	file, _ := os.Open("trivia.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "`")
		txn.Exec("insert into trivia (question, answer) values ($1,$2)", parts[0], parts[1])
	}
}

// Down is executed when this migration is rolled back
func Down_20141030200300(txn *sql.Tx) {
	txn.Exec("truncate table trivia")
}
