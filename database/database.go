package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	dbUser := "docker"
	dbPassword := "Docker123!"
	dbDatabase := "sampledb"
	dbHost := "127.0.0.1" // Docker Compose のサービス名を使用

	// DB接続文字列を構築
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)

	// データベース接続を初期化
	var err error
	DB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 接続確認
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the database successfully")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
