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

	return nil
}

// 記事のいいね数を +1 する関数
func UpdateNiceCount(articleID int) error {
	// トランザクションの開始
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 現在のいいね数を取得するクエリ
	const sqlGetNice = `
		SELECT nice
		FROM articles
		WHERE article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, articleID)

	// 現在のいいね数を取得
	var nicenum int
	if err := row.Scan(&nicenum); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get current nice count: %w", err)
	}

	// いいね数を +1 する更新処理
	const sqlUpdateNice = `
		UPDATE articles
		SET nice = ?
		WHERE article_id = ?;
	`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update nice count: %w", err)
	}

	// コミットして処理を確定
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("Nice count updated successfully for article_id: %d", articleID)
	return nil
}

func InsertArticle() {
	// 挿入する記事データ
	article := models.Article{
		Title:    "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "saki",
	}

	// SQLクエリ
	const sqlStr = `
		INSERT INTO articles (title, contents, username, nice, created_at)
		VALUES (?, ?, ?, 0, NOW());
	`

	// データを挿入
	result, err := DB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		log.Fatalf("Failed to insert article: %v", err)
	}

	// 挿入結果の確認
	lastInsertID, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Article inserted successfully! ID: %d, Rows affected: %d\n", lastInsertID, rowsAffected)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
