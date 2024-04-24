package main

import (
	"fmt"
	"net/http"

	"github.com/rolando-d3v/api_go/api/documento"
	"github.com/rolando-d3v/api_go/api/user"
)

func main() {

	// router := http.NewServeMux()
	// router.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("notas usuarios"))
	// })

	mux := http.NewServeMux()

	// define routes
	mux.Handle("/user/", user.UserMux())
	mux.Handle("/documento/", documento.DocumentoMux())
	// mux.Handle("/api/v1/auth/", AuthMux())

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	fmt.Println("Server run " + server.Addr)
	server.ListenAndServe()
}

// func AuthMux() http.Handler {
// 	authMux := http.NewServeMux()
// 	authMux.Handle("/signup", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		res.Write([]byte("You All Signed Up Misterr Wick 🧘🏽🧘🏽🧘🏽"))
// 	}))
// 	authMux.Handle("/resendVerificationEmail", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		res.Write([]byte("Your Access Has Been Resent Misterr Wick 🧘🏽🧘🏽🧘🏽"))
// 	}))

// 	return http.StripPrefix("/api/v1/auth", authMux)
// }
