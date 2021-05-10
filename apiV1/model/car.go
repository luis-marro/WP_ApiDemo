package model

type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Engine string `json:"engine"`
	Year   int    `json:"year"`
}
