package api

import (
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Config struct {
	SpotifyID     string `env:"SPOTIFY_ID"`
	SpotifySecret string `env:"SPOTIFY_SECRET"`
	RedirectUri   string `env:"REDIRECT_URI" envDefault:"http://localhost:8080/api/callback"`
}

type GlobalConfig struct {
	auth *spotifyauth.Authenticator
	// ch    chan *spotify.Client
	state string
}

var authConfig = GlobalConfig{
	auth: spotifyauth.New(spotifyauth.WithRedirectURL(os.Getenv("REDIRECT_URI")), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate)),
	// ch:    make(chan *spotify.Client),
	state: "34fFs29kd09",
}

// var Config Config
// if err := envconfig.Process(ctx, &Config); err != nil {
// 	panic(err) // TODO: handle error
// }

// var ENV = Config{
// SpotifyID:     os.Getenv("SPOTIFY_ID"),
// SpotifySecret: os.Getenv("SPOTIFY_SECRET"),
// }
