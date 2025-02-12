package testdata

import "github.com/yourname/reponame/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  4,
	},
}

// コメントデータ
var (
	CommentTestData1 = []models.Comment{
		{
			CommentID: 1,
			ArticleID: 1,
			Message:   "test comment1",
		},
		{
			CommentID: 2,
			ArticleID: 1,
			Message:   "second comment",
		},
	}

	CommentTestData2 = []models.Comment{} // 記事2にはコメントなし
)
