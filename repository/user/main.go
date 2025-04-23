package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func queryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}
}
