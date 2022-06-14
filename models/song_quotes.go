package models

type SongQuote struct {
	ID         int64  `json:"id" db:"id"`
	QuoteText  string `json:"quoteText" db:"quote_text"`
	SongTitle  string `json:"songTitle" db:"song_title"`
	AlbumTitle string `json:"albumTitle" db:"album_title"`
	AlbumYear  int    `json:"albumYear" db:"album_year"`
	BandName   string `json:"bandName" db:"band_name"`
}
