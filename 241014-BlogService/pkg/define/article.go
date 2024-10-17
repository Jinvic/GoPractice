package define

type ArticleInfo struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ArticleCreateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleEditReq struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

