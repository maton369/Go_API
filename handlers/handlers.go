package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/services"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

// **📝 POST /article**
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// **リクエストボディのデコード**
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		log.Printf("❌ JSONのデコードに失敗しました: %v", err)
		http.Error(w, "❌ JSONのデコードに失敗しました", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// **サービス層を呼び出す**
	newArticle, err := services.PostArticleService(reqArticle)
	if err != nil {
		log.Printf("❌ 記事の投稿に失敗しました: %v", err)
		http.Error(w, "❌ 記事の投稿処理中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	// **レスポンスをJSONで返す**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArticle)
}

// **📝 GET /article/list**
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	page, err := strconv.Atoi(req.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1 // デフォルトのページ番号
	}

	// **サービス層を呼び出す**
	articles, err := services.GetArticleListService(page)
	if err != nil {
		log.Printf("❌ 記事一覧の取得に失敗しました: %v", err)
		http.Error(w, "❌ 記事一覧の取得中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	// **レスポンスをJSONで返す**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// **📝 GET /article/{id}**
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		log.Printf("❌ 記事IDが無効です: %v", err)
		http.Error(w, "❌ 記事IDが無効です", http.StatusBadRequest)
		return
	}

	// **サービス層を呼び出す**
	article, err := services.GetArticleService(articleID)
	if err != nil {
		log.Printf("❌ 記事の取得に失敗しました: %v", err)
		http.Error(w, "❌ 記事の取得中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	// **レスポンスをJSONで返す**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// **📝 POST /article/nice**
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// **リクエストボディのデコード**
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		log.Printf("❌ JSONのデコードに失敗しました: %v", err)
		http.Error(w, "❌ JSONのデコードに失敗しました", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// **サービス層を呼び出す**
	updatedArticle, err := services.PostNiceService(reqArticle.ID)
	if err != nil {
		log.Printf("❌ いいねの更新に失敗しました: %v", err)
		http.Error(w, "❌ いいねの更新中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	// **レスポンスをJSONで返す**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedArticle)
}

// **📝 POST /comment**
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	// **リクエストボディのデコード**
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		log.Printf("❌ JSONのデコードに失敗しました: %v", err)
		http.Error(w, "❌ JSONのデコードに失敗しました", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// **サービス層を呼び出してコメントをDBに追加**
	newComment, err := services.PostCommentService(reqComment)
	if err != nil {
		log.Printf("❌ コメントの投稿に失敗しました: %v", err)
		http.Error(w, "❌ コメントの投稿処理中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	// **レスポンスをJSONで返す**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newComment)
}
