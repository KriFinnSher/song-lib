package handlers

import (
	"go.uber.org/zap"
	"song-lib/internal/usecase"
)

type SongHandler struct {
	songUseCase *usecase.SongUseCase
	logger      *zap.SugaredLogger
}

func NewSongHandler(songUseCase *usecase.SongUseCase, logger *zap.SugaredLogger) *SongHandler {
	return &SongHandler{songUseCase: songUseCase, logger: logger}
}

type Request struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpdateRequest struct {
	Artist      string `json:"artist"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Text        string ` json:"text"`
	SourceLink  string `json:"source_link"`
}

type SongTextResponse struct {
	TextParts []string `json:"text_parts"`
}

type SongResponse struct {
	ID          int    `json:"id"`
	Artist      string `json:"artist"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Text        string ` json:"text"`
	SourceLink  string `json:"source_link"`
}
