package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

type MusicController interface {
	SearchTrack(c *gin.Context)
}

type musicController struct {
	musicUsecase usecase.MusicUsecase
	jwtAuth      *auth.JWTMiddleware
}

func NewMusicController(musicUsecase usecase.MusicUsecase, jwtAuth *auth.JWTMiddleware) MusicController {
	return &musicController{
		musicUsecase: musicUsecase,
		jwtAuth:      jwtAuth,
	}
}

// SearchTrack godoc
// @Summary      search music track
// @Description  JWT 인증 토큰 기반 내 유저 정보 조회
// @Tags         music, tracks
// @Accept       json
// @Produce      json
// @Param request query SearchTrackRequest true "SearchTrack Request"
// @Security     BearerAuth
// @Success      200  {object}  SearchTrackResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/music/tracks [get]
func (m *musicController) SearchTrack(c *gin.Context) {
	var req SearchTrackRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		HandleError(c, ErrInvalidRequestBody)
		return
	}

	output, err := m.musicUsecase.SearchTrack(c, req.Keyword, req.Limit, req.Offset)
	if err != nil {
		HandleError(c, err)
		return
	}

	tracks := make([]Track, len(output.Tracks))
	total := output.Total
	for i, t := range output.Tracks {
		track := Track{ID: t.ID, Name: t.Name}
		artists := make([]Artist, len(t.Artists))
		for j, a := range t.Artists {
			artists[j] = Artist{ID: a.ID, Name: a.Name}
		}
		track.Artists = artists
		tracks[i] = track
	}

	res := SearchTrackResponse{Tracks: tracks, Total: total}
	c.JSON(http.StatusOK, res)
}
