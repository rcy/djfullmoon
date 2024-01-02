// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	EventID   string
	EventType string
	CreatedAt pgtype.Timestamptz
	Payload   []byte
}

type Result struct {
	ResultID  string
	SearchID  string
	StationID string
	CreatedAt pgtype.Timestamptz
	ExternID  string
	Url       string
	Thumbnail string
	Title     string
	Uploader  string
	Duration  float64
	Views     float64
}

type SchemaVersion struct {
	Version int32
}

type Search struct {
	SearchID  string
	StationID string
	CreatedAt pgtype.Timestamptz
	Query     string
	Status    string
}

type Station struct {
	StationID string
	CreatedAt pgtype.Timestamptz
	Slug      string
	Name      string
	Active    bool
}

type StationMessage struct {
	StationMessageID string
	CreatedAt        pgtype.Timestamptz
	Type             string
	StationID        string
	ParentID         string
	Nick             string
	Body             string
}

type Track struct {
	TrackID     string
	StationID   string
	CreatedAt   pgtype.Timestamptz
	Artist      string
	Title       string
	RawMetadata []byte
	Rotation    int32
	Plays       int32
	Skips       int32
	Playing     bool
}
