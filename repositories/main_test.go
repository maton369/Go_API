package repositories_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// `testDB`: テスト全体で使用するDB
var testDB *sql.DB

// `.env` をロードし、環境変数を取得
func loadEnv() error {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("⚠️ No .env file found: %v", err)
	}
	return nil
}

// **DB 接続情報を環境変数から取得**
var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbHost     string
	dbConn     string
)

// **環境変数をセットアップ**
func setupEnv() error {
	loadEnv() // `.env` をロード

	// **環境変数を取得**
	dbUser = os.Getenv("USERNAME")
	dbPassword = os.Getenv("USERPASS")
	dbDatabase = os.Getenv("DATABASE")
	dbHost = os.Getenv("DB_HOST")

	// **必須の環境変数チェック**
	if dbUser == "" || dbPassword == "" || dbDatabase == "" || dbHost == "" {
		return fmt.Errorf("❌ 必須の環境変数が設定されていません")
	}

	// **DB接続情報をフォーマット**
	dbConn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbDatabase)

	return nil
}

// **DB に接続する関数**
func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return fmt.Errorf("❌ データベース接続エラー: %w", err)
	}

	// 接続確認
	if err := testDB.Ping(); err != nil {
		return fmt.Errorf("❌ データベースPINGエラー: %w", err)
	}

	log.Println("✅ テスト用データベースに接続しました")
	return nil
}

// **テストデータをセットアップ**
func setupTestData() error {
	cmd := exec.Command("mysql",
		"-h", dbHost,
		"-u", dbUser,
		"-p"+dbPassword, dbDatabase,
		"-e", "source ./testdata/setupDB.sql",
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("❌ テストデータセットアップエラー: %w", err)
	}
	log.Println("✅ テストデータをセットアップしました")
	return nil
}

// **テストデータを削除**
func cleanupDB() error {
	cmd := exec.Command("mysql",
		"-h", dbHost,
		"-u", dbUser,
		"-p"+dbPassword, dbDatabase,
		"-e", "source ./testdata/cleanupDB.sql",
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("❌ クリーンアップエラー: %w", err)
	}
	log.Println("🗑️ データベースをクリーンアップしました")
	return nil
}

// **全テスト共通の前処理**
func setupDB() error {
	// 環境変数をセットアップ
	if err := setupEnv(); err != nil {
		return err
	}

	// DB に接続
	if err := connectDB(); err != nil {
		return err
	}

	// **データベースをクリーンアップ**
	if err := cleanupDB(); err != nil {
		return err
	}

	// **テストデータのセットアップ**
	if err := setupTestData(); err != nil {
		return err
	}

	return nil
}

// **全テスト共通の後処理**
func teardown() {
	cleanupDB() // **テストデータを削除**
	if testDB != nil {
		testDB.Close()
		log.Println("🛑 テストデータベースの接続を閉じました")
	}
}

// `TestMain` は全テストのセットアップとクリーンアップを担当する
func TestMain(m *testing.M) {
	// テスト用のデータベースをセットアップ
	if err := setupDB(); err != nil {
		log.Fatalf("❌ テストセットアップに失敗: %v", err)
	}

	// テストを実行
	code := m.Run()

	// 終了時にデータベース接続を閉じる
	teardown()

	// テストの終了コードで終了
	os.Exit(code)
}
