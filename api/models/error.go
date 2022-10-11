package models

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
