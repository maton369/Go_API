package repositories_test

import (
	"testing"

	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
	"github.com/yourname/reponame/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

// `SelectArticleList` 関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	} else {
		t.Logf("✅ TestSelectArticleList passed: expected %d articles, got %d", expectedNum, num)
	}
}

// `SelectArticleDetail` 関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID ||
				got.Title != test.expected.Title ||
				got.Contents != test.expected.Contents ||
				got.UserName != test.expected.UserName ||
				got.NiceNum != test.expected.NiceNum {
				t.Errorf("❌ %s failed: expected %+v, got %+v", test.testTitle, test.expected, got)
			} else {
				t.Logf("✅ %s passed: expected %+v, got %+v", test.testTitle, test.expected, got)
			}
		})
	}
}

// `InsertArticle` 関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	if newArticle.ID != expectedArticleNum {
		t.Errorf("❌ TestInsertArticle failed: expected ID %d, got %d", expectedArticleNum, newArticle.ID)
	} else {
		t.Logf("✅ TestInsertArticle passed: expected ID %d, got %d", expectedArticleNum, newArticle.ID)
	}

	// クリーンアップ処理
	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM articles
			WHERE title = ? AND contents = ? AND username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

// `UpdateNiceNum` 関数のテスト
func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	err := repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	got, _ := repositories.SelectArticleDetail(testDB, articleID)

	if got.NiceNum-testdata.ArticleTestData[articleID-1].NiceNum != 1 {
		t.Errorf("❌ TestUpdateNiceNum failed: expected %d, got %d",
			testdata.ArticleTestData[articleID].NiceNum,
			got.NiceNum)
	} else {
		t.Logf("✅ TestUpdateNiceNum passed: expected %d, got %d",
			testdata.ArticleTestData[articleID].NiceNum,
			got.NiceNum)
	}
}
