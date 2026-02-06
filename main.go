package main

import (
	"log"

	"github.com/gin-gonic/gin"

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

	router := gin.Default()

	// Aplicar CORS middleware
	router.Use(corsMiddleware())

	// Registrar rutas de cada entidad
	usuarioRoutes.UsuarioRoutes(router, db)
	recursoRoutes.RecursoRoutes(router, db)
	prestamoRoutes.PrestamoRoutes(router, db)

	// Health check
	router.GET("/health", healthCheck)

	// Iniciar servidor
	port := config.GetPort()
	log.Println("Servidor iniciado en http://localhost" + port)
	router.Run(port)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
