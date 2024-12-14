// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const activeStations = `-- name: ActiveStations :many
select station_id, created_at, slug, name, active, current_track_id, background_image_url, user_id from stations where active = true
`

func (q *Queries) ActiveStations(ctx context.Context) ([]Station, error) {
	rows, err := q.db.Query(ctx, activeStations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Station
	for rows.Next() {
		var i Station
		if err := rows.Scan(
			&i.StationID,
			&i.CreatedAt,
			&i.Slug,
			&i.Name,
			&i.Active,
			&i.CurrentTrackID,
			&i.BackgroundImageURL,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createGuestUser = `-- name: CreateGuestUser :one
insert into users(guest, user_id) values(true, $1) returning user_id, created_at, username, guest
`

func (q *Queries) CreateGuestUser(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRow(ctx, createGuestUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Username,
		&i.Guest,
	)
	return i, err
}

const createResult = `-- name: CreateResult :exec
insert into results(result_id, search_id, station_id, extern_id, url, thumbnail, title, uploader, duration, views) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
`

type CreateResultParams struct {
	ResultID  string
	SearchID  string
	StationID string
	ExternID  string
	URL       string
	Thumbnail string
	Title     string
	Uploader  string
	Duration  float64
	Views     float64
}

func (q *Queries) CreateResult(ctx context.Context, arg CreateResultParams) error {
	_, err := q.db.Exec(ctx, createResult,
		arg.ResultID,
		arg.SearchID,
		arg.StationID,
		arg.ExternID,
		arg.URL,
		arg.Thumbnail,
		arg.Title,
		arg.Uploader,
		arg.Duration,
		arg.Views,
	)
	return err
}

const createSearch = `-- name: CreateSearch :exec
insert into searches(search_id, station_id, query) values($1,$2,$3)
`

type CreateSearchParams struct {
	SearchID  string
	StationID string
	Query     string
}

func (q *Queries) CreateSearch(ctx context.Context, arg CreateSearchParams) error {
	_, err := q.db.Exec(ctx, createSearch, arg.SearchID, arg.StationID, arg.Query)
	return err
}

const createSession = `-- name: CreateSession :one
insert into sessions(session_id, user_id) values($1, $2) returning session_id
`

func (q *Queries) CreateSession(ctx context.Context, sessionID string, userID string) (string, error) {
	row := q.db.QueryRow(ctx, createSession, sessionID, userID)
	var session_id string
	err := row.Scan(&session_id)
	return session_id, err
}

const createStation = `-- name: CreateStation :one
insert into stations(station_id, slug, user_id, active) values($1, $2, $3, $4) returning station_id, created_at, slug, name, active, current_track_id, background_image_url, user_id
`

type CreateStationParams struct {
	StationID string
	Slug      string
	UserID    string
	Active    bool
}

func (q *Queries) CreateStation(ctx context.Context, arg CreateStationParams) (Station, error) {
	row := q.db.QueryRow(ctx, createStation,
		arg.StationID,
		arg.Slug,
		arg.UserID,
		arg.Active,
	)
	var i Station
	err := row.Scan(
		&i.StationID,
		&i.CreatedAt,
		&i.Slug,
		&i.Name,
		&i.Active,
		&i.CurrentTrackID,
		&i.BackgroundImageURL,
		&i.UserID,
	)
	return i, err
}

const createStationMessage = `-- name: CreateStationMessage :one
insert into station_messages(station_message_id, type, station_id, nick, body, parent_id) values($1, $2, $3, $4, $5, $6) returning station_message_id, created_at, type, station_id, parent_id, nick, body
`

type CreateStationMessageParams struct {
	StationMessageID string
	Type             string
	StationID        string
	Nick             string
	Body             string
	ParentID         string
}

func (q *Queries) CreateStationMessage(ctx context.Context, arg CreateStationMessageParams) (StationMessage, error) {
	row := q.db.QueryRow(ctx, createStationMessage,
		arg.StationMessageID,
		arg.Type,
		arg.StationID,
		arg.Nick,
		arg.Body,
		arg.ParentID,
	)
	var i StationMessage
	err := row.Scan(
		&i.StationMessageID,
		&i.CreatedAt,
		&i.Type,
		&i.StationID,
		&i.ParentID,
		&i.Nick,
		&i.Body,
	)
	return i, err
}

const createTrack = `-- name: CreateTrack :one
insert into tracks(track_id, station_id, artist, title, raw_metadata, rotation)
values($1,$2,$3,$4,$5, (coalesce((select min(rotation) from tracks where station_id = $2), 0)))
returning track_id, station_id, created_at, artist, title, raw_metadata, rotation, plays, skips, playing
`

type CreateTrackParams struct {
	TrackID     string
	StationID   string
	Artist      string
	Title       string
	RawMetadata []byte
}

func (q *Queries) CreateTrack(ctx context.Context, arg CreateTrackParams) (Track, error) {
	row := q.db.QueryRow(ctx, createTrack,
		arg.TrackID,
		arg.StationID,
		arg.Artist,
		arg.Title,
		arg.RawMetadata,
	)
	var i Track
	err := row.Scan(
		&i.TrackID,
		&i.StationID,
		&i.CreatedAt,
		&i.Artist,
		&i.Title,
		&i.RawMetadata,
		&i.Rotation,
		&i.Plays,
		&i.Skips,
		&i.Playing,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
insert into users(user_id, username) values($1, $2) returning user_id, created_at, username, guest
`

