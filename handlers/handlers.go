package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// リクエストボディをJSONデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// JSONエンコードしてレスポンス
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reqArticle); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var article *models.Article
	if articleID == models.Article1.ID {
		article = &models.Article1
	} else if articleID == models.Article2.ID {
		article = &models.Article2
	} else {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// リクエストボディをデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// 該当する記事の NiceNum をインクリメント
	if reqArticle.ID == models.Article1.ID {
		models.Article1.NiceNum++
		reqArticle = models.Article1
	} else if reqArticle.ID == models.Article2.ID {
		models.Article2.NiceNum++
		reqArticle = models.Article2
	} else {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// JSONエンコードしてレスポンス
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reqArticle); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	// リクエストボディをデコード
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// コメントデータのレスポンス
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reqComment); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
