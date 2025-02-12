package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// `.env` をロードする関数
func loadEnv() {
	paths := []string{".env", "../.env"}

	var err error
	for _, path := range paths {
		if _, err = os.Stat(path); err == nil {
			if err = godotenv.Load(path); err == nil {
				log.Printf("✅ 環境変数を読み込みました: %s", path)
				return
			}
		}
	}

	// `.env` が見つからなかった場合のログ
	log.Printf("⚠️ .env ファイルが見つかりません: %v", err)
}

// **DB に接続する関数**
func connectDB() (*sql.DB, error) {
	// `.env` をロード
	loadEnv()

	// `.env` のロード後に環境変数を取得
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dbHost := os.Getenv("DB_HOST")

	// 必須の環境変数が設定されているかチェック
	if dbUser == "" || dbPassword == "" || dbDatabase == "" || dbHost == "" {
		log.Println("DB_HOST:", dbHost)
		log.Println("DB_USER:", dbUser)
		log.Println("DB_PASSWORD:", dbPassword)
		log.Println("DB_NAME:", dbDatabase)
		return nil, fmt.Errorf("❌ 環境変数が正しく設定されていません")
	}

	// DB接続情報をフォーマット
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)

	// DB接続を試みる
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, fmt.Errorf("❌ データベース接続エラー: %w", err)
	}

	// 接続確認
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ データベースPINGエラー: %w", err)
	}

	log.Println("✅ データベースに接続しました")
	return db, nil
}
