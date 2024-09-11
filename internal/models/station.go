package models

type Station struct {
	Name string   `json:"name"`
	Line []string `json:"line"`
	Zone []int    `json:"zone"`
	ID   string   `json:"id"`
}
