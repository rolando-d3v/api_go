package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rolando-d3v/api_go/api/documento"
	"github.com/rolando-d3v/api_go/api/user"
)

func main() {

	// router := http.NewServeMux()
	// router.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("notas usuarios"))
	// })

	// value := os.Getenv("path")  //obtiene variables del sistema operativo
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error: No hay variable")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//obtiene variables del sistema operativo

	rol := os.Getenv("ROLANDO")

	fmt.Println("rolando", rol)

	mux := http.NewServeMux()

	// define routes
	mux.Handle("/user/", user.UserMux())
	mux.Handle("/documento/", documento.DocumentoMux())

	pepe := fmt.Sprintf("hola %s peru %.2f", "rolando", 2.4)
	fmt.Println(pepe)

	// Configurar el servidor
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Mostrar en consola el puerto en el que se ejecuta
	fmt.Println("Server running on http://0.0.0.0:" + port)

	// Iniciar el servidor y manejar errores
	log.Fatal(server.ListenAndServe())

	// result := fmt.Sprintf("%s%s", ":", port)
	// server := &http.Server{
	// 	Addr: result,
	// 	// Addr: ":" + port,
	// 	// Addr:    ":4000",
	// 	Handler: mux,
	// }
	// fmt.Println("Server run " + server.Addr)
	// server.ListenAndServe()
	// log.Fatal(server)
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

// package main

// import (
//     "crypto/rand"
//     "crypto/rsa"
//     "crypto/x509"
//     "encoding/pem"
//     "fmt"
//     "os"
// )

// func main() {
//     // Generar la llave privada
//     privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//     if err != nil {
//         fmt.Println("Error generando la llave privada:", err)
//         return
//     }

//     // Codificar la llave privada en formato PEM
//     privateKeyPEM := pem.EncodeToMemory(
//         &pem.Block{
//             Type:  "RSA PRIVATE KEY",
//             Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
//         },
//     )

//     // Escribir la llave privada en un archivo
//     err = os.WriteFile("private_key.pem", privateKeyPEM, 0600)
//     if err != nil {
//         fmt.Println("Error escribiendo la llave privada en el archivo:", err)
//         return
//     }
//     fmt.Println("Llave privada guardada en private_key.pem")

//     // Generar la llave pública correspondiente
//     publicKey := &privateKey.PublicKey

//     // Codificar la llave pública en formato PEM
//     publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
//     if err != nil {
//         fmt.Println("Error convirtiendo la llave pública a bytes:", err)
//         return
//     }

//     publicKeyPEM := pem.EncodeToMemory(
//         &pem.Block{
//             Type:  "PUBLIC KEY",
//             Bytes: publicKeyBytes,
//         },
//     )

//     // Escribir la llave pública en un archivo
//     err = os.WriteFile("public_key.pem", publicKeyPEM, 0644)
//     if err != nil {
//         fmt.Println("Error escribiendo la llave pública en el archivo:", err)
//         return
//     }
//     fmt.Println("Llave pública guardada en public_key.pem")
// }
