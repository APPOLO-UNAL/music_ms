// This script contains the configuration of the chi router
package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ms_music/app/internal/handler"
	"ms_music/app/internal/repository"
	"ms_music/app/internal/service"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
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
	// Repository is the repository
	Repository *repository.Repository
}

// Func to get a new token from spotify
func GetNewTokenSpotify(rp *repository.Repository) (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("no access_token found in response")
	}
	rp.SetToken(token)
	fmt.Println("The Spotify token is updated successfully", token)
	return token, nil
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

	// Repository
	rp := repository.NewTrackRepository(es)
	// Set the repository in the config
	s.Repository = &rp

	// Get a new Spotify token
	_, err = GetNewTokenSpotify(s.Repository)
	if err != nil {
		panic(fmt.Sprintf("Error getting Spotify token: %s", err))
	}

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
		// Endpoint to get the track by name
		r.Get("/music", hdMusic.GetTrackByName())
		// Endpoint to get the music
		r.Get("/music/tracks", hdMusic.GetAllTracks())
		// Endpoint to get the music by Artist
		r.Get("/music/tracks/artist", hdMusic.GetAllTrackByArtist())
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