func (q *Queries) CreateUser(ctx context.Context, userID string, username string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, userID, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Username,
		&i.Guest,
	)
	return i, err
}

const event = `-- name: Event :one
select event_id, event_type, created_at, payload from events where event_id = $1
`

func (q *Queries) Event(ctx context.Context, eventID string) (Event, error) {
	row := q.db.QueryRow(ctx, event, eventID)
	var i Event
	err := row.Scan(
		&i.EventID,
		&i.EventType,
		&i.CreatedAt,
		&i.Payload,
	)
	return i, err
}

const incrementTrackPlays = `-- name: IncrementTrackPlays :exec
update tracks set plays = plays + 1 where track_id = $1
`

func (q *Queries) IncrementTrackPlays(ctx context.Context, trackID string) error {
	_, err := q.db.Exec(ctx, incrementTrackPlays, trackID)
	return err
}

const incrementTrackRotation = `-- name: IncrementTrackRotation :exec
update tracks set rotation = rotation + 1 where track_id = $1
`

func (q *Queries) IncrementTrackRotation(ctx context.Context, trackID string) error {
	_, err := q.db.Exec(ctx, incrementTrackRotation, trackID)
	return err
}

const insertEvent = `-- name: InsertEvent :one
insert into events(event_id, event_type, payload) values ($1, $2, $3) returning event_id, event_type, created_at, payload
`

type InsertEventParams struct {
	EventID   string
	EventType string
	Payload   []byte
}

func (q *Queries) InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, insertEvent, arg.EventID, arg.EventType, arg.Payload)
	var i Event
	err := row.Scan(
		&i.EventID,
		&i.EventType,
		&i.CreatedAt,
		&i.Payload,
	)
	return i, err
}

const oldestUnplayedTrack = `-- name: OldestUnplayedTrack :one
select track_id, station_id, created_at, artist, title, raw_metadata, rotation, plays, skips, playing from tracks
where tracks.station_id = $1
and plays = 0
and rotation = (select min(rotation) from tracks where station_id = $1)
order by track_id asc
limit 1
`

func (q *Queries) OldestUnplayedTrack(ctx context.Context, stationID string) (Track, error) {
	row := q.db.QueryRow(ctx, oldestUnplayedTrack, stationID)
	var i Track
	err := row.Scan(
		&i.TrackID,
		&i.StationID,
		&i.CreatedAt,
		&i.Artist,
		&i.Title,
		&i.RawMetadata,
		&i.Rotation,
		&i.Plays,
		&i.Skips,
		&i.Playing,
	)
	return i, err
}

const randomTrack = `-- name: RandomTrack :one
select track_id, station_id, created_at, artist, title, raw_metadata, rotation, plays, skips, playing from tracks
where tracks.station_id = $1
and plays > 0
and rotation = (select min(rotation) from tracks where station_id = $1)
order by random()
limit 1
`

func (q *Queries) RandomTrack(ctx context.Context, stationID string) (Track, error) {
	row := q.db.QueryRow(ctx, randomTrack, stationID)
	var i Track
	err := row.Scan(
		&i.TrackID,
		&i.StationID,
		&i.CreatedAt,
		&i.Artist,
		&i.Title,
		&i.RawMetadata,
		&i.Rotation,
		&i.Plays,
		&i.Skips,
		&i.Playing,
	)
	return i, err
}

const results = `-- name: Results :many
select result_id, search_id, station_id, created_at, extern_id, url, thumbnail, title, uploader, duration, views from results where search_id = $1
`

