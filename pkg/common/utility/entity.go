package utility

import (
	"time"

	"gorm.io/gorm"
)

type PrivateRoom struct {
	gorm.Model
	UserID           string
	ConferenceID     string
	Type             string `gorm:"column:type;default:'private'"`
	Title            string
	Description      string
	Interest         string
	Recording        bool
	Chat             bool
	Broadcast        bool
	Participantlimit uint
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type PrivateRoomParticipants struct {
	gorm.Model
	UserID       string
	ConferenceID string
	Permission   bool
	CamStatus    string
	MicStatus    string
	JoinTime     time.Time
	ExitTime     time.Time
	Role         string
}

type GroupRoom struct {
	gorm.Model
	UserID           string
	ConferenceID     string
	GroupID          string
	Type             string `gorm:"column:type;default:'group'"`
	Title            string
	Description      string
	Interest         string
	Recording        bool
	Chat             bool
	Broadcast        bool
	Participantlimit uint
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type GroupRoomParticipants struct {
	gorm.Model
	UserID       string
	ConferenceID string
	GroupID      string
	Permission   bool
	CamStatus    string
	MicStatus    string
	JoinTime     time.Time
	ExitTime     time.Time
	Role         string
}

type PublicRoom struct {
	gorm.Model
	UserID           string
	ConferenceID     string
	Type             string `gorm:"column:type;default:'public'"`
	Title            string
	Description      string
	Interest         string
	Recording        bool
	Chat             bool
	Broadcast        bool
	Participantlimit uint
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type PublicRoomParticipants struct {
	gorm.Model
	UserID       string
	ConferenceID string
	GroupID      string
	Permission   bool
	CamStatus    string
	MicStatus    string
	JoinTime     time.Time
	ExitTime     time.Time
	Role         string
}
