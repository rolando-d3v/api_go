package user

import "net/http"

func UserMux() http.Handler{
	userMux := http.NewServeMux()
	userMux.Handle("GET /profile", http.HandlerFunc(GetIdUser))
	userMux.Handle("GET /all-user", http.HandlerFunc(GetAllUser))
	return http.StripPrefix("/user", userMux)
}

