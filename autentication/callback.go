package api

import (
	"net/http"

	"github.com/zmb3/spotify/v2"
)

// TODO: refactor to use the global config
var authChannel = make(chan *spotify.Client)

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	// tok, err := authConfig.auth.Token(r.Context(), authConfig.state, r)
	// if err != nil {
	// 	http.Error(w, "Couldn't get token", http.StatusForbidden)
	// 	log.Fatal(err)
	// }
	// if st := r.FormValue("state"); st != authConfig.state {
	// 	http.NotFound(w, r)
	// 	log.Fatalf("State mismatch: %s != %s\n", st, authConfig.state)
	// }

	// // use the token to get an authenticated client
	// client := spotify.New(authConfig.auth.Client(r.Context(), tok))
	// fmt.Fprintf(w, "Login Completed!")
	// fmt.Printf("Login Completed! %v", client)
	// authChannel <- client
	// http.Redirect(w, r, "/api/login?state=autenticated", http.StatusSeeOther)

}
