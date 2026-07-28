package models

type Block struct {

	Number int `json:"number"`

	Hash string `json:"hash"`

	TxCount int `json:"tx_count"`

	Timestamp string `json:"timestamp"`

}
