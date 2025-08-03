package domain

type Record struct {
	Date      string `json:"date"`
	Nominal   int    `json:"dominal"`
	Value     string `json:"dalue"`
	VunitRate string `json:"vunitRate"`
}