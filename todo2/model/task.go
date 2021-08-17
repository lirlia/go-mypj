package task

type Task struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Done  bool   `json:"done,omitempty"`
}
