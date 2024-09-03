package database_test

import (
	"context"
	"fmt"
	"os"
	"reportingPlatform/database"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvFile(t *testing.T) {
    err := godotenv.Load("../.env")
    if err != nil {
        t.Errorf("Error loading env file")
    }
    test_env := os.Getenv("TEST")
    if test_env != "helloworld123" {
        t.Errorf("Test env value is incorrect")
    }
}

func TestDatabaseConnected(t *testing.T) {
    db, err := database.GetDatabase()
    defer db.Close(context.Background())
    if err != nil {
        t.Errorf("Error connecting to database")
    }
}

func TestQuery(t *testing.T) {
    err := godotenv.Load("../.env")
    db, err := database.GetDatabase()
    defer db.Close(context.Background())
    if err != nil {
        t.Errorf("Error connecting to database")
    }
    test_schema_path := os.Getenv("TEST_SCHEMA_PATH")
    fmt.Println(test_schema_path)
    test_string := "Hello, World!"
    var res string;
    err = db.QueryRow(context.Background(), "SELECT $1;", test_string).Scan(&res);
    if err != nil {
        t.Errorf("Database query returned error: %v", err)
    }
    if res != test_string {
        t.Errorf(
            "Database did not return required string: needed: %s, received: %s",
            test_string,
            res,
        )
    }
}
