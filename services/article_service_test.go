package services_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/yourname/reponame/services"
)

var aSer *services.MyAppService

// **環境変数をロード**
func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("❌ .env ファイルの読み込みに失敗しました: %v", err)
	}
	log.Println("✅ .env ファイルを読み込みました")
}

func TestMain(m *testing.M) {
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

	// ✅ MySQL に接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println("DB接続エラー:", err)
		os.Exit(1)
	}
	defer db.Close() // ✅ DB 接続のクリーンアップ

	// ✅ サービスインスタンスを作成
	aSer = services.NewMyAppService(db)

	// ✅ ベンチマークテストの実行
	exitCode := m.Run()

	// ✅ テスト終了コードを返す
	os.Exit(exitCode)
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
