package spotifyclient

import (
	"context"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyClient interface {
	Search(ctx context.Context, query string, t spotify.SearchType) (*spotify.SearchResult, error)
}

type spotifyClient struct {
	client *spotify.Client
}

func New(ctx context.Context, clientID, clientSecret string) (SpotifyClient, error) {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	return &spotifyClient{
		client: client,
	}, nil
}

func (c *spotifyClient) Search(ctx context.Context, query string, t spotify.SearchType) (*spotify.SearchResult, error) {
	result, err := c.client.Search(ctx, query, t)
	if err != nil {
		return nil, err
	}
	return result, nil
}
