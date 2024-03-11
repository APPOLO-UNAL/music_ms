package service

import (
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
func (rp *TrackService) GetTrackByName(trackName string) (interface{}, error) {
	// Bussiness logic ...

	// Check if trackName is in the database
	// ...
	// If not use the API to get the track
	track, err := rp.rp.GetTrackByName(trackName)

	return track, err
}
