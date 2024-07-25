package types

type Message struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	UserType string `json:"userType"`
}
