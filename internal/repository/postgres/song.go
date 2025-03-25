package postgres

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"song-lib/internal/models"
	"strings"
)

type SongRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewSongRepo(db *sqlx.DB, logger *zap.SugaredLogger) *SongRepo {
	return &SongRepo{db: db, logger: logger}
}

func (s *SongRepo) Exist(ctx context.Context, songID int) bool {
	query, args, err := sq.Select("COUNT(*) > 0").
		From("songs").
		Where(sq.Eq{"id": songID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for Exist", "error", err)
		return false
	}

	var exists bool
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		s.logger.Errorw("DB error in Exist", "songID", songID, "error", err)
		return false
	}

	s.logger.Debugw("Exist check", "songID", songID, "exists", exists)
	return exists
}

func (s *SongRepo) GetSongs(ctx context.Context, filter models.SongFilter) ([]models.Song, error) {
	query := sq.Select("id", "artist", "title", "release_date", "text", "source_link").
		From("songs")

	if filter.Artist != "" {
		query = query.Where(sq.Like{"artist": "%" + filter.Artist + "%"})
	}
	if filter.Title != "" {
		query = query.Where(sq.Like{"title": "%" + filter.Title + "%"})
	}
	if filter.ReleaseDate != "" {
		query = query.Where(sq.Eq{"release_date": filter.ReleaseDate})
	}
	if filter.Text != "" {
		query = query.Where(sq.Like{"text": "%" + filter.Text + "%"})
	}
	if filter.SourceLink != "" {
		query = query.Where(sq.Eq{"source_link": filter.SourceLink})
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	sqlQuery, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for GetSongs", "error", err)
		return nil, err
	}

	s.logger.Debugw("Executing GetSongs query", "query", sqlQuery, "args", args)

	rows, err := s.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		s.logger.Errorw("Failed to execute GetSongs query", "error", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			s.logger.Warnw("Failed to close rows", "error", err)
		}
	}(rows)

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Artist, &song.Title, &song.ReleaseDate, &song.Text, &song.SourceLink); err != nil {
			s.logger.Errorw("Failed to scan row in GetSongs", "error", err)
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		s.logger.Errorw("Rows iteration error in GetSongs", "error", err)
		return nil, err
	}

	s.logger.Infow("Successfully retrieved songs", "count", len(songs))
	return songs, nil
}

func (s *SongRepo) GetSongText(ctx context.Context, songID int) ([]string, error) {
	query, args, err := sq.Select("text").
		From("songs").
		Where(sq.Eq{"id": songID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for GetSongText", "error", err)
		return nil, err
	}

	s.logger.Debugw("Executing GetSongText query", "query", query, "args", args)

	var text string
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&text)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warnw("Song text not found", "songID", songID)
			return nil, nil
		}
		s.logger.Errorw("Failed to fetch song text", "songID", songID, "error", err)
		return nil, err
	}

	textParts := strings.Split(text, "\n\n")
	s.logger.Infow("Successfully retrieved song text", "songID", songID, "parts", len(textParts))
	return textParts, nil
}

func (s *SongRepo) CreateSong(ctx context.Context, song models.Song) error {
	query, args, err := sq.Insert("songs").
		Columns("artist", "title", "release_date", "text", "source_link").
		Values(song.Artist, song.Title, song.ReleaseDate, song.Text, song.SourceLink).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for CreateSong", "error", err)
		return err
	}

	s.logger.Infow("Executing CreateSong query", "query", query, "args", args)
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Errorw("Failed to execute CreateSong query", "error", err)
		return err
	}

	s.logger.Infow("Song created successfully", "title", song.Title, "artist", song.Artist)
	return nil
}

func (s *SongRepo) ChangeSong(ctx context.Context, song models.Song) error {
	query, args, err := sq.Update("songs").
		Set("artist", song.Artist).
		Set("title", song.Title).
		Set("release_date", song.ReleaseDate).
		Set("text", song.Text).
		Set("source_link", song.SourceLink).
		Where(sq.Eq{"id": song.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for ChangeSong", "error", err)
		return err
	}

	s.logger.Infow("Executing ChangeSong query", "query", query, "args", args)
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Errorw("Failed to execute ChangeSong query", "error", err)
		return err
	}

	s.logger.Infow("Song updated successfully", "songID", song.ID)
	return nil
}

func (s *SongRepo) DeleteSong(ctx context.Context, songID int) error {
	query, args, err := sq.Delete("songs").
		Where(sq.Eq{"id": songID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		s.logger.Errorw("Failed to build SQL query for DeleteSong", "error", err)
		return err
	}

	s.logger.Infow("Executing DeleteSong query", "query", query, "args", args)
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Errorw("Failed to execute DeleteSong query", "error", err)
		return err
	}

	s.logger.Infow("Song deleted successfully", "songID", songID)
	return nil
}
