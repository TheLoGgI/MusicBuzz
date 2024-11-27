package api

import (
	"fmt"
	"net/http"

	"lasseaakjaer.com/musicbuzz/config"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var autenticationState = r.URL.Query().Get("state")
	if autenticationState == "" {

	}

	// if autenticationState == "autenticated" {
	// 	client := <-authChannel

	// 	fmt.Printf("Login Completed! %v", client)

	// 	// use the client to make calls that require authorization
	// 	user, err := client.CurrentUser(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// fmt.Println("\nYou are logged in as:", user.ID)

	// 	w.Write([]byte("<p>You are logged in as:" + user.ID + "</p>"))
	// 	return
	// }

	auth := config.AuthConfig.Auth
	// token, err := auth.Token(r.Context(), config.AuthConfig.State, r)
	// if err != nil {
	// 	http.Error(w, "Couldn't get token", http.StatusForbidden)
	// 	log.Fatal(err)
	// }

	// fmt.Println(token)
	url := auth.AuthURL(config.AuthConfig.State)

	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf("<a href='%s'>Log in to Spotify</a>", url)))

	// autenticatedClient := spotify.New(config.AuthConfig.Auth.Client(r.Context(), accessToken))
	// fmt.Println("Login Completed!", autenticatedClient)
	// fmt.Println("Login Completed!", client)
	// fmt.Printf("Login Completed! %v", client)

	// use the client to make calls that require authorization
	// user, err := client.CurrentUser(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("\nYou are logged in as:", user.ID)

	// 	w.Write([]byte("<p>You are logged in as:" + user.ID + "</p>"))

	// }()

	// var url = fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=user-read-private%%20user-read-email&state=34fFs29kd09", spotifyID, "http://localhost:8080/api/me")
	// http.Redirect(w, r, url, http.StatusSeeOther)

}
