package documento

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetIdDocumento(w http.ResponseWriter, r *http.Request) {

    id_string := r.PathValue("id")
    id, err := strconv.Atoi(id_string)
    if err != nil {
        http.Error(w, "Invalid to parse id", http.StatusBadRequest)
    }

    fmt.Printf("Heyyy Am getting called From Controllers")
    fmt.Fprintf(w, "numero: %d", id)

    // w.Write([]byte(`{"message": "documento one id"}`))

}
func GetAllDocumento(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"message": "documento all "}`))

}