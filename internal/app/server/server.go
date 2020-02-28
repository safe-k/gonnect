package server

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type server struct {
	db *sqlx.DB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .server file found")
	}
}

func Serve() {
	dbUrl := os.Getenv("DB")
	if dbUrl == "" {
		log.Fatalln("DB environment variable not set")
	}

	DB, err := sqlx.Connect("mysql", dbUrl)
	if err != nil {
		log.Fatalln("Could not connect to DB", err)
	}

	defer DB.Close()

	s := &server{DB}

	http.Handle("/player/connect", s.handlePlayerConnect())

	defer func() {
		if err := recover(); err != nil {
			log.Println("Could not complete request:", err)
			return
		}
	}()

	const port = ":5000"
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
