package main

import (
	"context"
	"fmt"
	"os"
	"reportingPlatform/database"
)


func main() {
    db, err := database.GetDatabase()
    if err != nil {
        fmt.Println("Error getting database connection:", err)
        os.Exit(1)
    }
    defer db.Close(context.Background())

    var msg string
    err = db.QueryRow(context.Background(), "SELECT firstname from teachers limit 1").Scan(&msg)
    if err != nil {
        fmt.Println("Error Querying database:", err)
        os.Exit(1)
    }

    fmt.Println(msg)
}

