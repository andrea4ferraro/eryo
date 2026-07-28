package models

type Transaction struct {

	Hash string `json:"hash"`

	From string `json:"from"`

	To string `json:"to"`

	Value float64 `json:"value"`

	Status string `json:"status"`

}
