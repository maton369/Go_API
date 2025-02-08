package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/yourname/reponame/models"
)

var DB *sql.DB

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func InitDB() error {
	// .env ファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// 環境変数からデータベース接続情報を取得
	dbUser := getEnv("USERNAME")
	dbPassword := getEnv("USERPASS")
	dbDatabase := getEnv("DATABASE")
	dbHost := getEnv("DB_HOST")

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

	// 特定の記事を取得するクエリ
	articleID := 1
	const sqlStr = `
		SELECT article_id, title, contents, username, nice, created_at
		FROM articles
		WHERE article_id = ?;
	`

	// クエリ実行（単一記事取得）
	row := DB.QueryRow(sqlStr, articleID)

	var article models.Article
	var createdTime sql.NullTime

	// データをスキャン
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No article found with the given ID")
			return nil
		}
		return fmt.Errorf("failed to scan row: %w", err)
	}

	// NULLチェックを行い、値を設定
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	// 結果を出力
	fmt.Printf("Article: %+v\n", article)

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
