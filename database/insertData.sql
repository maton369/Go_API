-- 記事データ 2 つを挿入
INSERT INTO articles (title, contents, username, nice, created_at) 
VALUES ('firstPost', 'This is my first blog', 'saki', 2, NOW());

INSERT INTO articles (title, contents, username, nice, created_at) 
VALUES ('2nd', 'Second blog post', 'saki', 4, NOW());

-- コメントデータ 2 つを挿入
INSERT INTO comments (article_id, message, created_at) 
VALUES (1, '1st comment yeah', NOW());

INSERT INTO comments (article_id, message, created_at) 
VALUES (1, 'welcome', NOW());
