package adapter

// import (
// 	"net/http"

// 	"github.com/zmb3/spotify/v2"
// 	spotifyauth "github.com/zmb3/spotify/v2/auth"
// )

// func main() {
// 	var redirectURL = "http://localhost:8080/callback"
// 	// the redirect URL must be an exact match of a URL you've registered for your application
// 	// scopes determine which permissions the user is prompted to authorize
// 	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))

// 	// get the user to this URL - how you do that is up to you
// 	// you should specify a unique state string to identify the session
// 	url := auth.AuthURL(state)
// }

// // the user will eventually be redirected back to your redirect URL
// // typically you'll have a handler set up like the following:
// func redirectHandler(w http.ResponseWriter, r *http.Request) {
// 	// use the same state string here that you used to generate the URL
// 	token, err := auth.Token(r.Context(), state, r)
// 	if err != nil {
// 		http.Error(w, "Couldn't get token", http.StatusNotFound)
// 		return
// 	}
// 	// create a client using the specified token
// 	client := spotify.New(auth.Client(r.Context(), token))

// 	// the client can now be used to make authenticated requests
// }
