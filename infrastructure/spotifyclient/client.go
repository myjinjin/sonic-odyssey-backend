package spotifyclient

import (
	"net/http"

	"github.com/zmb3/spotify/v2"
)

type SpotifyClient interface {
	// TODO
}

type spotifyClient struct {
	client *spotify.Client
}

func New(httpClient *http.Client) SpotifyClient {
	client := spotify.New(httpClient)
	return &spotifyClient{
		client: client,
	}
}
