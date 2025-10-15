package dto

import "task0/models"

// ResponseDto represents the JSON structure returned by the /me endpoint.
type ResponseDto struct {
	Status    string      `json:"status"`
	User      models.User `json:"user"`
	Timestamp string      `json:"timestamp"`
	Fact      string      `json:"fact"`
}
