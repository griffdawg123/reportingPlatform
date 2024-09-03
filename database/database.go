package database

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Database struct {
    DB *pgx.Conn
}

func GetDatabase() (*pgx.Conn, error) {
    err := godotenv.Load("../.env")
    if err != nil {
        fmt.Errorf("Error loading environment variable: %v\n", err)
        return nil, err
    }
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        fmt.Errorf("Error connecting to database: %v\n", err)
        return nil, err
    }
    return conn, nil
}

func ApplySQL(db *pgx.Conn, schemaPath string) error {
    file, err := os.Open(schemaPath)
    if err != nil {
        return err
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(split_by_semicolon)
    for scanner.Scan() {
        query := scanner.Text()
        fmt.Println(query)
        _, err = db.Query(context.Background(), query)
        if err != nil {
            return err
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

func split_by_semicolon(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF {
        if data != nil {
            malformedFile := errors.New("File is malformed, file does not end with a ;")
            return 0, nil, malformedFile
        } else {
            return 0, nil, nil
        }
    }
    i := 0
    for data[i] != byte(';') {
        i++
    }
    return i, data[0:i+1], nil
}

