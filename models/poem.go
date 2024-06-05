package models

type Poem struct {
	Title     string   `json:"title"`
	Lines     []string `json:"lines"`
	Author    string   `json:"author"`
	Linecount string   `json:"linecount"`
}
