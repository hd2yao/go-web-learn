package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func MysqlDemoCode() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_web_scout")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}
