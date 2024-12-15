package configs

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func main() {
	dsn := ""
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected")
}