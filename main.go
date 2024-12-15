package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/websocket"
	"lasseaakjaer.com/musicbuzz/api"
	ws "lasseaakjaer.com/musicbuzz/websocket"
)

// THIS FILE IS ONLY FOR DEVELOPMENT
// THE API WILL BE HOSTED ON VERCEL THAT ONLY ALLOWS FOR SPECIFIC ROUTE FORMAT.

func main() {
	route := Route{
		router: &http.ServeMux{},
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(route.router),
		// TLSConfig: &tls.Config{},
	}

	// Hosting on vercel means that every endpoint is run serverless, from the api folder. /api/something
	// route.get("/api/login", api.Login)
	// route.get("/api/me", api.Me)
	// route.get("/api/callback", api.CompleteAuth)

	// route.group("/api/v1/auth/", middlewareAuth(api.AuthMux()))

	fmt.Println("Server is running on port 8080")
	route.get("/", api.App)

	// http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
	// 	// Handle WebSocket connection
	// 	fmt.Println("New Connection", ws.RemoteAddr())
	// 	ws.
	// }))
	wsServer := ws.CreateServer()
	route.ws("/ws", wsServer.AddWebSocketHandler)
	route.get("/send-response", wsServer.SendGroupMessageHandler)
	route.post("/group", wsServer.GetGroupMembers)

	// http.ListenAndServe(":8080", route.router)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}

type Route struct {
	router *http.ServeMux
	res    http.ResponseWriter
	req    http.Request
}

func (route Route) ws(pattern string, handler websocket.Handler) {
	route.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		route.res = w
		route.req = *r
		route.methodeAllowed(http.MethodGet)

		handler.ServeHTTP(w, r)
	})

}

func (route Route) get(pattern string, handler http.HandlerFunc) {
	route.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		route.res = w
		route.req = *r
		route.methodeAllowed(http.MethodGet)

		handler(w, r)
	})
}

func (route Route) post(pattern string, handler http.HandlerFunc) {
	route.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		route.res = w
		route.req = *r
		route.methodeAllowed(http.MethodPost)

		handler(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Change '*' to a specific origin for better security
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := getJWTToken(r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token, _, err := jwt.NewParser().ParseUnverified(jwtToken, jwt.MapClaims{
			"role": "admin",
		})

		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// if !token.Valid {
		// 	fmt.Println("Token is not valid")
		// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		// 	return
		// }
		expirationDate, err := token.Claims.GetExpirationTime()
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		fmt.Println(expirationDate.Unix(), time.Now().Unix())
		if expirationDate.Unix() < time.Now().Unix() {
			fmt.Println("Token has expired")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// if token["role"] == "admin" {
		// 	fmt.Println("Token not autenticated")
		// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		// 	return
		// }

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims)
		}

		fmt.Println(token)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func getJWTToken(r *http.Request) (string, error) {
	var cookieString = r.Header["Cookie"][0]
	var cookiesStruct = mapCookieHeader(cookieString)

	decodedCookie, err := url.QueryUnescape(cookiesStruct["sb-127-auth-token"])
	if err != nil {
		return "", err
	}
	decodedCookie = strings.TrimPrefix(decodedCookie, `["`)
	decodedCookie = strings.Split(decodedCookie, "\",")[0]
	return decodedCookie, nil
}

func mapCookieHeader(cookieString string) map[string]string {
	var cookieMap = make(map[string]string)
	var cookieArray = strings.Split(cookieString, "; ")

	for _, cookie := range cookieArray {
		var cookieSplit = strings.Split(cookie, "=")
		cookieMap[cookieSplit[0]] = cookieSplit[1]
	}

	return cookieMap
}

func (route Route) group(pattern string, handler http.Handler) {
	route.router.Handle(pattern, handler)
}

func (route Route) methodeAllowed(methode string) {
	if route.req.Method != methode {
		http.Error(route.res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (route Route) allowCORS() {
	route.res.Header().Set("Access-Control-Allow-Origin", "*")
	route.res.Header().Set("Access-Control-Allow-Methods", "*")
	route.res.Header().Set("Access-Control-Allow-Headers", "*")
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
