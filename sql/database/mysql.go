package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBPing() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_web_scout")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping 的时候会与数据库建立连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("connected\n")
}

func MysqlDemoCode() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_web_scout")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // 创建表
		query := `
        CREATE TABLE users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (id)
        );`
		if _, err = db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // 插入新数据
		username := "Joshua"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // 查询单个用户
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // 查询所有用户
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{ // 删除
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
