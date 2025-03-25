package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Delete godoc
// @Summary Delete a song by its ID
// @Description Delete a song from the database using its unique ID.
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} Response "Song was deleted successfully"
// @Failure 400 {object} Response "Invalid song ID"
// @Failure 404 {object} Response "Song with this ID isn't present"
// @Failure 500 {object} Response "Failed to delete song"
// @Router /api/songs/{id} [delete]
func (s *SongHandler) Delete(ctx echo.Context) error {
	songID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.logger.Warnw("invalid song id", "error", err, "input", ctx.Param("id"))
		return ctx.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "invalid song id",
		})
	}

	if exist := s.songUseCase.Exist(ctx.Request().Context(), songID); !exist {
		s.logger.Warnw("song not found", "song_id", songID)
		return ctx.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "song with this id isn't present",
		})
	}

	if err = s.songUseCase.DeleteSong(ctx.Request().Context(), songID); err != nil {
		s.logger.Errorw("failed to delete song", "song_id", songID, "error", err)
		return ctx.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "failed to delete song",
		})
	}

	s.logger.Infow("song deleted successfully", "song_id", songID)
	return ctx.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "song was deleted successfully",
	})
}
