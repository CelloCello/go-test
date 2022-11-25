package net

type Response struct {
	Message string      `json:"msg"`
	Payload interface{} `json:"payload"`
	Error   string      `json:"error"`
}
