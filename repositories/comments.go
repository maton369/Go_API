package repositories

import (
	"database/sql"
	"fmt"

	"github.com/yourname/reponame/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at)
		VALUES (?, ?, NOW());
	`
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	lastInsertID, _ := result.LastInsertId()
	comment.CommentID = int(lastInsertID)
	return comment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT comment_id, article_id, message, created_at
		FROM comments
		WHERE article_id = ?;
	`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, fmt.Errorf("failed to select comment list: %w", err)
	}
	defer rows.Close()

	var commentArray []models.Comment
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		if err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
