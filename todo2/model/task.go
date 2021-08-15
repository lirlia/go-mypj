package task

type Task struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Done  bool   `json:"done,omitempty"`
}
