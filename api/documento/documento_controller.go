package documento

import (
    "fmt"
    "net/http"
)

func GetIdDocumento(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"message": "documento one"}`))

}
func GetAllDocumento(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"message": "documento all "}`))

}