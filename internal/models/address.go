package models

type Address struct {

	ID string `json:"id"`

	Balance float64 `json:"balance"`

	TxCount int `json:"tx_count"`

	Network string `json:"network"`

}
