package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {

	//Load enviroment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro em carregar as variaveis de ambiente : %v", err)
	}

	//Get enviroment variables
	dbHost := os.Getenv("DATABASE_HOSTNAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	//Mount connection string using enviroment variables
	connection := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUser, dbName, dbPass, dbHost, dbPort,
	)

	// Connect to database
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Check connection with database
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro com encontrado ao tentar conectar com o banco de dados: %v", err)
	}

	return db

}
