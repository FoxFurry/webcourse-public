package model

type Song struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Album string `json:"album"`
	Genre string `json:"genre"`
	Year int `json:"year"`
}
