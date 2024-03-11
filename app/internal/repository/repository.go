package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"ms_music/app/internal"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

// Repository is the repository
type Repository struct {
	es         *elasticsearch.Client // Elasticsearch client
	authToken  string                // Auth token for spotify
	mapping    string                // Mapping for the track entity
	boolMapped bool                  // Bool to check if the mapping is already done
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
}

// NewTrackRepository returns a new TrackRepository
func NewTrackRepository(es *elasticsearch.Client) Repository {
	repo := Repository{
		es: es,
		mapping: `{
			"mappings": {
			  "properties": {
				"tracks": {
				  "properties": {
					"href": {
					  "type": "text"
					},
					"total": {
					  "type": "text"
					},
					"items": {
					  "properties": {
						"album": {
						  "properties": {
							"album_type": {
							  "type": "text"
							},
							"total_tracks": {
							  "type": "text"
							},
							"available_markets": {
							  "type": "text"
							},
							"external_urls": {
							  "properties": {
								"spotify": {
								  "type": "text"
								}
							  }
							},
							"id": {
							  "type": "text"
							},
							"imagesFr": {
							  "properties": {
								"height": {
								  "type": "text"
								},
								"width": {
								  "type": "text"
								},
								"url": {
								  "type": "text"
								}
							  }
							},
							"name": {
							  "type": "text"
							},
							"release_date": {
							  "type": "text"
							},
							"release_date_precision": {
							  "type": "text"
							}
						  }
						},
						"artists": {
						  "properties": {
							"external_urls": {
							  "properties": {
								"spotify": {
								  "type": "text"
								}
							  }
							},
							"id": {
							  "type": "text"
							},
							"name": {
							  "type": "text"
							},
							"followers": {
							  "properties": {
								"href": {
								  "type": "text"
								},
								"total": {
								  "type": "text"
								}
							  }
							},
							"genres": {
							  "type": "text"
							},
							"images": {
							  "properties": {
								"height": {
								  "type": "text"
								},
								"width": {
								  "type": "text"
								},
								"url": {
								  "type": "text"
								}
							  }
							},
							"popularity": {
							  "type": "text"
							}
						  }
						}
					  }
					}
				  }
				}
			  }
			}
		  }`,
		boolMapped: false,
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
	repo.boolMapped = true
}

// CreateIndexWithMapping creates an index with the given mapping
func (repo *Repository) CreateIndexWithMapping(indexName string) error {
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(repo.mapping),
	}

	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return internal.ErrCreatingIndex
	}

	return nil
}

// Delete index
func (repo *Repository) DeleteIndex(indexName string) error {
	// Delete the index
	deleteIndexReq := esapi.IndicesDeleteRequest{
		Index: []string{indexName},
	}

	res, err := deleteIndexReq.Do(context.Background(), repo.es)
	if err != nil {
		return err
	}
	if res.IsError() {
		return internal.ErrCreatingIndex
	}

	// Recreate the index with the correct mapping
	createIndexReq := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(repo.mapping), // replace with your mapping
	}

	res, err = createIndexReq.Do(context.Background(), repo.es)
	if err != nil {
		return err
	}
	if res.IsError() {
		return internal.ErrCreatingIndex
	}
	return nil
}

// IndexTrack indexes a track
func (repo *Repository) IndexTrack(indexName string, track SpotifyResponse) error {
	// Verify that the indices exist
	repo.VerifyIndices()

	if !repo.boolMapped {
		err := repo.CreateIndexWithMapping("tracks")
		if err != nil {
			return internal.ErrCreateTrack
		}
		repo.boolMapped = true
	}

	// Convert the track to JSON
	trackJSON, err := json.Marshal(track)
	if err != nil {
		return err
	}

	// Index the track
	req := esapi.IndexRequest{
		Index:   indexName,
		Body:    bytes.NewReader(trackJSON),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return internal.ErrorIndexingDate
	}

	return nil

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
		return SpotifyResponse{}, internal.ErrBadRequest
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

// GetTracksElasticSearch retrieves tracks from Elasticsearch and returns the max_score
func (repo *Repository) GetTracksElasticSearch(indexName string, trackName string) ([]SpotifyResponse, float64, error) {
	var tracks []SpotifyResponse

	// Define the search query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"simple_query_string": map[string]interface{}{
				"query": trackName,
			},
		},
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return tracks, 0, internal.ErrBadRequest
	}

	// Perform the search request
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  bytes.NewReader(queryJSON),
	}

	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return tracks, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return tracks, 0, internal.ErrTrackNotFound
	}

	// Parse the response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return tracks, 0, internal.ErrInternalServerError
	}

	// Extract the max_score from the response
	maxScore, ok := r["hits"].(map[string]interface{})["max_score"].(float64)
	if !ok {
		return tracks, 0, internal.ErrMaxScore
	}

	// Extract the tracks from the response
	if hits, ok := r["hits"].(map[string]interface{}); ok {
		if hitsArray, ok := hits["hits"].([]interface{}); ok {
			for _, hit := range hitsArray {
				if source, ok := hit.(map[string]interface{})["_source"]; ok {
					var track SpotifyResponse
					trackJSON, err := json.Marshal(source)
					if err != nil {
						return tracks, maxScore, internal.ErrInternalServerError
					}
					if err := json.Unmarshal(trackJSON, &track); err != nil {
						return tracks, maxScore, internal.ErrInternalServerError
					}
					tracks = append(tracks, track)
				}
			}
		}
	}

	return tracks, maxScore, nil
}
