package database

import (
	"context"
	"fmt"
	"main/src/util"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection() (*pgxpool.Pool) {
	dbpool, err := pgxpool.New(context.Background(), util.GetEnvVar("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}