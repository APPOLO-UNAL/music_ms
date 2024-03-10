package main

import (
	"fmt"
	"ms_music/app/internal/application"
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
	// Start a goroutine to get a new token every hour
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for {
			application.GetNewToken()
			<-ticker.C
		}
	}()
	err := server.Run()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
