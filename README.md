# music_ms

Repository with the microservice music.

## Endpoints

### Artist

#### Get
Endpoint to retrieve information about an artist.
- **GET** `/artists/:id`

#### Create
Endpoint to create a new artist.
- **POST** `/artists`

#### Update
Endpoint to update information of an existing artist.
- **PUT** `/artists/:id`

#### Delete
Endpoint to delete an existing artist.
- **DELETE** `/artists/:id`

### Search

#### Search by Genre
Endpoint to search for artists based on genre.
- **GET** `/artists/search?genre=:genre`

#### Search by Name
Endpoint to search for artists based on name.
- **GET** `/artists/search?name=:name`

#### Search with Limit
Endpoint to search for artists with a specified limit on the number of results.
- **GET** `/artists/search?limit=:limit`

#### Search by Followers
Endpoint to search for artists based on the number of followers.
- **GET** `/artists/search?followers=:followers`

#### Search by Popularity
Endpoint to search for artists based on popularity.
- **GET** `/artists/search?popularity=:popularity`

## Albums

#### Get
Endpoint to retrieve information about an album.
- **GET** `/albums/:id`

#### Create
Endpoint to create a new album.
- **POST** `/albums`

#### Update
Endpoint to update information of an existing album.
- **PUT** `/albums/:id`

#### Delete
Endpoint to delete an existing album.
- **DELETE** `/albums/:id`

#### Search
Endpoint to search for albums based on specified criteria.
- **GET** `/albums/search`

Possible Query Parameters:
- `q`: Search query to filter albums by name, artist, or other relevant information.
- `limit`: The maximum number of albums to be returned in the response.
- `genre`: Filter albums by genre.
- `release_date`: Filter albums by release date or range of release dates.
- `artist`: Filter albums by artist name or ID.
- `min_tracks`: Filter albums by the minimum number of tracks.
- `max_tracks`: Filter albums by the maximum number of tracks.
- `min_popularity`: Filter albums by minimum popularity score.
- `max_popularity`: Filter albums by maximum popularity score.
- `available_markets`: Filter albums by available markets (ISO 3166-1 alpha-2 country codes).

### Album

#### Get
Endpoint to retrieve information about an album.
- **GET** `/albums/:id`

#### Create
Endpoint to create a new album.
- **POST** `/albums`

#### Update
Endpoint to update information of an existing album.
- **PUT** `/albums/:id`

#### Delete
Endpoint to delete an existing album.
- **DELETE** `/albums/:id`

#### Search
Endpoint to search for albums based on specified criteria.
- **GET** `/albums/search?q=:query`

## Tracks

#### Get
Endpoint to retrieve information about a track.
- **GET** `/tracks/:id`

#### Create
Endpoint to create a new track.
- **POST** `/tracks`

#### Update
Endpoint to update information of an existing track.
- **PUT** `/tracks/:id`

#### Delete
Endpoint to delete an existing track.
- **DELETE** `/tracks/:id`

#### Search
Endpoint to search for tracks based on specified criteria.
- **GET** `/tracks/search?q=:query`

Possible Query Parameters:
- `q`: Search query to filter tracks by name, artist, album, or other relevant information.
- `limit`: The maximum number of tracks to be returned in the response.
- `artist`: Filter tracks by artist name or ID.
- `album`: Filter tracks by album name or ID.
- `genre`: Filter tracks by genre.
- `release_date`: Filter tracks by release date or range of release dates.
- `min_duration`: Filter tracks by minimum duration (in milliseconds).
- `max_duration`: Filter tracks by maximum duration (in milliseconds).
- `min_popularity`: Filter tracks by minimum popularity score.
- `max_popularity`: Filter tracks by maximum popularity score.
- `available_markets`: Filter tracks by available markets (ISO 3166-1 alpha-2 country codes).

This README provides detailed endpoints for managing artists, albums, and tracks, along with search functionality with various query parameters for filtering results based on specific criteria.

## Commands to use elasticsearch

1. Create ElasticSearch Image
`docker pull docker.elastic.co/elasticsearch/elasticsearch:7.5.2
`

2. Create a Docker Volume `docker volume create elasticsearch
`

3. Create a new command to run ElasticSearch 

```bash
#! /bin/bash

docker rm -f elasticsearch
docker run -d --name elasticsearch -p 9200:9200 -e discovery.type=single-node \
    -v elasticsearch:/usr/share/elasticsearch/data \
    docker.elastic.co/elasticsearch/elasticsearch:7.5.2
docker ps
```


4. Change executbable 
5. 