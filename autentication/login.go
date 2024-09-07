package api

import "net/http"

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

func Login(w http.ResponseWriter, r *http.Request) {

	// 	var autenticationState = r.URL.Query().Get("state")
	// 	fmt.Println("autenticationState:", autenticationState)
	// 	if autenticationState == "" {
	// 		url := authConfig.auth.AuthURL(authConfig.state)
	// 		// fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// 		w.Header().Set("Content-Type", "text/html")
	// 		w.Write([]byte(fmt.Sprintf("<a href='%s'>Log in to Spotify</a>", url)))
	// 		return
	// 	}

	// 	if autenticationState == "autenticated" {
	// 		client := <-authChannel

	// 		fmt.Printf("Login Completed! %v", client)

	// 		// use the client to make calls that require authorization
	// 		user, err := client.CurrentUser(context.Background())
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		// fmt.Println("\nYou are logged in as:", user.ID)

	// 		w.Write([]byte("<p>You are logged in as:" + user.ID + "</p>"))
	// }

	// autenticatedClient := spotify.New(authConfig.auth.Client(r.Context(), accessToken))
	// fmt.Println("Login Completed!", client)
	// 	// fmt.Printf("Login Completed! %v", client)

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
