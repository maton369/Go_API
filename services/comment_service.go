package services

import (
	"fmt"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// **📝 PostCommentService** (コメント投稿)
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, fmt.Errorf("❌ [Service: PostCommentService] データベース接続に失敗しました: %w", err)
	}
	defer db.Close()

	// コメントをデータベースに挿入
	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, fmt.Errorf("❌ [Service: PostCommentService] コメントの投稿に失敗しました: %w", err)
	}

	return newComment, nil
}
