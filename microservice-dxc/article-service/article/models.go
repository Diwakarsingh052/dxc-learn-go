package article

type Article struct {
	ArticleId   string `json:"article_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorEmail string `json:"author_email"`
}

type NewArticle struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorEmail string `json:"author_email"`
}
