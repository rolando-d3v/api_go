package documento

import "net/http"

func DocumentoMux() http.Handler{
	docuMux := http.NewServeMux()
	docuMux.Handle("GET /id-doc/{id}", http.HandlerFunc(GetIdDocumento))
	docuMux.Handle("GET /all-doc", http.HandlerFunc(GetAllDocumento))
	return http.StripPrefix("/documento", docuMux)
}

