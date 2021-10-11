package model

type Session struct {
	CurrentSong Song `json:"current_song"`
	Position int `json:"position"`
}