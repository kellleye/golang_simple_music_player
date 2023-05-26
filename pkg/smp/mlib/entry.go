package library

type MusicEntry struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Source string `json:"source"`
	Type   string `json:"type"`
}
