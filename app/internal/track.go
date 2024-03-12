package internal

import (
	"errors"
	"net/http"
)

// SpotifyResponse is the response from spotify
type SpotifyResponse struct {
	Artist struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			}
			Followers struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			Images []struct {
				Height int    `json:"height"`
				Width  int    `json:"width"`
				URL    string `json:"url"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			URI        string `json:"uri"`
		} `json:"items"`
	} `json:"artists"`
	Tracks struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
		Items []struct {
			Album struct {
				AlbumType   string `json:"album_type"`
				TotalTracks int    `json:"total_tracks"`
				ExtUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				AlbumID string `json:"id"`
				Images  []struct {
					Height int    `json:"height"`
					Width  int    `json:"width"`
					URL    string `json:"url"`
				} `json:"imagesFr"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				}
				ArtistID  string `json:"id"`
				Name      string `json:"name"`
				Followers struct {
					Href  string `json:"href"`
					Total int    `json:"total"`
				}
				Genres []string `json:"genres"`
				Images []struct {
					Height int    `json:"height"`
					Width  int    `json:"width"`
					URL    string `json:"url"`
				}
				Popularity int `json:"popularity"`
			} `json:"artists"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"tracks"`
	Albums struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
		Items []struct {
			AlbumType    string `json:"album_type"`
			TotalTracks  int    `json:"total_tracks"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				Width  int    `json:"width"`
				URL    string `json:"url"`
			} `json:"images"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
		} `json:"items"`
	} `json:"albums"`
}

var (
	// ErrArtistNotFound is returned when the artist is not found
	ErrArtistNotFound = errors.New("artist not found")
	// ErrGetIndex is returned when ocurrs an error getting the index
	ErrGetIndex = errors.New("error getting index")
	// ErrMaxScore is returned when the track is not indexed
	ErrMaxScore = errors.New("error max score")
	// ErrorIndexingDate is returned when the track is not indexed
	ErrorIndexingDate = errors.New("error indexing track")
	//ErrCreatingIndex is returned when the index is not created
	ErrCreatingIndex = errors.New("error creating index")
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
	GetTrackByName(trackName string) (Track, error)             // Get a track by name
	GetTrackByNameAndArtistName(artist string) ([]Track, error) // Get tracks by artist
	GetTrackByAlbum(album string) ([]Track, error)              // Get tracks by album
	GetTrackByGenre(genre string) ([]Track, error)              // Get tracks by genre
	GetAllTracksPopularityElasticSearch(minPopularity, maxPopularity int, indexName string) ([]SpotifyResponse, error)
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
	GetAllTrackByAlbum() http.HandlerFunc       // Get all tracks by album
	GetAllTrackByGenre() http.HandlerFunc       // Get all tracks by genre
	GetAllTrackByPopularity() http.HandlerFunc  // Get all tracks by popularity
	GetAllTrackByReleaseDate() http.HandlerFunc // Get all tracks by release date

	// Search Methods
	SearchTrack() http.HandlerFunc // Search tracks according some query parameters
	// Create Methods
	CreateTrack() http.HandlerFunc

	// Update Methods
	UpdateTrack() http.HandlerFunc

	// Delete Methods
	DeleteTrack() http.HandlerFunc
}
