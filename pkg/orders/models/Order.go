package models

type Order struct {
	Id         string   `json:"id"`
	CustomerId int      `json:"customerId"`
	Items      []string `json:"items"`
	Total      float64  `json:"total"`
}
