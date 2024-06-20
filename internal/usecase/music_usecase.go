package usecase

import (
	"context"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/spotifyclient"
	"github.com/zmb3/spotify/v2"
)

type MusicUsecase interface {
	SearchTrack(ctx context.Context, keyword string) (*SearchTrackOutput, error)
}

type musicUsecase struct {
	spotifyClient spotifyclient.SpotifyClient
}

func NewMusicUsecase(ctx context.Context, spotifyClient spotifyclient.SpotifyClient) MusicUsecase {
	return &musicUsecase{
		spotifyClient: spotifyClient,
	}
}

func (u *musicUsecase) SearchTrack(ctx context.Context, keyword string) (*SearchTrackOutput, error) {
	searchResult, err := u.spotifyClient.Search(ctx, keyword, spotify.SearchTypeTrack)
	if err != nil {
		return nil, ErrSearchingSpotify
	}

	searchOutput := new(SearchTrackOutput)
	tracks := make([]Track, len(searchResult.Tracks.Tracks))
	for i, t := range searchResult.Tracks.Tracks {
		tracks[i] = Track{ID: string(t.ID), Name: t.Name}
		artists := []Artist{}
		for _, a := range t.Artists {
			artists = append(artists, Artist{ID: string(a.ID), Name: a.Name})
		}
		tracks[i].Artists = artists
	}

	searchOutput.Tracks = tracks
	return searchOutput, nil
}
