package models

type IssueTag struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Color string `json:"color"`
}

type CreateTagBody struct {
	Project int    `json:"projectId" binding:"required"`
	Text    string `json:"text" binding:"required"`
	Color   string `json:"color" binding:"required"`
}