func (q *Queries) Results(ctx context.Context, searchID string) ([]Result, error) {
	rows, err := q.db.Query(ctx, results, searchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Result
	for rows.Next() {
		var i Result
		if err := rows.Scan(
			&i.ResultID,
			&i.SearchID,
			&i.StationID,
			&i.CreatedAt,
			&i.ExternID,
			&i.URL,
			&i.Thumbnail,
			&i.Title,
			&i.Uploader,
			&i.Duration,
			&i.Views,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const search = `-- name: Search :one
select search_id, station_id, created_at, query, status from searches where search_id = $1
`

func (q *Queries) Search(ctx context.Context, searchID string) (Search, error) {
	row := q.db.QueryRow(ctx, search, searchID)
	var i Search
	err := row.Scan(
		&i.SearchID,
		&i.StationID,
		&i.CreatedAt,
		&i.Query,
		&i.Status,
	)
	return i, err
}

const sessionUser = `-- name: SessionUser :one
select users.user_id, users.created_at, users.username, users.guest from sessions
join users on users.user_id = sessions.user_id
where sessions.expires_at > now()
and session_id = $1
`

func (q *Queries) SessionUser(ctx context.Context, sessionID string) (User, error) {
	row := q.db.QueryRow(ctx, sessionUser, sessionID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Username,
		&i.Guest,
	)
	return i, err
}

const setSearchStatusCompleted = `-- name: SetSearchStatusCompleted :exec
update searches set status = 'completed' where search_id = $1
`

func (q *Queries) SetSearchStatusCompleted(ctx context.Context, searchID string) error {
	_, err := q.db.Exec(ctx, setSearchStatusCompleted, searchID)
	return err
}

const setSearchStatusFailed = `-- name: SetSearchStatusFailed :exec
update searches set status = 'failed' where search_id = $1
`

func (q *Queries) SetSearchStatusFailed(ctx context.Context, searchID string) error {
	_, err := q.db.Exec(ctx, setSearchStatusFailed, searchID)
	return err
}

const setStationCurrentTrack = `-- name: SetStationCurrentTrack :exec
update stations set current_track_id = $1 where station_id = $2
`

func (q *Queries) SetStationCurrentTrack(ctx context.Context, currentTrackID pgtype.Text, stationID string) error {
	_, err := q.db.Exec(ctx, setStationCurrentTrack, currentTrackID, stationID)
	return err
}

const station = `-- name: Station :one
select station_id, created_at, slug, name, active, current_track_id, background_image_url, user_id from stations where slug = $1
`

func (q *Queries) Station(ctx context.Context, slug string) (Station, error) {
	row := q.db.QueryRow(ctx, station, slug)
	var i Station
	err := row.Scan(
		&i.StationID,
		&i.CreatedAt,
		&i.Slug,
		&i.Name,
		&i.Active,
		&i.CurrentTrackID,
		&i.BackgroundImageURL,
		&i.UserID,
	)
	return i, err
}

const stationCurrentTrack = `-- name: StationCurrentTrack :one
select tracks.track_id, tracks.station_id, tracks.created_at, tracks.artist, tracks.title, tracks.raw_metadata, tracks.rotation, tracks.plays, tracks.skips, tracks.playing from stations join tracks on stations.current_track_id = tracks.track_id where stations.station_id = $1
`

func (q *Queries) StationCurrentTrack(ctx context.Context, stationID string) (Track, error) {
	row := q.db.QueryRow(ctx, stationCurrentTrack, stationID)
	var i Track
	err := row.Scan(
		&i.TrackID,
		&i.StationID,
		&i.CreatedAt,
		&i.Artist,
		&i.Title,
		&i.RawMetadata,
		&i.Rotation,
		&i.Plays,
		&i.Skips,
		&i.Playing,
	)
	return i, err
}

const stationMessages = `-- name: StationMessages :many
select station_message_id, created_at, type, station_id, parent_id, nick, body from station_messages where station_id = $1 order by station_message_id desc limit 500
`

func (q *Queries) StationMessages(ctx context.Context, stationID string) ([]StationMessage, error) {
	rows, err := q.db.Query(ctx, stationMessages, stationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StationMessage
	for rows.Next() {
		var i StationMessage
		if err := rows.Scan(
			&i.StationMessageID,
			&i.CreatedAt,
			&i.Type,
			&i.StationID,
			&i.ParentID,
			&i.Nick,
			&i.Body,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const track = `-- name: Track :one
select track_id, station_id, created_at, artist, title, raw_metadata, rotation, plays, skips, playing from tracks where track_id = $1
`

func (q *Queries) Track(ctx context.Context, trackID string) (Track, error) {
	row := q.db.QueryRow(ctx, track, trackID)
	var i Track
	err := row.Scan(
		&i.TrackID,
		&i.StationID,
		&i.CreatedAt,
		&i.Artist,
		&i.Title,
		&i.RawMetadata,
		&i.Rotation,
		&i.Plays,
		&i.Skips,
		&i.Playing,
	)
	return i, err
}

const trackRequestStationMessage = `-- name: TrackRequestStationMessage :one
select station_message_id, created_at, type, station_id, parent_id, nick, body from station_messages
where station_id = $1
and type = 'TrackRequested'
and parent_id = $2
`

func (q *Queries) TrackRequestStationMessage(ctx context.Context, stationID string, parentID string) (StationMessage, error) {
	row := q.db.QueryRow(ctx, trackRequestStationMessage, stationID, parentID)
	var i StationMessage
	err := row.Scan(
		&i.StationMessageID,
		&i.CreatedAt,
		&i.Type,
		&i.StationID,
		&i.ParentID,
		&i.Nick,
		&i.Body,
	)
	return i, err
}

const updateStationMessage = `-- name: UpdateStationMessage :exec
update station_messages set type = $1, body = $2 where station_message_id = $3
`

type UpdateStationMessageParams struct {
	Type             string
	Body             string
	StationMessageID string
}

func (q *Queries) UpdateStationMessage(ctx context.Context, arg UpdateStationMessageParams) error {
	_, err := q.db.Exec(ctx, updateStationMessage, arg.Type, arg.Body, arg.StationMessageID)
	return err
}

const user = `-- name: User :one
select user_id, created_at, username, guest from users where user_id = $1
`

func (q *Queries) User(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRow(ctx, user, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Username,
		&i.Guest,
	)
	return i, err
}
