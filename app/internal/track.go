package internal

// This script contains the logic to handle, service and repostory to  the track entity
type Track struct {
	Href             string   // The Spotify URL for the track.
	Limit            int      // The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Album            Album    // The album on which the track appears. The album object includes a link in href to full information about the album.
	Artists          []Artist // The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	AvailableMarkets []string // The markets in which the track is available: ISO 3166-1 alpha-2 country codes.
	DiscNumber       int      // The disc number (usually 1 unless the album consists of more than one disc).
	DurationMs       int      // The track length in milliseconds.
	Explicit         bool     // Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	SpotifyURL       string   // The Spotify URL for the track.
	SpotifyID        string   // The Spotify ID for the track.
	Name             string   // The name of the track.
	Popularity       int      // The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	PreviewURL       string   // A link to a 30 second preview (MP3 format) of the track. Can be null.
	TrackNumber      int      // The number of the track. If an album has several discs, the track number is the number on the specified disc.

}
