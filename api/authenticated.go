package api

import (
	"fmt"
	"net/http"
)

func AuthMux() http.Handler {
	authMux := http.NewServeMux()

	authMux.Handle("/signup", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Signup")

		res.Write([]byte("You All Signed Up Misterr Wick 🧘🏽🧘🏽🧘🏽"))
	}))

	authMux.Handle("/resendVerificationEmail", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Print("resendVerificationEmail")
		res.Write([]byte("Your Access Has Been Resent Misterr Wick 🧘🏽🧘🏽🧘🏽"))
	}))

	return http.StripPrefix("/api/v1/auth", authMux)
}

// url := authConfig.auth.AuthURL(authConfig.state)
// // fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

// w.Header().Set("Content-Type", "text/html")
// w.Write([]byte(fmt.Sprintf("<a href='%s'>Log in to Spotify</a>", url)))

// // wait for auth to complete

// client := <-authChannel

// fmt.Printf("Login Completed! %v", client)

// // use the client to make calls that require authorization
// user, err := client.CurrentUser(context.Background())
// if err != nil {
// 	log.Fatal(err)
// }
// // fmt.Println("\nYou are logged in as:", user.ID)

// w.Write([]byte("<p>You are logged in as:" + user.ID + "</p>"))

// var url = fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=user-read-private%%20user-read-email&state=34fFs29kd09", spotifyID, "http://localhost:8080/api/me")
// http.Redirect(w, r, url, http.StatusSeeOther)
