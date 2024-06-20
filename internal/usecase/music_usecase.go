package usecase

import (
	"context"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/spotifyclient"
	"github.com/zmb3/spotify/v2"
)

type MusicUsecase interface {
	SearchTrack(ctx context.Context, keyword string, limit, offset *int) (*SearchTrackOutput, error)
}

type musicUsecase struct {
	spotifyClient spotifyclient.SpotifyClient
}

func NewMusicUsecase(ctx context.Context, spotifyClient spotifyclient.SpotifyClient) MusicUsecase {
	return &musicUsecase{
		spotifyClient: spotifyClient,
	}
}

func (u *musicUsecase) SearchTrack(ctx context.Context, keyword string, limit, offset *int) (*SearchTrackOutput, error) {
	opts := []spotify.RequestOption{}
	if limit != nil && offset != nil {
		opts = append(opts, spotify.Limit(*limit))
		opts = append(opts, spotify.Offset(*offset))
	}
	searchResult, err := u.spotifyClient.Search(ctx, keyword, spotify.SearchTypeTrack, opts...)
	if err != nil {
		return nil, ErrSearchingSpotify
	}

	searchOutput := new(SearchTrackOutput)
	tracks := make([]Track, len(searchResult.Tracks.Tracks))
	total := searchResult.Tracks.Total
	for i, t := range searchResult.Tracks.Tracks {
		tracks[i] = Track{ID: string(t.ID), Name: t.Name}
		artists := []Artist{}
		for _, a := range t.Artists {
			artists = append(artists, Artist{ID: string(a.ID), Name: a.Name})
		}
		tracks[i].Artists = artists
	}

	searchOutput.Tracks = tracks
	searchOutput.Total = int(total)
	return searchOutput, nil
}
