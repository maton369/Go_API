-- 記事データを格納するテーブル
CREATE TABLE IF NOT EXISTS articles (
    article_id  INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(100) NOT NULL,
    contents    TEXT NOT NULL,
    username    VARCHAR(100) NOT NULL,
    nice        INTEGER NOT NULL,
    created_at  DATETIME
);

-- コメントデータを格納するテーブル
CREATE TABLE IF NOT EXISTS comments (
    comment_id  INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    article_id  INTEGER UNSIGNED NOT NULL,
    message     TEXT NOT NULL,
    created_at  DATETIME,
    FOREIGN KEY (article_id) REFERENCES articles(article_id)
);

-- 記事データの挿入
INSERT INTO articles (title, contents, username, nice, created_at) VALUES
    ('firstPost', 'This is my first blog', 'saki', 2, NOW()),
    ('2nd', 'Second blog post', 'saki', 4, NOW());

-- コメントデータの挿入
INSERT INTO comments (article_id, message, created_at) VALUES
    (1, '1st comment yeah', NOW()),
    (1, 'welcome', NOW());
