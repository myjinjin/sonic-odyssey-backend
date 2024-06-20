package usecase

import (
	"context"
	"testing"

	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zmb3/spotify/v2"
)

func TestMusicUsecase_SearchTrack_Success(t *testing.T) {
	// Setup
	spotifyClient := &mocks.SpotifyClient{}

	ctx := context.Background()
	musicUsecase := NewMusicUsecase(ctx, spotifyClient)

	keyword := "One"
	limit := 10
	offset := 0
	searchType := spotify.SearchTypeTrack

	// Expectations
	spotifyClient.On("Search", ctx, keyword, spotify.SearchType(searchType), mock.AnythingOfType("spotify.RequestOption"), mock.AnythingOfType("spotify.RequestOption")).Return(&spotify.SearchResult{
		Tracks: &spotify.FullTrackPage{
			Tracks: []spotify.FullTrack{
				{
					SimpleTrack: spotify.SimpleTrack{
						ID:   "2up3OPMp9Tb4dAKM2erWXQ",
						Name: "One",
						Artists: []spotify.SimpleArtist{
							{
								ID:   "DpdlalAks",
								Name: "Aimee mann",
							},
						},
					},
				},
			},
		},
	}, nil)

	// Execute
	output, err := musicUsecase.SearchTrack(ctx, keyword, &limit, &offset)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.Tracks)

	// Verify
	spotifyClient.AssertExpectations(t)
}

func TestMusicUsecase_SearchTrack_SearchingSpotifyError(t *testing.T) {
	// Setup
	spotifyClient := &mocks.SpotifyClient{}

	ctx := context.Background()
	musicUsecase := NewMusicUsecase(ctx, spotifyClient)

	keyword := "One"
	limit := 10
	offset := 0
	searchType := spotify.SearchTypeTrack

	// Expectations
	spotifyClient.On("Search", ctx, keyword, spotify.SearchType(searchType), mock.AnythingOfType("spotify.RequestOption"), mock.AnythingOfType("spotify.RequestOption")).Return(nil, ErrSearchingSpotify)

	// Execute
	output, err := musicUsecase.SearchTrack(ctx, keyword, &limit, &offset)

	// Assert
	assert.ErrorIs(t, err, ErrSearchingSpotify)
	assert.Nil(t, output)

	// Verify
	spotifyClient.AssertExpectations(t)
}

func TestMusicUsecase_SearchTrack_EmptyResult(t *testing.T) {
	// Setup
	spotifyClient := &mocks.SpotifyClient{}

	ctx := context.Background()
	musicUsecase := NewMusicUsecase(ctx, spotifyClient)

	keyword := "NonExistentTrack"
	limit := 10
	offset := 0
	searchType := spotify.SearchTypeTrack

	// Expectations
	spotifyClient.On("Search", ctx, keyword, spotify.SearchType(searchType), mock.AnythingOfType("spotify.RequestOption"), mock.AnythingOfType("spotify.RequestOption")).Return(&spotify.SearchResult{
		Tracks: &spotify.FullTrackPage{
			Tracks: []spotify.FullTrack{},
		},
	}, nil)

	// Execute
	output, err := musicUsecase.SearchTrack(ctx, keyword, &limit, &offset)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Empty(t, output.Tracks)

	// Verify
	spotifyClient.AssertExpectations(t)
}
