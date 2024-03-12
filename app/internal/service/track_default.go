package service

import (
	"errors"
	"fmt"
	"ms_music/app/internal"
	"ms_music/app/internal/repository"
)

// This script contains the logic to handle, service and repostory to  the track entity
type TrackService struct {
	rp repository.Repository
}

// NewTrackService returns a new TrackService
func NewTrackService(trackRepository repository.Repository) TrackService {
	return TrackService{
		rp: trackRepository,
	}
}

// Get a track by name
func (sv *TrackService) GetTrackByName(trackName string) (interface{}, error) {
	// Bussiness logic ...

	// Check if trackName is in the database
	trackList, maxScore, err := sv.rp.GetTracksElasticSearch("tracks", trackName)
	fmt.Println("MaxScore: ", maxScore)
	if errors.Is(err, internal.ErrTrackNotFound) || len(trackList) == 0 || maxScore < 3 { // If the track does not exist in the database use the API to get the track
		fmt.Println("No existe en la base de datos o es menor a 8")
		track, err := sv.rp.GetTrackByName(trackName)
		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest
			default:
				return nil, internal.ErrInternalServerError
			}
		}
		// Save the track in the database
		err = sv.rp.IndexTrack("tracks", track)

		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest
			case internal.ErrorIndexingDate:
				// Delete the index and try to index the track again
				sv.rp.DeleteIndex("tracks")
				err = sv.rp.IndexTrack("tracks", track)
				if err != nil {
					return nil, internal.ErrInternalServerError
				}
			default:
				return nil, internal.ErrInternalServerError
			}

		}

		return track, nil
	}

	return trackList, err
}

// Get a track by name and artist
func (sv *TrackService) GetTrackByArtistAndName(artistName string, trackName string) (interface{}, error) {
	// Bussiness logic ...

	// Check if trackName is in the database
	trackList, maxScore, err := sv.rp.GetTrackByArtistElasticSearch("tracks", artistName, trackName)
	fmt.Println("MaxScore: ", maxScore)
	if errors.Is(err, internal.ErrTrackNotFound) || maxScore < 3 { // If the track does not exist in the database use the API to get the track
		fmt.Println("No existe en la base de datos o es menor a 8")
		// Get the track from the Spotify API by artist and name
		track, err := sv.rp.GetTrackByArtist(trackName, artistName)
		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest
			default:
				return nil, internal.ErrInternalServerError
			}
		}
		// Save the track in the database
		err = sv.rp.IndexTrack("tracks", track)
		fmt.Println("El error es :", err)
		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest

			case internal.ErrorIndexingDate:
				// Delete the index and try to index the track again
				sv.rp.DeleteIndex("tracks")
				err = sv.rp.IndexTrack("tracks", track)
				if err != nil {
					return nil, internal.ErrInternalServerError
				}
			case internal.ErrCreatingIndex:
				return nil, internal.ErrCreatingIndex
			default:
				return nil, internal.ErrInternalServerError
			}

		}
		return track, nil
	}
	return trackList, nil
}

// Get all tracks
func (sv *TrackService) GetAllTracks() (interface{}, error) {
	// Bussiness logic ...

	// Get all tracks from the database
	trackList, err := sv.rp.GetAllTracksElasticSearch("tracks")
	if err != nil {
		switch err {
		case internal.ErrBadRequest:
			return nil, internal.ErrBadRequest
		case internal.ErrTrackNotFound:
			return nil, internal.ErrTrackNotFound
		default:
			return nil, internal.ErrInternalServerError
		}
	}
	return trackList, nil
}

// Get all tracks by album

func (sv *TrackService) GetAllTracksByAlbum(albumName string) (interface{}, error) {
	// Bussiness logic ...

	// Get all tracks from the database
	trackList, err := sv.rp.GetAllTracksByAlbumElasticSearch("tracks", albumName)

	// Check if there is an error
	if err != nil {
		switch err {
		case internal.ErrBadRequest:
			return nil, internal.ErrBadRequest
		case internal.ErrTrackNotFound:
			return nil, internal.ErrTrackNotFound
		default:
			return nil, internal.ErrInternalServerError
		}
	}
	// Unmarshall

	//fmt.Println("album", trackList["albums"].(string))
	if len(trackList) == 0 {
		// Use the API to get the tracks
		album, err := sv.rp.GetAllTracksByAlbum(albumName)
		//fmt.Println("El album es: ", album)
		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest
			default:
				return nil, internal.ErrInternalServerError
			}

		}

		err = sv.rp.IndexTrackByArtist("tracks", album)

		if err != nil {
			switch err {
			case internal.ErrBadRequest:
				return nil, internal.ErrBadRequest
			case internal.ErrorIndexingDate:

				return nil, internal.ErrInternalServerError

			default:
				return nil, internal.ErrInternalServerError
			}
		}

		return album, nil
	}
	return trackList, nil
}

// Get all tracks by popularity
func (sv *TrackService) GetTrackByPopularity(start int, end int) (interface{}, error) {
	// Bussiness logic ...

	album, err := sv.rp.GetAllTracksPopularityElasticSearch(start, end, "tracks")
	if err != nil {
		switch err {
		case internal.ErrBadRequest:
			return nil, internal.ErrBadRequest
		case internal.ErrTrackNotFound:
			return nil, internal.ErrTrackNotFound
		default:
			return nil, internal.ErrInternalServerError
		}

	}
	return album, nil
}
