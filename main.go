package main

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

/*
schema

CREATE TABLE `users` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	`created_at` datetime NOT NULL,
	`updated_at` datetime NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `posts` (
	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` int(11) unsigned NOT NULL,
	`title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	`created_at` datetime NOT NULL,
	`updated_at` datetime NOT NULL,
	PRIMARY KEY (`id`),
	CONSTRAINT `posts_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
*/

func main() {
	c := mysql.Config{
		DBName:    "db",
		User:      "soda",
		Passwd:    "soda",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.UTC,
	}
	var err error
	db, err = sql.Open("mysql", c.FormatDSN())
	defer db.Close()
	if err != nil {
		panic(err)
	}
	users, err := SelectUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		println(user.Name)
	}
}

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SelectUsers() ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var id int
		var name string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			panic(err)
		}
		users = append(users, User{id, name, createdAt, updatedAt})
	}
	return users, nil
}
