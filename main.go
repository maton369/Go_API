package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yourname/reponame/api"

	_ "github.com/go-sql-driver/mysql"
)

// **環境変数をロード**
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("❌ .env ファイルの読み込みに失敗しました: %v", err)
	}
	log.Println("✅ .env ファイルを読み込みました")
}

func main() {
	// **環境変数をロード**
	loadEnv()

	// **環境変数から DB 接続情報を取得**
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dbHost := os.Getenv("DB_HOST")

	// **必須の環境変数が設定されているかチェック**
	if dbUser == "" || dbPassword == "" || dbDatabase == "" || dbHost == "" {
		log.Fatalf("❌ 必須の環境変数が設定されていません")
	}

	// **DB 接続情報をフォーマット**
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)

	// **DB に接続**
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatalf("❌ データベース接続エラー: %v", err)
	}
	defer db.Close()

	r := api.NewRouter(db)

	// **サーバー起動**
	log.Println("✅ サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
