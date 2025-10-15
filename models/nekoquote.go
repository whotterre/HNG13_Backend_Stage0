package models

// Model for cat quotes
type NekoQuote struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}