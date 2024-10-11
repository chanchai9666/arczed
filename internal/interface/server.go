package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"arczed/configs"
	"arczed/internal/interface/database"
)

type Server struct {
	port   int
	db     database.Service
	config configs.Config
}

func NewServer() *http.Server {

	configs, err := configs.LoadConfig()
	if err != nil {
		log.Fatalln("Load ENV ERROR : ", err.Error())
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:   port,
		db:     database.New(configs),
		config: *configs,
	}
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
