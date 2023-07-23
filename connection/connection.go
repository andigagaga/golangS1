package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DataBaseConnect() {

	// connection string
	databaseurl := "postgres://postgres:01082003@localhost:5434/b47-s1"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseurl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("seccesfully connect to data base")
}
