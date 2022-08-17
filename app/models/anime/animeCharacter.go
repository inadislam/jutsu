package models

type CharJpg struct {
	ImageURL string `json:"image_url"`
}
type CharWebp struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
}
type CharImages struct {
	Jpg  CharJpg  `json:"jpg"`
	Webp CharWebp `json:"webp"`
}
type Character struct {
	MalID  int    `json:"mal_id"`
	URL    string `json:"url"`
	Images CharImages `json:"images"`
	Name   string `json:"name"`
}
type VAImages struct {
	Jpg CharJpg `json:"jpg"`
}
type Person struct {
	MalID  int    `json:"mal_id"`
	URL    string `json:"url"`
	Images VAImages `json:"images"`
	Name   string `json:"name"`
}
type VoiceActors struct {
	Person   Person `json:"person"`
	Language string `json:"language"`
}
type Characters struct {
	Character   Character     `json:"character"`
	Role        string        `json:"role"`
	VoiceActors []VoiceActors `json:"voice_actors"`
}