package repositories_test

import (
	"testing"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	"github.com/yourname/reponame/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

// `SelectCommentList` 関数のテスト
func TestSelectCommentList(t *testing.T) {
	commentTests := []struct {
		testTitle string
		articleID int
		expected  int // 期待するコメント数
	}{
		{
			testTitle: "記事 1 のコメント取得",
			articleID: 1,
			expected:  len(testdata.CommentTestData1), // 事前にセットアップされたコメント数
		},
		{
			testTitle: "記事 2 のコメント取得 (コメントなし)",
			articleID: 2,
			expected:  0, // 記事 2 にはコメントなし
		},
	}

	for _, test := range commentTests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectCommentList(testDB, test.articleID)
			if err != nil {
				t.Fatalf("❌ %s 失敗: %v", test.testTitle, err)
			}

			if len(got) != test.expected {
				t.Errorf("❌ %s 失敗: 期待値 %d 件, 実際 %d 件", test.testTitle, test.expected, len(got))
			} else {
				t.Logf("✅ %s 成功: 期待値 %d 件, 実際 %d 件", test.testTitle, test.expected, len(got))
			}
		})
	}
}

// `InsertComment` 関数のテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "This is a test comment",
	}

	// データベースにコメントを追加
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatalf("❌ TestInsertComment 失敗: %v", err)
	}

	// ID が自動採番されているか確認
	if newComment.CommentID == 0 {
		t.Errorf("❌ TestInsertComment 失敗: 挿入されたコメント ID が 0")
	} else {
		t.Logf("✅ TestInsertComment 成功: 挿入されたコメント ID = %d", newComment.CommentID)
	}

	// 挿入したデータを削除
	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM comments WHERE comment_id = ?;
		`
		testDB.Exec(sqlStr, newComment.CommentID)
	})
}
