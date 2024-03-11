package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// NewTrackRepository returns a new TrackRepository
func NewTrackRepository(es *elasticsearch.Client) Repository {
	repo := Repository{
		es: es,
		mapping: `{
			"mappings": {
				"properties": {
					"artists": {
						"properties": {
							"href": {
								"type": "text"
							},
							"total": {
								"type": "integer"
							},
							"items": {
								"properties": {
									"external_urls": {
										"properties": {
											"spotify": {
												"type": "text"
											}
										}
									},
									"followers": {
										"properties": {
											"href": {
												"type": "text"
											},
											"total": {
												"type": "integer"
											}
										}
									},
									"genres": {
										"type": "text"
									},
									"href": {
										"type": "text"
									},
									"images": {
										"properties": {
											"height": {
												"type": "integer"
											},
											"width": {
												"type": "integer"
											},
											"url": {
												"type": "text"
											}
										}
									},
									"name": {
										"type": "text"
									},
									"popularity": {
										"type": "integer"
									},
									"uri": {
										"type": "text"
									}
								}
							}
						}
					},
					"tracks": {
						"properties": {
							"href": {
								"type": "text"
							},
							"total": {
								"type": "integer"
							},
							"items": {
								"properties": {
									"album": {
										"properties": {
											"album_type": {
												"type": "text"
											},
											"total_tracks": {
												"type": "integer"
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
														"type": "integer"
													},
													"width": {
														"type": "integer"
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
												"type": "date", // Cambiado de "text" a "date"
												"format": "strict_date_optional_time||epoch_millis" // Asegúrate de que el formato de fecha coincide con tus datos
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
														"type": "integer"
													}
												}
											},
											"genres": {
												"type": "text"
											},
											"images": {
												"properties": {
													"height": {
														"type": "integer"
													},
													"width": {
														"type": "integer"
													},
													"url": {
														"type": "text"
													}
												}
											},
											"popularity": {
												"type": "integer"
											}
										}
									},
									"external_urls": {
										"properties": {
											"spotify": {
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
		} `,
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

func (repo *Repository) IndexTrackByArtist(indexName string, response SpotifyResponse) error {
	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Create an index request
	req := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(responseJSON)),
		Refresh: "true",
	}

	// Perform the index request
	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println(res.IsError())
		fmt.Println(res.String())
		return internal.ErrCreatingIndex
	}

	return nil
}

// GetTrackByName returns a track by name from Spotify
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
			"multi_match": map[string]interface{}{
				"query":     trackName,
				"fields":    []string{"tracks.items.album.name"},
				"fuzziness": "AUTO",
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

// GetTrackByArtistElasticSearch retrieves tracks from Elasticsearch by artist and track name
func (repo *Repository) GetTrackByArtistElasticSearch(indexName string, artistName string, trackName string) (SpotifyResponse, float64, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"artists.items.name": artistName,
						},
					},
					{
						"match": map[string]interface{}{
							"tracks.items.album.name": trackName,
						},
					},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return SpotifyResponse{}, 0, err
	}

	res, err := repo.es.Search(
		repo.es.Search.WithContext(context.Background()),
		repo.es.Search.WithIndex(indexName),
		repo.es.Search.WithBody(&buf),
		repo.es.Search.WithTrackTotalHits(true),
		repo.es.Search.WithPretty(),
	)
	if err != nil {
		return SpotifyResponse{}, 0, err
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return SpotifyResponse{}, 0, err
	}

	hits, ok := r["hits"].(map[string]interface{})
	if !ok || hits == nil {
		return SpotifyResponse{}, 0, fmt.Errorf("no hits from Elasticsearch")
	}

	maxScore, ok := hits["max_score"].(float64)
	if !ok {
		return SpotifyResponse{}, 0, fmt.Errorf("no max_score from Elasticsearch")
	}

	var spotifyResponses SpotifyResponse
	if err := json.Unmarshal([]byte(hits["hits"].([]interface{})[0].(map[string]interface{})["_source"].(string)), &spotifyResponses); err != nil {
		return SpotifyResponse{}, 0, err
	}

	return spotifyResponses, maxScore, nil
}

