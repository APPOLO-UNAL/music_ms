package handler

import (
	"fmt"
	"ms_music/app/internal"
	"ms_music/app/internal/service"
	"ms_music/app/platform/web/response"
	"net/http"
)

type TrackJSONDefault struct {
	AlbumID          string            `json:"album_id"`
	Href             string            `json:"href"`
	Limit            int               `json:"limit"`
	Album            internal.Album    `json:"album"`
	Artist           []internal.Artist `json:"artist"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	SpotifyURL       string            `json:"spotify_url"`
	SpotifyID        string            `json:"spotify_id"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	TrackNumber      int               `json:"track_number"`
}

type TrackHandler struct {
	TrackService service.TrackService
}

// NewTrackService returns a new TrackService
func NewTrackHandler(trackService service.TrackService) TrackHandler {
	return TrackHandler{
		TrackService: trackService,
	}
}

// Get Methods
// Get a track
func (t TrackHandler) GetTrackByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trackName := r.URL.Query().Get("name")

		if trackName == "" {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}
		track, err := t.TrackService.GetTrackByName(trackName)
		if err != nil {
			switch err {
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		if track == nil {
			response.Error(w, http.StatusNotFound, "Track not found")
		}
		fmt.Println(track)
		response.JSON(w, http.StatusOK, track)

	}

}

// Get all tracks
func (h TrackHandler) GetAllTracks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bussiness logic
		// ...

		// Use the service to get all the tracks
		trackList, err := h.TrackService.GetAllTracks()

		if err != nil {
			switch err {
			case internal.ErrTrackNotFound:
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		// Return the tracks
		response.JSON(w, http.StatusOK, trackList)
	}
}

// Get tracks by artist
func (t TrackHandler) GetAllTrackByArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trackName := r.URL.Query().Get("track")
		artistName := r.URL.Query().Get("artist")

		if trackName == "" || artistName == "" {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}

		track, err := t.TrackService.GetTrackByArtistAndName(trackName, artistName)
		if err != nil {
			switch err {
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		if track == nil {
			response.Error(w, http.StatusNotFound, "Track not found")
			return
		}
		response.JSON(w, http.StatusOK, track)
	}
}

// Get all tracks by album
func (t TrackHandler) GetAllTrackByAlbum() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bussiness logic

		album := r.URL.Query().Get("name")

		if album == "" {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Use the service to get all the tracks
		trackList, err := t.TrackService.GetAllTracksByAlbum(album)

		if err != nil {
			switch err {
			case internal.ErrTrackNotFound:
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		// Return the tracks
		response.JSON(w, http.StatusOK, trackList)
	}
}

// Get alls tracks by popularity
func (t TrackHandler) GetAllTrackByPopularity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		popularity := r.URL.Query().Get("popularity")
		//tracks, err := t.TrackService.GetTrackByPopularity(popularity)
		// handle response here
		fmt.Println(popularity)
	}
}

// Get alls tracks by release date
func (t TrackHandler) GetAllTrackByReleaseDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		releaseDate := r.URL.Query().Get("release_date")
		//tracks, err := t.TrackService.GetTrackByReleaseDate(releaseDate)
		// handle response here
		fmt.Println(releaseDate)
	}
}
