package services

import (
	"fmt"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// **📝 GetArticleService** (記事詳細取得)
func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: GetArticleService] データベース接続に失敗しました: %w", err)
	}
	defer db.Close()

	// 1. 記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: GetArticleService] 記事詳細の取得に失敗しました: %w", err)
	}

	// 2. コメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: GetArticleService] コメント一覧の取得に失敗しました: %w", err)
	}

	// 3. 記事にコメントを紐付ける
	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// **📝 PostArticleService** (記事投稿)
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: PostArticleService] データベース接続に失敗しました: %w", err)
	}
	defer db.Close()

	// 記事をデータベースに挿入
	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: PostArticleService] 記事の投稿に失敗しました: %w", err)
	}

	return newArticle, nil
}

// **📝 GetArticleListService** (記事一覧取得)
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, fmt.Errorf("❌ [Service: GetArticleListService] データベース接続に失敗しました: %w", err)
	}
	defer db.Close()

	// 記事一覧を取得
	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, fmt.Errorf("❌ [Service: GetArticleListService] 記事一覧の取得に失敗しました: %w", err)
	}

	return articles, nil
}

// **📝 PostNiceService** (いいね更新)
func PostNiceService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: PostNiceService] データベース接続に失敗しました: %w", err)
	}
	defer db.Close()

	// `nice` 数を更新
	err = repositories.UpdateNiceNum(db, articleID)
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: PostNiceService] いいねの更新に失敗しました: %w", err)
	}

	// 更新後のデータを取得
	updatedArticle, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, fmt.Errorf("❌ [Service: PostNiceService] 記事の再取得に失敗しました: %w", err)
	}

	return updatedArticle, nil
}
