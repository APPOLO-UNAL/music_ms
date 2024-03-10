package internal

// This script contains the logic to handle, service and repostory to  the artist entity

type Artist struct {
	Href          string   // The Spotify URL for the artist.
	Limit         int      // The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	External_urls []string // Known external URLs for this artist.
	Followers     int      // Information about the followers of the artist.
	Genres        []string // A list of the genres the artist is associated with. For example: "Prog Rock" , "Post-Grunge". (If not yet classified, the array is empty.)
	SpotifyID     string   // The Spotify ID for the artist.
	Name          string   // The name of the artist.

	Images     []Image // The cover art for the album in various sizes, widest first.
	Popularity int     // The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular.
	SpotifyURI string  // The Spotify URI for the artist.
}

// ArtistRepository is the interface that provides artist methods
type ArtistRepository interface {
	// Get Methods
	GetArtist(id string) (Artist, error)                    // Get an artist
	GetArtistByName(name string) (Artist, error)            // Get an artist by name
	GetArtistByGenre(genre string) ([]Artist, error)        // Get artists by genre
	GetArtistByPopularity(popularity int) ([]Artist, error) // Get artists by popularity
	GetArtistByFollowers(followers int) ([]Artist, error)   // Get artists by followers
	GetArtistByCountry(country string) ([]Artist, error)    // Get artists by country

	// Create Methods
	CreateArtist(artist Artist) (Artist, error)

	// Update Methods
	UpdateArtist(artist Artist) (Artist, error)

	// Delete Methods
	DeleteArtist(id string) error
}

// ArtistService is the interface that provides artist methods

type ArtistService interface {
	// Get Methods
	GetArtist(id string) (Artist, error)                    // Get an artist
	GetArtistByName(name string) (Artist, error)            // Get an artist by name
	GetArtistByGenre(genre string) ([]Artist, error)        // Get artists by genre
	GetArtistByPopularity(popularity int) ([]Artist, error) // Get artists by popularity
	GetArtistByFollowers(followers int) ([]Artist, error)   // Get artists by followers
	GetArtistByCountry(country string) ([]Artist, error)    // Get artists by country

	// Create Methods
	CreateArtist(artist Artist) (Artist, error)

	// Update Methods
	UpdateArtist(artist Artist) (Artist, error)

	// Delete Methods
	DeleteArtist(id string) error
}
