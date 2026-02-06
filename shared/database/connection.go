package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	SSLMode     string
	SSLCertPath string
}

func NewConfig() *Config {
	return &Config{
		Host:        getEnv("DB_HOST", "localhost"),
		Port:        getEnv("DB_PORT", "5432"),
		User:        getEnv("DB_USER", "postgres"),
		Password:    getEnv("DB_PASSWORD", "postgres"),
		DBName:      getEnv("DB_NAME", "postgres"),
		SSLMode:     getEnv("DB_SSLMODE", "disable"),
		SSLCertPath: getEnv("DB_CA_CERT", ""),
	}
}

func (c *Config) Connect() (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s sslrootcert=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
		c.SSLCertPath,
	)

	log.Println("DB connection string:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos exitosa")
	return db, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
