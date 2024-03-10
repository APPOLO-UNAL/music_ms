package internal

// This script contains the logic to handle, service and repostory to  the album entity

type Album struct {
	Href                 string   // The Spotify URL for the album.
	AlbumType            string   // The type of the album: one of "album" , "single" , or "compilation".
	TotalTracks          int      // The total number of tracks in the album. In case of an album takedown, the value may be null.
	AvailableMarkets     []string // The markets in which the album is available: ISO 3166-1 alpha-2 country codes.
	SpotifyURL           string   // The Spotify URL for the album.
	SpotifyID            string   // The Spotify ID for the album.
	Images               []Image  // The cover art for the album in various sizes, widest first.
	Name                 string   // The name of the album. In case of an album takedown, the value may be null.
	ReleaseDate          string   // The date the album was first released, for example "1981-12-15". Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleaseDatePrecision string   // The precision with which release_date value is known: "year" , "month" , or "day".

}
