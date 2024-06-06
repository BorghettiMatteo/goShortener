package models

type Poem struct {
	Title     string   `json:"title" bosn:"title"`
	Lines     []string `json:"lines" bson:"lines"`
	Author    string   `json:"author" bson:"author"`
	Linecount string   `json:"linecount" bson:"linecount"`
}
