package internal

import (
	"errors"
	"net/http"
)

var (
	// ErrCreateTrack is returned when the track is not created
	ErrCreateTrack = errors.New("error creating track")
	// ErrBadRequest is returned when the request is invalid
	ErrBadRequest = errors.New("bad request")
	// ErrTrackNotFound is returned when the track is not found
	ErrTrackNotFound = errors.New("track not found")
	// InternalServerError is returned when the track is not found
	ErrInternalServerError = errors.New("internal server error")
)

// This script contains the logic to handle, service and repostory to  the track entity
type Track struct {
	Href             string   // The Spotify URL for the track.
	Limit            int      // The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Album            Album    // The album on which the track appears. The album object includes a link in href to full information about the album.
	Artists          []Artist // The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	AvailableMarkets []string // The markets in which the track is available: ISO 3166-1 alpha-2 country codes.
	DiscNumber       int      // The disc number (usually 1 unless the album consists of more than one disc).
	DurationMs       int      // The track length in milliseconds.
	Explicit         bool     // Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	SpotifyURL       string   // The Spotify URL for the track.
	SpotifyID        string   // The Spotify ID for the track.
	Name             string   // The name of the track.
	Popularity       int      // The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	PreviewURL       string   // A link to a 30 second preview (MP3 format) of the track. Can be null.
	TrackNumber      int      // The number of the track. If an album has several discs, the track number is the number on the specified disc.

}

// TrackService is the interface that provides track methods
type TrackService interface {
	// Get Methods
	GetTrackByName(trackName string) (Track, error)                        // Get a track by name
	GetTrack(id string) (Track, error)                                     // Get a track
	GetTrackByArtist(artist string) ([]Track, error)                       // Get tracks by artist
	GetTrackByAlbum(album string) ([]Track, error)                         // Get tracks by album
	GetTrackByGenre(genre string) ([]Track, error)                         // Get tracks by genre
	GetTrackByPopularity(popularity int) ([]Track, error)                  // Get tracks by popularity
	GetTrackByDuration(duration int) ([]Track, error)                      // Get tracks by duration
	GetTrackByReleaseDate(releaseDate string) ([]Track, error)             // Get tracks by release date
	GetTrackBetweenDuration(duration1 int, duration2 int) ([]Track, error) // Get tracks between duration
	GetTrackAvailableInMarket(market string) ([]Track, error)              // Get tracks available in market

	// Create Methods
	CreateTrack(track Track) (Track, error)

	// Update Methods
	UpdateTrack(track Track) (Track, error)

	// Delete Methods
	DeleteTrack(id string) error
}

// TrackHandler is the interface that provides track methods
type TrackHandler interface {
	// Get Methods
	GetTrackByName() http.HandlerFunc      // Get a track
	GetAllTracks() http.HandlerFunc        // Get all tracks
	GetAllTrackByArtist() http.HandlerFunc // Get tracks by artist

	// General Methods Get
	GetAllTrackByAlbum() http.HandlerFunc           // Get all tracks by album
	GetAllTrackByGenre() http.HandlerFunc           // Get all tracks by genre
	GetAllTrackByPopularity() http.HandlerFunc      // Get all tracks by popularity
	GetAllTrackByDuration() http.HandlerFunc        // Get all tracks by duration
	GetAllTrackByReleaseDate() http.HandlerFunc     // Get all tracks by release date
	GetAllTrackBetweenDuration() http.HandlerFunc   // Get all tracks between duration
	GetAllTrackAvailableInMarket() http.HandlerFunc // Get all tracks available in market

	// Search Methods
	SearchTrack() http.HandlerFunc // Search tracks according some query parameters
	// Create Methods
	CreateTrack() http.HandlerFunc

	// Update Methods
	UpdateTrack() http.HandlerFunc

	// Delete Methods
	DeleteTrack() http.HandlerFunc
}
