package models

type Anime struct {
	MalID          int           `json:"mal_id"`
	URL            string        `json:"url"`
	Images         []Images      `json:"images"`
	Trailer        []Trailer     `json:"trailer"`
	Title          string        `json:"title"`
	TitleEnglish   string        `json:"title_english"`
	TitleJapanese  string        `json:"title_japanese"`
	TitleSynonyms  []string      `json:"title_synonyms"`
	Type           string        `json:"type"`
	Source         string        `json:"source"`
	Episodes       int           `json:"episodes"`
	Status         string        `json:"status"`
	Airing         bool          `json:"airing"`
	Aired          Aired         `json:"aired"`
	Duration       string        `json:"duration"`
	Rating         string        `json:"rating"`
	Score          float64       `json:"score"`
	ScoredBy       int           `json:"scored_by"`
	Rank           int           `json:"rank"`
	Popularity     int           `json:"popularity"`
	Members        int           `json:"members"`
	Favorites      int           `json:"favorites"`
	Synopsis       string        `json:"synopsis"`
	Background     string        `json:"background"`
	Season         string        `json:"season"`
	Year           int           `json:"year"`
	Broadcast      []Broadcast   `json:"broadcast"`
	Producers      []Producers   `json:"producers"`
	Licensors      []Licensors   `json:"licensors"`
	Studios        []Studios     `json:"studios"`
	Genres         []Genres      `json:"genres"`
	ExplicitGenres []interface{} `json:"explicit_genres"`
	Themes         []Themes      `json:"themes"`
	Demographics   []interface{} `json:"demographics"`
}

type Jpg struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}
type Webp struct {
	ImageURL      string `json:"image_url"`
	SmallImageURL string `json:"small_image_url"`
	LargeImageURL string `json:"large_image_url"`
}
type Images struct {
	Jpg  Jpg  `json:"jpg"`
	Webp Webp `json:"webp"`
}
type TrailerImages struct {
	ImageURL        string `json:"image_url"`
	SmallImageURL   string `json:"small_image_url"`
	MediumImageURL  string `json:"medium_image_url"`
	LargeImageURL   string `json:"large_image_url"`
	MaximumImageURL string `json:"maximum_image_url"`
}
type Trailer struct {
	YoutubeID     string        `json:"youtube_id"`
	URL           string        `json:"url"`
	EmbedURL      string        `json:"embed_url"`
	TrailerImages TrailerImages `json:"trailer_images"`
}
type From struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type To struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type Prop struct {
	From From `json:"from"`
	To   To   `json:"to"`
}
type Aired struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Prop   Prop   `json:"prop"`
	String string `json:"string"`
}
type Broadcast struct {
	Day      string `json:"day"`
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
	String   string `json:"string"`
}
type Producers struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
type Licensors struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
type Studios struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
type Genres struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
type Themes struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
