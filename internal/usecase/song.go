package usecase

import (
	"context"
	"go.uber.org/zap"
	"song-lib/internal/externalAPI"
	"song-lib/internal/models"
	"song-lib/internal/usecase/song"
)

type SongUseCase struct {
	Repo        song.Repository
	ExternalAPI externalAPI.Client
	logger      *zap.SugaredLogger
}

func NewSongInstance(repo song.Repository, externalAPI externalAPI.Client, logger *zap.SugaredLogger) *SongUseCase {
	return &SongUseCase{Repo: repo, ExternalAPI: externalAPI, logger: logger}
}

func (s *SongUseCase) Exist(ctx context.Context, songID int) bool {
	s.logger.Debugw("Checking if song exists", "songID", songID)
	exists := s.Repo.Exist(ctx, songID)
	s.logger.Debugw("Song existence check completed", "songID", songID, "exists", exists)
	return exists
}

func (s *SongUseCase) AddSong(ctx context.Context, group string, songTitle string) error {
	s.logger.Infow("Adding new song", "group", group, "songTitle", songTitle)

	externalData, err := s.ExternalAPI.GetSongDetails(ctx, group, songTitle)
	if err != nil {
		s.logger.Errorw("Failed to fetch song details from external API", "group", group, "songTitle", songTitle, "error", err)
		return err
	}

	songInstance := models.Song{
		Artist:      group,
		Title:       songTitle,
		ReleaseDate: externalData.ReleaseDate,
		Text:        externalData.Text,
		SourceLink:  externalData.Link,
	}

	err = s.Repo.CreateSong(ctx, songInstance)
	if err != nil {
		s.logger.Errorw("Failed to add song to the database", "song", songInstance, "error", err)
		return err
	}

	s.logger.Infow("Song added successfully", "song", songInstance)
	return nil
}

func (s *SongUseCase) ChangeSong(ctx context.Context, song models.Song) error {
	s.logger.Infow("Updating song", "songID", song.ID, "title", song.Title)

	err := s.Repo.ChangeSong(ctx, song)
	if err != nil {
		s.logger.Errorw("Failed to update song", "songID", song.ID, "title", song.Title, "error", err)
		return err
	}

	s.logger.Infow("Song updated successfully", "songID", song.ID, "title", song.Title)
	return nil
}

func (s *SongUseCase) DeleteSong(ctx context.Context, songID int) error {
	s.logger.Infow("Deleting song", "songID", songID)

	err := s.Repo.DeleteSong(ctx, songID)
	if err != nil {
		s.logger.Errorw("Failed to delete song", "songID", songID, "error", err)
		return err
	}

	s.logger.Infow("Song deleted successfully", "songID", songID)
	return nil
}

func (s *SongUseCase) GetSongs(ctx context.Context, filter models.SongFilter) ([]models.Song, error) {
	s.logger.Infow("Retrieving songs", "filter", filter)

	songs, err := s.Repo.GetSongs(ctx, filter)
	if err != nil {
		s.logger.Errorw("Failed to retrieve songs", "filter", filter, "error", err)
		return nil, err
	}

	s.logger.Infow("Successfully retrieved songs", "count", len(songs))
	return songs, nil
}

func (s *SongUseCase) GetSongText(ctx context.Context, songID int) ([]string, error) {
	s.logger.Infow("Retrieving song text", "songID", songID)

	text, err := s.Repo.GetSongText(ctx, songID)
	if err != nil {
		s.logger.Errorw("Failed to retrieve song text", "songID", songID, "error", err)
		return nil, err
	}

	s.logger.Infow("Successfully retrieved song text", "songID", songID, "parts", len(text))
	return text, nil
}
