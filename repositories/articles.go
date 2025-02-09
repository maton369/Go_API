package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourname/reponame/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		INSERT INTO articles (title, contents, username, nice, created_at)
		VALUES (?, ?, ?, 0, NOW());
	`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, fmt.Errorf("failed to insert article: %w", err)
	}

	lastInsertID, _ := result.LastInsertId()
	article.ID = int(lastInsertID)
	return article, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice, created_at
		FROM articles
		LIMIT ? OFFSET ?;
	`
	limit := 5
	offset := (page - 1) * limit

	rows, err := db.Query(sqlStr, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to select article list: %w", err)
	}
	defer rows.Close()

	var articleArray []models.Article
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		if err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}
		articleArray = append(articleArray, article)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice, created_at
		FROM articles
		WHERE article_id = ?;
	`
	var article models.Article
	var createdTime sql.NullTime

	row := db.QueryRow(sqlStr, articleID)
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Article{}, fmt.Errorf("article not found")
		}
		return models.Article{}, fmt.Errorf("failed to select article detail: %w", err)
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	const sqlGetNice = `
		SELECT nice
		FROM articles
		WHERE article_id = ?;
	`
	var nicenum int
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Scan(&nicenum); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get current nice count: %w", err)
	}

	const sqlUpdateNice = `
		UPDATE articles
		SET nice = ?
		WHERE article_id = ?;
	`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update nice count: %w", err)
	}

	return tx.Commit()
}
