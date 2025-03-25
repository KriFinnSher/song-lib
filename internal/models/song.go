package models

type Song struct {
	ID          int    `db:"id" json:"id"`
	Artist      string `db:"artist" json:"artist"`
	Title       string `db:"title" json:"title"`
	ReleaseDate string `db:"release_date" json:"release_date"`
	Text        string `db:"text" json:"text"`
	SourceLink  string `db:"source_link" json:"source_link"`
}

type SongFilter struct {
	Artist      string
	Title       string
	ReleaseDate string
	Text        string
	SourceLink  string
	Limit       uint64
	Offset      uint64
}
