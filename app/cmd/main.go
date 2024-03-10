package main

import (
	"fmt"
	"ms_music/app/internal/application"
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
	err := server.Run()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
