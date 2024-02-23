package model

import "time"

type Speedrun struct {
	PlayerName  string    `json:"player_name"`
	Game        string    `json:"game"`
	Category    string    `json:"category"`
	Time        string    `json:"time"`
	SubmittedAt time.Time `json:"submitted_at"`
}
