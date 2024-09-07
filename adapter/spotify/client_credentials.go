// This example demonstrates how to authenticate with Spotify using the
// client credentials flow.  Note that this flow does not include authorization
// and can't be used to access a user's private data.
//
// Make sure you set the SPOTIFY_ID and SPOTIFY_SECRET environment variables
// prior to running this example.
package adapter

import (
	"context"
	"fmt"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyAdapter struct {
}

func (adapter SpotifyAdapter) Credentials() {

	if os.Getenv("SPOTIFY_ID") == "" || os.Getenv("SPOTIFY_SECRET") == "" {
		log.Fatal("please set SPOTIFY_ID and SPOTIFY_SECRET")
	}

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
	// msg, page, err := client.FeaturedPlaylists(ctx)
	// if err != nil {
	// 	log.Fatalf("couldn't get features playlists: %v", err)
	// }

	playlistsItems, err := client.GetPlaylistItems(ctx, "37i9dQZF1DX60OAKjsWlA2")
	if err != nil {
		log.Fatalf("couldn't get playlist items: %v", err)
	}
	fmt.Println(playlistsItems.Items)
	// client.GetPlaylist(ctx, "37i9dQZF1DXcBWIGoYBM5M")

	// user, err := client.CurrentUser(ctx)
	// if err != nil {
	// 	log.Fatalf("couldn't get private user: %v", err)
	// }
	// fmt.Println("User:", &user.DisplayName)

	fmt.Println("It's Hits Denmark")
	for _, track := range playlistsItems.Items {
		fmt.Println(track.AddedBy.DisplayName, track.Track.Track.Name, track.Track.Track.ID, track.Track.Track.Artists[0].Name, track.Track.Track.Album.Name, track.Track.Track.Album.Images[0].URL)
	}
	// for _, playlist := range page.Playlists {
	// 	fmt.Println("  ", playlist.Name, playlist.ID)

	// 	client.GetPlaylistItems(ctx, "37i9dQZF1DX60OAKjsWlA2")

	// }
}
