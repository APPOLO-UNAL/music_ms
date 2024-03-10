// This script contains the configuration of the chi router
package application

import (
	"fmt"
	"ms_music/app/internal/handler"
	"ms_music/app/internal/repository"
	"ms_music/app/internal/service"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/chi/v5"
)

// ConfigServerChi configures the chi router
type ConfigServerChi struct {
	// Addr is the address where the server will listen to
	Addr string
	// ReadTimeout is the maximum duration for reading the entire request, including the body
	ReadTimeout int
	// WriteTimeout is the maximum duration before timing out writes of the response
	WriteTimeout int
	// IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled
	IdleTimeout int
	// Port is the port where the server will listen to
	Port int
}

// NewConfigServerChi creates a new ConfigServerChi
func NewServerChi(cfg *ConfigServerChi) *ConfigServerChi {
	// Default Config
	defaultConf := &ConfigServerChi{
		Addr:         ":8080",
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  15,
		Port:         8080,
	}
	if cfg != nil {
		defaultConf = cfg
	}
	return defaultConf
}

// Run runs the server

func (s *ConfigServerChi) Run() (err error) {

	// Depedencies
	// - Database Connection
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(fmt.Sprintf("Error creating the client: %s", err))
	}
	fmt.Println(es)
	// Repository
	rp := repository.NewTrackRepository()

	// Service
	sv := service.NewTrackService(rp)

	// Handler
	hd := handler.NewTrackHandler(sv)

	// Router

	router := chi.NewRouter()

	// Middlewares
	// ... No implemented yet

	// Endpoints
	// - Endpoints Music
	buildEndpointMusic(router, hd)
	// - Endpoints Artist
	buildEndpointArtist(router, hd)
	// - Endpoints Album
	buildEndpointAlbum(router, hd)

	err = http.ListenAndServe(s.Addr, router)
	return
}

// Function to build the endpoints of the music
func buildEndpointMusic(router *chi.Mux, hd handler.TrackHandler) {

	// Configure the handler
	hdMusic := hd

	// Router group for the music
	router.Route("/api/v1", func(r chi.Router) {
		// Endpoint to get the music
		r.Get("/music", hdMusic.GetAllTracks())
		// Endpoint to get the music by Artist
		r.Get("/music/{id}", hdMusic.GetAllTrackByArtist())
		// Endpoint to get the music by Album
		r.Get("/music/album/{album}", hdMusic.GetAllTrackByAlbum())
		// Endpoint to get the music by Genre
		r.Get("/music/genre/{genre}", hdMusic.GetAllTrackByGenre())
		// Endpoint to get the music by Popularity
		r.Get("/music/popularity/{genre}", hdMusic.GetAllTrackByPopularity())
		// Endpoint to get the music by Duration
		r.Get("/music/duration/{duration}", hdMusic.GetAllTrackByDuration())
		// Endpoint to get the music by Release Date
		r.Get("/music/releasedate/{release_date}", hdMusic.GetAllTrackByReleaseDate())
		// Endpoint to get the music between Duration
		r.Get("/music/duration/{start_duration}/{end_duration}", hdMusic.GetAllTrackBetweenDuration())
		// Endpoint to get the music available in a country
		r.Get("/music/markets/{market}", hdMusic.GetAllTrackAvailableInMarket())

		// Endpoints using query params
		// Search
		r.Get("/music/search", hdMusic.SearchTrack()) // Search by artist and album

		// Create Method
		r.Post("/music", hdMusic.CreateTrack())

		// Update Method
		r.Put("/music/{id}", hdMusic.UpdateTrack())

		// Delete Method
		r.Delete("/music/{id}", hdMusic.DeleteTrack())

	})
}

// Function to build the endpoints of the artist
func buildEndpointArtist(router *chi.Mux, hd handler.TrackHandler) {
}

// Function to build the endpoints of the album
func buildEndpointAlbum(router *chi.Mux, hd handler.TrackHandler) {
}
