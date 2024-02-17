package models

type Task struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Date        string `json:"date,omitempty"`
	Complete    bool   `json:"complete,omitempty"`
}
