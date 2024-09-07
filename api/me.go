package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

func Me(w http.ResponseWriter, r *http.Request) {

	fmt.Println("COOKIE: ", r.Cookies())
	authToken, _ := r.Cookie("sb-127-auth-token")
	fmt.Println("COOKIE Token: ", authToken)
	// r.Header.Set("Content-Type", "application/json")
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatalf("couldn't get private user: %v", err)
	}
	fmt.Println("User:", user)

	userPlaylists, err := client.CurrentUsersPlaylists(ctx, spotify.Limit(10))
	if err != nil {
		log.Fatalf("couldn't get private user: %v", err)
	}
	fmt.Println("userPlaylists:", userPlaylists)

	w.Write([]byte("Hello from the API"))
}
