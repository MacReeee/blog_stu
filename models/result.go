package models

type Result struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
	Code  int    `json:"code"`
}
