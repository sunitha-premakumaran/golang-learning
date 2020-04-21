package cyoa

// Story ...
type Story map[string]Chapter

// Chapter ...
type Chapter struct {
	Title   string          `json:"title"`
	Story   []string        `json:"story"`
	Options []ChapterOption `json:"options"`
}

// ChapterOption ..
type ChapterOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc_name"`
}
