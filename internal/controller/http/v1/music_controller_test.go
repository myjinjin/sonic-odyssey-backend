package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMusicController_SearchTrack(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		defer func() { mockMusicUsecase.Mock.ExpectedCalls = nil }()

		keyword := "One"
		limit := 10
		offset := 0
		mockOutput := &usecase.SearchTrackOutput{
			Tracks: []usecase.Track{
				{ID: "2up3OPMp9Tb4dAKM2erWXQ", Name: "One", Artists: []usecase.Artist{{ID: "DpdlalAks", Name: "Aimee mann"}}},
			},
		}
		mockMusicUsecase.On("SearchTrack", mock.Anything, keyword, &limit, &offset).Return(mockOutput, nil)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/music/tracks?keyword=One&limit=10&offset=0", nil)

		userID := uint(1)
		token, _, _ := testUserJwtAuth.TokenGenerator(&auth.UserPayload{UserID: userID})
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		var res SearchTrackResponse
		err := json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockOutput.Tracks[0].ID, res.Tracks[0].ID)
		assert.Equal(t, mockOutput.Tracks[0].Name, res.Tracks[0].Name)
		assert.Equal(t, mockOutput.Tracks[0].Artists[0].ID, res.Tracks[0].Artists[0].ID)
		assert.Equal(t, mockOutput.Tracks[0].Artists[0].Name, res.Tracks[0].Artists[0].Name)
		mockMusicUsecase.AssertExpectations(t)
	})

	t.Run("InvalidRequestBody", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/music/tracks?limit=abc", nil)

		userID := uint(1)
		token, _, _ := testUserJwtAuth.TokenGenerator(&auth.UserPayload{UserID: userID})
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockMusicUsecase.AssertNotCalled(t, "SearchTrack")
	})

	t.Run("Unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/music/tracks?keyword=One&limit=10&offset=0", nil)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockMusicUsecase.AssertNotCalled(t, "SearchTrack")
	})

	t.Run("SearchingSpotifyError", func(t *testing.T) {
		defer func() { mockMusicUsecase.Mock.ExpectedCalls = nil }()
		keyword := "One"
		limit := 10
		offset := 0
		mockMusicUsecase.On("SearchTrack", mock.Anything, keyword, &limit, &offset).Return(nil, usecase.ErrSearchingSpotify)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/music/tracks?keyword=One&limit=10&offset=0", nil)

		userID := uint(1)
		token, _, _ := testUserJwtAuth.TokenGenerator(&auth.UserPayload{UserID: userID})
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockMusicUsecase.AssertExpectations(t)
	})
}
