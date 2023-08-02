package models

type Director struct {
	ID        string
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
