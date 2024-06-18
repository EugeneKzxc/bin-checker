package models

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type BinBank struct {
	Bin  int    `db:"bin"`
	Bank string `db:"bank"`
}
