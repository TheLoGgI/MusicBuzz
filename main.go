package main

import (
	"log"
	"net/http"

	"lasseaakjaer.com/musicbuzz/api"
)

// THIS FILE IS ONLY FOR DEVELOPMENT
// THE API WILL BE HOSTED ON VERCEL THAT ONLY ALLOWS FOR SPECIFIC ROUTE FORMAT.

func main() {
	// godotenv.Load()
	route := Route{
		router: &http.ServeMux{},
	}

	// Hosting on vercel means that every endpoint is run serverless, from the api folder. /api/something
	route.get("/api/login", api.Login)
	route.get("/api/me", api.Me)
	route.get("/api/callback", api.CompleteAuth)
	route.get("/api/autenticated", api.Authenticated)

	route.get("/", api.App)

	err := http.ListenAndServe(":8080", route.router)
	if err != nil {
		log.Fatal(err)
	}
}

type Route struct {
	router *http.ServeMux
	res    http.ResponseWriter
	req    http.Request
}

func (route Route) get(pattern string, handler http.HandlerFunc) {
	route.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		route.res = w
		route.req = *r
		route.methodeAllowed(http.MethodGet)

		handler(w, r)
	})
}

func (route Route) methodeAllowed(methode string) {
	if route.req.Method != methode {
		http.Error(route.res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// const redirectURI = "http://localhost:8080/api/callback"

// var (
// 	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
// 	ch    = make(chan *spotify.Client)
// 	state = "abc123"
// )

// func main() {
// 	// first start an HTTP server
// 	http.HandleFunc("/api/callback", completeAuth)
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Got request for:", r.URL.String())
// 	})
// 	go func() {
// 		err := http.ListenAndServe(":8080", nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	url := auth.AuthURL(state)
// 	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

// 	// wait for auth to complete
// 	client := <-ch

// 	// use the client to make calls that require authorization
// 	user, err := client.CurrentUser(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("You are logged in as:", user.ID)
// }

// func completeAuth(w http.ResponseWriter, r *http.Request) {
// 	tok, err := auth.Token(r.Context(), state, r)
// 	if err != nil {
// 		http.Error(w, "Couldn't get token", http.StatusForbidden)
// 		log.Fatal(err)
// 	}
// 	if st := r.FormValue("state"); st != state {
// 		http.NotFound(w, r)
// 		log.Fatalf("State mismatch: %s != %s\n", st, state)
// 	}

// 	// use the token to get an authenticated client
// 	client := spotify.New(auth.Client(r.Context(), tok))
// 	fmt.Fprintf(w, "Login Completed!")
// 	ch <- client
// }
