package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    source := "root:svc-ftth@tcp(localhost:3306)/svc_ftth"

    fmt.Printf("Connect to : %s", source)

    db, err := sql.Open("mysql", source)
    if err != nil {
        fmt.Println(err)
    }

    defer db.Close()
}
