package utility

import (
	"time"

	"gorm.io/gorm"
)

type ConferenceRoom struct {
	gorm.Model
	UserID           string
	Type             string
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

type ConferenceParticipants struct {
	gorm.Model
	UserID       string
	ConferenceID uint
	Permission   bool
	CamStatus    string
	MicStatus    string
	JoinTime     time.Time
	ExitTime     time.Time
	Role         string
}
