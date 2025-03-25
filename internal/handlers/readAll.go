package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"song-lib/internal/models"
	"strconv"
)

// GetSongs godoc
// @Summary Get all songs with filtering and pagination
// @Description Get a list of songs based on filter criteria like artist, title, release date, text, and source link with pagination (limit and offset).
// @Tags songs
// @Accept json
// @Produce json
// @Param artist query string false "Artist name"
// @Param title query string false "Song title"
// @Param release_date query string false "Release date"
// @Param text query string false "Text content"
// @Param source_link query string false "Source link"
// @Param limit query int false "Limit of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} SongResponse "List of songs"
// @Failure 400 {object} Response "Invalid query parameters"
// @Failure 500 {object} Response "Failed to fetch songs"
// @Router /api/songs/filter [get]
func (s *SongHandler) GetSongs(ctx echo.Context) error {
	logger := zap.L()
	logger.Debug("handling GetSongs request")

	filter := models.SongFilter{
		Artist:      ctx.QueryParam("artist"),
		Title:       ctx.QueryParam("title"),
		ReleaseDate: ctx.QueryParam("release_date"),
		Text:        ctx.QueryParam("text"),
		SourceLink:  ctx.QueryParam("source_link"),
	}

	if limitStr := ctx.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			logger.Warn("invalid limit value", zap.String("limit", limitStr), zap.Error(err))
			return ctx.JSON(http.StatusBadRequest, Response{
				Code:    400,
				Message: "invalid limit value",
			})
		}
		filter.Limit = uint64(limit)
	}

	if offsetStr := ctx.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			logger.Warn("invalid offset value", zap.String("offset", offsetStr), zap.Error(err))
			return ctx.JSON(http.StatusBadRequest, Response{
				Code:    400,
				Message: "invalid offset value",
			})
		}
		filter.Offset = uint64(offset)
	}

	logger.Debug("Fetching songs with filter", zap.Any("filter", filter))

	songs, err := s.songUseCase.GetSongs(ctx.Request().Context(), filter)
	if err != nil {
		logger.Error("failed to fetch songs", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "failed to fetch songs",
		})
	}

	var resp []SongResponse
	for _, song := range songs {
		resp = append(resp, SongResponse{
			ID:          song.ID,
			Artist:      song.Artist,
			Title:       song.Title,
			ReleaseDate: song.ReleaseDate,
			Text:        song.Text,
			SourceLink:  song.SourceLink,
		})
	}

	logger.Info("Successfully fetched songs", zap.Int("count", len(resp)))

	return ctx.JSON(http.StatusOK, resp)
}
