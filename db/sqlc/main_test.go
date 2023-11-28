package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	connString := os.Getenv("PG_CONN_STRING")

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer conn.Close(context.Background())

	testQueries = New(conn)

	result := m.Run()

	os.Exit(result)
}
