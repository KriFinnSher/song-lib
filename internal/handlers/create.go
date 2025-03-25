package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Create godoc
// @Summary Create a new song
// @Description Create a new song by providing the group and song title.
// @Tags songs
// @Accept json
// @Produce json
// @Param song body Request true "Song data"
// @Success 200 {object} Response "Song was created successfully"
// @Failure 400 {object} Response "Invalid request body"
// @Failure 500 {object} Response "Failed to create song"
// @Router /api/songs [post]
func (s *SongHandler) Create(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		s.logger.Warnw("invalid request body", "error", err)
		return ctx.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "invalid request body",
		})
	}

	if err := s.songUseCase.AddSong(ctx.Request().Context(), req.Group, req.Song); err != nil {
		s.logger.Errorw("failed to create song", "error", err)
		return ctx.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "failed to create song",
		})
	}

	s.logger.Infow("song created successfully", "group", req.Group, "song", req.Song)
	return ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "song was created successfully",
	})
}
