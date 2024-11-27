package config

import (
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

// type envs struct {
// 	SpotifyID     string `env:"SPOTIFY_ID"`
// 	SpotifySecret string `env:"SPOTIFY_SECRET"`
// 	RedirectUri   string `env:"REDIRECT_URI" envDefault:"http://localhost:8080/api/callback"`
// }

type spotifyAuthConfig struct {
	Auth    *spotifyauth.Authenticator
	Channel chan *spotify.Client
	State   string
}

var AuthConfig = spotifyAuthConfig{
	Auth:    spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("REDIRECT_URI")), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate)),
	Channel: make(chan *spotify.Client),
	State:   "34fFs29kd09",
}
