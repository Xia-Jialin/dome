package dome01

//Chapter 章节
type Chapter struct {
	ID          int    `json:"id,omitempty"`
	ChapterName string `json:"chapter_name,omitempty"`
	ChapterURL  string `json:"chapter_url,omitempty"`
}
