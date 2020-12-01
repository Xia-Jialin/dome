package dome01

import "testing"

func TestChapter_GetChapterContent(t *testing.T) {
	tests := []struct {
		name string
		c    Chapter
		want string
	}{
		// TODO: Add test cases.
		{name:"123",
		c:Chapter{ID:1, ChapterName:"", ChapterURL:"http://www.xbiquge.la/15/15409/13232685.html"},
		want:""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got,_ := tt.c.GetChapterContent(); got != tt.want {
				t.Errorf("Chapter.GetChapterContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
