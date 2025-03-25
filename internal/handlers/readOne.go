package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Get godoc
// @Summary Get song text by song ID
// @Description Get the text of a song by its ID. If the song is not found, returns a 404 error.
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} SongTextResponse "Song text retrieved successfully"
// @Failure 400 {object} Response "Invalid song ID"
// @Failure 404 {object} Response "Song not found"
// @Failure 500 {object} Response "Internal server error"
// @Router /api/songs/{id} [get]
func (s *SongHandler) Get(ctx echo.Context) error {
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

	logger.Debug("Fetching song text", zap.Int("songID", songID))
	text, err := s.songUseCase.GetSongText(ctx.Request().Context(), songID)
	if err != nil {
		logger.Error("Failed to fetch song text", zap.Int("songID", songID), zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "failed to fetch song",
		})
	}

	logger.Info("Successfully retrieved song text", zap.Int("songID", songID))
	resp := SongTextResponse{
		TextParts: text,
	}
	return ctx.JSON(http.StatusOK, resp)
}
