package main

import (
	"fmt"
	"ms_music/app/internal/application"
	"ms_music/app/internal/repository"
	"time"
)

func main() {
	fmt.Println("Initializing server ...")

	// env
	// ...

	// app
	// -config

	// Empty config for now to use the default config
	cfg := &application.ConfigServerChi{
		Addr: ":8080",
		Port: 8080,
	}
	server := application.NewServerChi(cfg)

	// Initialize Repository
	server.Repository = &repository.Repository{}

	rp := server.Repository

	// Start a goroutine to get a new token to do the request in the Spotify API every hour
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for {
			application.GetNewTokenSpotify(rp)
			<-ticker.C
		}
	}()

	err := server.Run()

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	/*
		fmt.Println("Server started successfully")
		token, err := application.GetNewTokenSpotify(repository.Repository{})
		if err != nil {
			fmt.Println("Error getting Spotify token:", err)
		} else {
			fmt.Println("Spotify token:", token)
		}*/
}
