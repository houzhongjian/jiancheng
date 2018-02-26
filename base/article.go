package base

//Article 文章.
type Article struct {
	Model
	MenuId      int64  `json:"menu_id"`
	Title       string `json:"title"`
	Pic         string `json:"pic"`
	Author      string `json:"author"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
