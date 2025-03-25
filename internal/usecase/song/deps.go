package song

import (
	"context"
	"song-lib/internal/models"
)

type Repository interface {
	Exist(ctx context.Context, songID int) bool
	GetSongs(ctx context.Context, filter models.SongFilter) ([]models.Song, error)
	GetSongText(ctx context.Context, songID int) ([]string, error)
	CreateSong(ctx context.Context, song models.Song) error
	ChangeSong(ctx context.Context, song models.Song) error
	DeleteSong(ctx context.Context, songID int) error
}
