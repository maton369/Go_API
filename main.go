package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/database"
	"github.com/yourname/reponame/handlers"
)

func main() {

	// データベースを初期化
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.CloseDB()

	// database.InsertArticle()

	// database.UpdateNiceCount(1)

	// 実際のアプリケーション処理（例としてログ出力）
	log.Println("Application is running...")

	r := mux.NewRouter()

	// 各ルートにメソッド制限を追加
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
