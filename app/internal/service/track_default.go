package service

import (
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
	// ...
	// If not use the API to get the track
	track, err := sv.rp.GetTrackByName(trackName)
	if err != nil {
		switch err {
		case internal.ErrBadRequest:
			return nil, internal.ErrBadRequest
		default:
			return nil, internal.ErrInternalServerError
		}
	}

	// Index the track in the database

	err = sv.rp.IndexTrack("tracks", track)
	fmt.Println("IndexTrack", err)
	if err != nil {
		switch err {
		case internal.ErrBadRequest:
			return nil, internal.ErrBadRequest
		default:
			return nil, internal.ErrInternalServerError
		}
	}
	return track, err
}
