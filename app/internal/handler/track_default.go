package handler

import (
	"fmt"
	"ms_music/app/internal"
	"ms_music/app/internal/service"
	"ms_music/app/platform/web/response"
	"net/http"
	"strconv"
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
		// Bussiness Logic
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		// Convert string to int
		startInt, err := strconv.Atoi(start)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Bad Request: start should be a number")
			return
		}

		endInt, err := strconv.Atoi(end)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Bad Request: end should be a number")
			return
		}

		if startInt < 0 || endInt < 0 || endInt < startInt {
			response.Error(w, http.StatusBadRequest, "Bad Request: start and end should be positive numbers")
			return
		}

		// Use the service to get all the tracks

		tracks, err := t.TrackService.GetTrackByPopularity(startInt, endInt)

		if err != nil {
			switch err {
			case internal.ErrTrackNotFound:
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		// handle response here
		response.JSON(w, http.StatusOK, tracks)
		fmt.Println(start, end)
	}
}

// Get alls tracks by release date
func (t TrackHandler) GetAllTrackByReleaseDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bussiness Logic
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")

		if start == "" || end == "" {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}
		// Use the service to get all the tracks

		tracks, err := t.TrackService.GetTrackByReleaseDate(start, end)
		fmt.Println("error", err)
		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				response.Error(w, http.StatusBadRequest, err.Error())
			case internal.ErrTrackNotFound:
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		response.JSON(w, http.StatusOK, tracks)

	}

}

// GetAllArtist returns all the artists
func (t TrackHandler) GetAllArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bussiness logic
		artistName := r.URL.Query().Get("name")
		if artistName == "" {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}
		// Use the service to get all the artists
		artistList, err := t.TrackService.GetAllArtist(artistName)

		if err != nil {
			switch err {
			case internal.ErrArtistNotFound:
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		// Return the artists
		response.JSON(w, http.StatusOK, artistList)
	}
}

func (t TrackHandler) GetID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Bussiness logic
		trackID := r.URL.Query().Get("track")
		albumID := r.URL.Query().Get("album")
		artistID := r.URL.Query().Get("artist")

		if trackID != "" && albumID == "" && artistID == "" {
			track, err := t.TrackService.GetTrackByID(trackID)
			fmt.Println("error track", err)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.JSON(w, http.StatusOK, track)
			return
		}

		if trackID == "" && albumID != "" && artistID == "" {
			album, err := t.TrackService.GetAlbumByID(albumID)
			fmt.Println("error album", err)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.JSON(w, http.StatusOK, album)
			return
		}

		if trackID == "" && albumID == "" && artistID != "" {
			artist, err := t.TrackService.GetArtistByID(artistID)
			fmt.Println("error artist", err)
			if err != nil {
				response.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.JSON(w, http.StatusOK, artist)
			return
		}
		fmt.Println("error paso")
		// Return the artists
		response.Error(w, http.StatusInternalServerError, "Internal Server Error")

	}
}
