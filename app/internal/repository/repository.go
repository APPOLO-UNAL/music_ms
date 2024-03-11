package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

// Repository is the repository
type Repository struct {
	es        *elasticsearch.Client
	authToken string
}

// SpotifyResponse is the response from spotify
type SpotifyResponse struct {
	Tracks struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
		Items []struct {
			Album struct {
				AlbumType        string   `json:"album_type"`
				TotalTracks      int      `json:"total_tracks"`
				AvailableMarkets []string `json:"available_markets"`
				ExtUrls          struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				AlbumID string `json:"id"`
				Images  []struct {
					Height int    `json:"height"`
					Width  int    `json:"width"`
					URL    string `json:"url"`
				} `json:"images"`
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
		} `json:"items"`
	} `json:"tracks"`
}

// NewTrackRepository returns a new TrackRepository
func NewTrackRepository(es *elasticsearch.Client) Repository {
	repo := Repository{
		es: es,
	}

	repo.VerifyIndices()
	return repo
}

// SetToken sets the token to the repository
func (repo *Repository) SetToken(token string) {
	repo.authToken = token
}

// Function to verify if the indices exist, if not, create them
func (repo *Repository) VerifyIndices() {
	indices := []string{"track", "album", "artist"}

	for _, index := range indices {
		res, err := repo.es.Indices.Exists([]string{index})
		if err != nil {
			log.Fatalf("Error checking if index exists: %s", err)
		}

		// If the index does not exist, create it
		if res.StatusCode == 404 {
			req := esapi.IndicesCreateRequest{
				Index: index,
			}

			res, err := req.Do(context.Background(), repo.es)
			if err != nil {
				log.Fatalf("Error creating index: %s", err)
			}
			if res.IsError() {
				log.Fatalf("Error response from server: %s", res.String())
			}
		}
	}
}

// GetTrackByName returns a track by name
func (repo *Repository) GetTrackByName(name string) (SpotifyResponse, error) {
	// Create the request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	if err != nil {
		return SpotifyResponse{}, err
	}
	// Add the query parameters
	q := req.URL.Query()
	q.Add("q", name)
	q.Add("type", "track")
	req.URL.RawQuery = q.Encode()

	// Add the authorization header
	req.Header.Add("Authorization", "Bearer "+repo.authToken)

	// Make the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return SpotifyResponse{}, err
	}
	defer res.Body.Close()
	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SpotifyResponse{}, err
	}
	// Unmarshal the response
	var spotifyResponse SpotifyResponse
	err = json.Unmarshal(body, &spotifyResponse)
	if err != nil {
		return SpotifyResponse{}, err
	}
	// Print the response
	return spotifyResponse, nil
}
