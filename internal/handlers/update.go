package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"song-lib/internal/models"
	"strconv"
)

// Update godoc
// @Summary Update a song by ID
// @Description Update an existing song by its ID. If the song is not found, returns a 404 error.
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param body body UpdateRequest true "Song data to update"
// @Success 200 {object} Response "Song updated successfully"
// @Failure 400 {object} Response "Invalid song ID or request body"
// @Failure 404 {object} Response "Song not found"
// @Failure 500 {object} Response "Internal server error"
// @Router /api/songs/{id} [put]
func (s *SongHandler) Update(ctx echo.Context) error {
	logger := zap.L()

	songID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Warn("Invalid song ID", zap.String("param", ctx.Param("id")), zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "invalid song id",
		})
	}

	logger.Debug("Checking if song exists", zap.Int("songID", songID))
	if exist := s.songUseCase.Exist(ctx.Request().Context(), songID); !exist {
		logger.Warn("Song not found", zap.Int("songID", songID))
		return ctx.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "song with this id isn't present",
		})
	}

	var req UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Warn("Invalid request body", zap.Int("songID", songID), zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "invalid request body",
		})
	}

	song := models.Song{
		ID:          songID,
		Artist:      req.Artist,
		Title:       req.Title,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		SourceLink:  req.SourceLink,
	}

	logger.Info("Updating song", zap.Int("songID", songID), zap.Any("updateData", req))
	if err = s.songUseCase.ChangeSong(ctx.Request().Context(), song); err != nil {
		logger.Error("Failed to update song", zap.Int("songID", songID), zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "failed to update song",
		})
	}

	logger.Info("Song updated successfully", zap.Int("songID", songID))
	return ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "song was updated successfully",
	})
}
