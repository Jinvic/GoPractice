package define

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CompletedAt string `json:"completed_at,omitempty"`
}

type TodoListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetListRsp struct {
	List  []TodoListItem `json:"list"`
	Total int            `json:"total"`
}
