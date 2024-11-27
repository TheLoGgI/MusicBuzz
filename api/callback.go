package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify/v2"
	"lasseaakjaer.com/musicbuzz/config"
)

// TODO: refactor to use the global config
// var authChannel = make(chan *spotify.Client)

func CompleteAuth(w http.ResponseWriter, r *http.Request) {

	config := config.AuthConfig
	fmt.Println(config)

	tok, err := config.Auth.Token(r.Context(), config.State, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != config.State {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, config.State)
	}

	// use the token to get an authenticated client
	client := spotify.New(config.Auth.Client(r.Context(), tok))
	user, err := client.CurrentUser(r.Context())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Login Completed!", user.ID, user.DisplayName)
	// authChannel <- client

	// w.Write([]byte("Login Completed!"))
	redirectUri := fmt.Sprintf("/%s?state=autenticated", user.ID)
	http.Redirect(w, r, redirectUri, http.StatusSeeOther)
	// http.Redirect(w, r, "/api/login?state=autenticated", http.StatusSeeOther)

}
