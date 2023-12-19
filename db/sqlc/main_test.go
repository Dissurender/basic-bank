package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	testStore   Store
	testQueries Queries
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	connString := os.Getenv("PG_CONN_STRING")

	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testStore = NewStore(connPool)

	result := m.Run()

	os.Exit(result)
}
