# music_ms

Repository with the microservice music.

## Endpoints
### healthy
Check ms 
- **GET** `/`

### Tracks

#### Get Track By Artist and Name Track
Endpoint to retrieve a track by its artist and name.
- **GET** `/api/v1/music/tracks/artist?track=:track&artist=:artist`

#### Get Track By Name
Endpoint to retrieve a track by its name.
- **GET** `/api/v1/music?name=:name`

#### Get All Tracks
Endpoint to retrieve all tracks.
- **GET** `/api/v1/music/tracks`

#### Get All Tracks By Album
Endpoint to retrieve all tracks from a specific album.
- **GET** `/api/v1/music/album?name=:name`

#### Get All Tracks By Popularity
Endpoint to retrieve all tracks within a specific popularity range.
- **GET** `/api/v1/music/artist/popularity?start=:start&end=:end`

#### Get All Tracks By Release Date
Endpoint to retrieve all tracks released within a specific date range.
- **GET** `/api/v1/music/releasedate?start=:start_date&end=:end_date`

### Artists

#### Get All Artists
Endpoint to retrieve all artists.
- **GET** `/api/v1/music/artist`
Endpoint to get information about an artist 

- **GET** `/api/v1/music/artist?name=:name`
This README provides detailed endpoints for managing artists, albums, and tracks, along with search functionality with various query parameters for filtering results based on specific criteria.

## Commands to use elasticsearch

1. `colima start` 
2. `docker build -t dockerfile .`
3. docker-compose up`
4. `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}music_ms-elasticsearch-1 ``
5.  Probar en postman