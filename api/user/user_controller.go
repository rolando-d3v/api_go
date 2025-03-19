package user

import (
    // "fmt"
    "net/http"
    "fmt"
	"time"
)

func GetIdUser(w http.ResponseWriter, r *http.Request) {

    // fmt.Printf("Heyyy Am getting called From Controllers")

    w.Write([]byte(`{"msg": "React one user 2022..."}`))

}
func GetAllUser(w http.ResponseWriter, r *http.Request) {

    // fmt.Printf("Heyyy Am getting called From Controllers")

    	// Iniciar el cronómetro
	start := time.Now()

	// Crear un slice con 9,000,000 elementos
	registros := make([]int, 90000000)
	for i := 0; i < 90000000; i++ {
		registros[i] = i + 1
	}

	// Leer los registros con un bucle `for`
	for i := 0; i < len(registros); i++ {
		dato := registros[i] // Simulación de lectura
		_ = dato             // Evitar warnings de variable no utilizada
	}

	// Calcular el tiempo total transcurrido
	elapsed := time.Since(start)
	fmt.Printf("Tiempo total: %s\n", elapsed)

    w.Write([]byte(`{"message": "working users...", "pepe": 15}`))

}