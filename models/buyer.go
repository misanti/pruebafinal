package models

type Buyer struct {
	ID       int    `json:"id"`
	Code     int    `json:"code"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
