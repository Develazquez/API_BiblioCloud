package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	prestamoRoutes "biblioteca-api/prestamos/infrastructure/routes"
	recursoRoutes "biblioteca-api/recursos/infrastructure/routes"
	"biblioteca-api/shared/config"
	"biblioteca-api/shared/database"
	usuarioRoutes "biblioteca-api/usuarios/infrastructure/routes"
)

func main() {
	config.LoadEnv()

	dbConfig := database.NewConfig()
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	usuarioRoutes.UsuarioRoutes(router, db)
	recursoRoutes.RecursoRoutes(router, db)
	prestamoRoutes.PrestamoRoutes(router, db)

	router.Use(corsMiddleware)


	port := config.GetPort()
	log.Println("Servidor iniciado en http://localhost" + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


