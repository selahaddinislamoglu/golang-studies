package model

type ArticleSummary struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
}

type Article struct {
	ArticleSummary
	Content string `json:"content"`
}
