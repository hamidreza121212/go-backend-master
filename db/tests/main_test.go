package testutil

import (
    "database/sql"
    "log"
    "os"
    "testing"

    db "go-backend-master/db/sqlc"
	_ "github.com/lib/pq"
)

const (
    dbDriver = "postgres"
    dbSource = "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *db.Queries

// TestMain sets up the database connection and initializes test queries.
func TestMain(m *testing.M) {
    conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    // Ensure the database connection is closed after tests run
    defer func() {
        if err = conn.Close(); err != nil {
            log.Fatal("cannot close db connection:", err)
        }
    }()

    testQueries = db.New(conn)

    // Run tests
    os.Exit(m.Run())
}