// GetTrackByArtist retrieves tracks from Spotify by artist and track name
func (repo *Repository) GetTrackByArtist(artistName string, trackName string) (SpotifyResponse, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	if err != nil {
		return SpotifyResponse{}, err
	}

	q := req.URL.Query()
	q.Add("q", artistName+" "+trackName)
	q.Add("type", "track")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer "+repo.authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SpotifyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SpotifyResponse{}, err
	}
	var spotifyResponses SpotifyResponse
	err = json.Unmarshal(body, &spotifyResponses)
	if err != nil {
		return SpotifyResponse{}, err
	}

	return spotifyResponses, nil
}

// GetAllTracksElasticSearch retrieves all tracks from Elasticsearch
func (repo *Repository) GetAllTracksElasticSearch(indexName string) ([]SpotifyResponse, error) {

	// Define the search query
	query := `{
        "query": {
            "match_all": {}
        }
    }`

	// Perform the search request
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(query),
	}

	// Make the request
	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return nil, fmt.Errorf("failed to perform search request: %w", err)
	}
	defer res.Body.Close()

	// Check if response is not OK
	if res.IsError() {
		return nil, fmt.Errorf("response status: %s", res.Status())
	}

	// Decode the response body
	var response struct {
		Hits struct {
			Hits []struct {
				Source SpotifyResponse `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	// Extract SpotifyResponse from hits
	var tracks []SpotifyResponse
	for _, hit := range response.Hits.Hits {
		tracks = append(tracks, hit.Source)
	}

	return tracks, nil
}

// GetAllTracksByAlbumElasticSearch retrieves all tracks from Elasticsearch by album
func (repo *Repository) GetAllTracksByAlbumElasticSearch(indexName string, albumName string) ([]SpotifyResponse, error) {
	var tracks []SpotifyResponse

	// Define the search query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"tracks.items.album.name": map[string]interface{}{
					"query": albumName,
					"slop":  10,
				},
			},
		},
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return tracks, internal.ErrBadRequest
	}

	// Perform the search request
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  bytes.NewReader(queryJSON),
	}

	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		return tracks, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return tracks, internal.ErrTrackNotFound
	}

	// Parse the response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return tracks, internal.ErrInternalServerError
	}

	// Extract the tracks from the response
	if hits, ok := r["hits"].(map[string]interface{}); ok {
		if hitsArray, ok := hits["hits"].([]interface{}); ok {
			for _, hit := range hitsArray {
				if source, ok := hit.(map[string]interface{})["_source"]; ok {
					var track SpotifyResponse
					trackJSON, err := json.Marshal(source)
					if err != nil {
						return tracks, internal.ErrInternalServerError
					}
					if err := json.Unmarshal(trackJSON, &track); err != nil {
						return tracks, internal.ErrInternalServerError
					}
					tracks = append(tracks, track)
				}
			}
		}
	}

	return tracks, nil
}

// GetAllTracksByAlbum retrieves all tracks from Spotify by album
func (repo *Repository) GetAllTracksByAlbum(albumName string) (SpotifyResponse, error) {
	// Replace spaces with +
	albumName = strings.ReplaceAll(albumName, " ", "+")

	// URL request
	url := "https://api.spotify.com/v1/search?q=" + albumName + "&type=album,track"

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SpotifyResponse{}, internal.ErrBadRequest
	}

	// Append the authorization header
	req.Header.Add("Authorization", "Bearer "+repo.authToken)

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SpotifyResponse{}, internal.ErrInternalServerError
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SpotifyResponse{}, internal.ErrInternalServerError
	}

	// Deserialize the response body
	var spotifyResponse SpotifyResponse
	err = json.Unmarshal(body, &spotifyResponse)
	if err != nil {
		return SpotifyResponse{}, internal.ErrInternalServerError
	}

	// Return the response
	return spotifyResponse, nil
}
