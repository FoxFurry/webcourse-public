package model

import "math/rand"

type Song struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Album string `json:"album"`
	Genre string `json:"genre"`
	Year int `json:"year"`
}

var playlist = []Song{
	{
		Title:  "Juicebox",
		Author: "The Strokes",
		Album:  "First Impressions of Earth",
		Genre:  "Rock",
		Year:   2005,
	},
	{
		Title:  "Sunborn",
		Author: "Muse",
		Album:  "Snowbiz",
		Genre:  "Alternative/Indie",
		Year:   1999,
	},
	{
		Title:  "Around the World",
		Author: "Daft Punk",
		Album:  "Homework",
		Genre:  "Techno",
		Year:   1997,
	},
	{
		Title:  "Rain in My Hearth",
		Author: "Frank Sinatra",
		Album:  "Cycles",
		Genre:  "Jazz",
		Year:   1968,
	},
	{
		Title:  "Summer Wine",
		Author: "Nancy Sinatra",
		Album:  "Nancy in London",
		Genre:  "Pop",
		Year:   1966,
	},
	{
		Title:  "Karma Police",
		Author: "Radiohead",
		Album:  "OK Computer",
		Genre:  "Alternative rock",
		Year:   2017,
	},
	{
		Title:  "Feel Good Inc",
		Author: "Gorillaz",
		Album:  "Demon Days",
		Genre:  "Rock",
		Year:   2005,
	},
	{
		Title:  "Enjoy the Silence",
		Author: "Depeche Mode",
		Album:  "Violator",
		Genre:  "Alternative/Indie",
		Year:   1990,
	},
}

func GetRandomSong() Song {
	randIdx := rand.Int()% len(playlist)
	return playlist[randIdx]
}
