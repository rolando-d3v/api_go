package user

import (
    "fmt"
    "net/http"
)

func GetIdUser(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"message": "working one user..."}`))

}
func GetAllUser(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"message": "working users..."}`))

}