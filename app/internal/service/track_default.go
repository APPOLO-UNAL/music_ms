package service

import (
	repository "ms_music/app/internal/repository"
)

// This script contains the logic to handle, service and repostory to  the track entity
type TrackService struct {
	TrackRepository repository.TrackRepository
}

// NewTrackService returns a new TrackService
func NewTrackService(trackRepository repository.TrackRepository) TrackService {
	return TrackService{
		TrackRepository: trackRepository,
	}
}
