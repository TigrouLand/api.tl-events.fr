package entities

import (
	"github.com/google/uuid"
	"github.com/tigrouland/api/mongo"
)

type ArchivedGame struct {
	ID   int8   `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`

	Type   string `bson:"_t" json:"type"`
	Status string `bson:"status" json:"status"`

	Alive     []mongo.DBRef `bson:"alive" json:"-"`
	AliveUUID []uuid.UUID   `json:"alive"`

	Players     []mongo.DBRef `bson:"players" json:"-"`
	PlayersUUID []uuid.UUID   `json:"players"`

	Moles     []mongo.DBRef `bson:"baseMoles" json:"-"`
	MolesUUID []uuid.UUID   `json:"moles,omitempty"`

	Modifiers []string `bson:"modifiers" json:"modifiers"`
	Events    []string `bson:"events" json:"events,omitempty"`

	Logs []string `bson:"logs" json:"logs"`

	TeamRefs []mongo.DBRef `bson:"baseTeams" json:"-"`
	Teams    []Team        `json:"teams,omitempty"`

	PlayerTeams map[string]string `bson:"playerTeams" json:"playerTeams,omitempty"`
	FinalTeams  map[string]string `bson:"finalTeams" json:"finalTeams,omitempty"`
	PlayerRoles map[string]string `bson:"playerRoles" json:"playerRoles,omitempty"`

	Hours   int8 `bson:"hours" json:"hours"`
	Minutes int8 `bson:"minutes" json:"minutes"`
	Seconds int8 `bson:"seconds" json:"seconds"`

	StartDate    int64 `bson:"startDate" json:"startDate"`
	ScheduleDate int64 `bson:"scheduleDate" json:"scheduleDate"`
	ArchiveDate  int64 `bson:"archiveDate" json:"archiveDate"`
}

type Game struct {
	ID   int8   `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`

	Type   string `bson:"type" json:"type"`
	Status string `bson:"status" json:"status"`

	// Alive     []primitive.Binary `bson:"alive" json:"-"`
	// AliveUUID []uuid.UUID        `json:"alive"`

	Players     []mongo.DBRef `bson:"players" json:"-"`
	PlayersUUID []uuid.UUID   `json:"players"`

	// Moles []mongo.DBRef `bson:"baseMoles" json:"-"`
	// MolesUUID []uuid.UUID `json:"moles,omitempty"`

	Modifiers []string `bson:"modifiers" json:"modifiers"`
	// Events    []string `bson:"events" json:"events,omitempty"`

	// Logs []string `bson:"logs" json:"logs"`

	TeamRefs []mongo.DBRef `bson:"baseTeams" json:"-"`
	Teams    []Team        `json:"teams,omitempty"`

	// PlayerTeams map[string]string `bson:"playerTeams" json:"playerTeams,omitempty"`
	// PlayerRoles map[string]string `bson:"playerRoles" json:"playerRoles,omitempty"`

	Hours   int8 `bson:"hours" json:"hours"`
	Minutes int8 `bson:"minutes" json:"minutes"`
	Seconds int8 `bson:"seconds" json:"seconds"`

	StartDate    int64 `bson:"startDate" json:"startDate"`
	ScheduleDate int64 `bson:"scheduleDate" json:"scheduleDate"`
	ArchiveDate  int64 `bson:"archiveDate" json:"archiveDate"`
}
