package internal

// TrackRepository is the interface that provides track methods
type TrackRepository interface {
	// Configuration Methods
	VerifyIndices() error // Verify if the indices exist, if not, create them

	// Get Methods
	GetTrack(id string) (Track, error)                                     // Get a track
	GetTrackByArtist(artist string) ([]Track, error)                       // Get tracks by artist
	GetTrackByAlbum(album string) ([]Track, error)                         // Get tracks by album
	GetTrackByGenre(genre string) ([]Track, error)                         // Get tracks by genre
	GetTrackByPopularity(popularity int) ([]Track, error)                  // Get tracks by popularity
	GetTrackByDuration(duration int) ([]Track, error)                      // Get tracks by duration
	GetTrackByReleaseDate(releaseDate string) ([]Track, error)             // Get tracks by release date
	GetTrackBetweenDuration(duration1 int, duration2 int) ([]Track, error) // Get tracks between duration
	GetTrackAvailableInMarket(market string) ([]Track, error)              // Get tracks available in market

	// Create Methods
	CreateTrack(track Track) (Track, error)

	// Update Methods
	UpdateTrack(track Track) (Track, error)

	// Delete Methods
	DeleteTrack(id string) error
}
