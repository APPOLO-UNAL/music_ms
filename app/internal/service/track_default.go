package service

import (
	"errors"
	"fmt"
	"ms_music/app/internal"
	repository "ms_music/app/internal/repository"
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
	if errors.Is(err, internal.ErrTrackNotFound) || len(trackList) == 0 || maxScore < 8 { // If the track does not exist in the database use the API to get the track
		fmt.Println("No existe en la base de datos o es menor a 8")
		track, err := sv.rp.GetTrackByName(trackName)
		fmt.Println("error1 ", err)
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
		fmt.Println("error2 ", err)

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

		// Add the track to trackList
		trackList = append(trackList, track)
	}

	return trackList, err
}
