package utility

import (
	"time"

	"gorm.io/gorm"
)

type PrivateRoom struct {
	gorm.Model
	UserID           string `gorm:"not null"`
	SdpOffer         string
	IceCandidate     string
	ConferenceID     string `gorm:"unique;not null"`
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
	SdpAnswer    string
	IceCandidate string
	Permission   bool
	CamStatus    string
	MicStatus    string
	JoinTime     time.Time
	ExitTime     time.Time
	Role         string
}

type StreamRoom struct {
	gorm.Model
	HostID      string
	StreamID    string
	Title       string
	Description string
	ThumbnailID string
	Interest    string
	Status      string
	AvatarID    string
	UserName    string
}

type StreamRoomParticipants struct {
	gorm.Model
	StreamID      string
	ParticipantID string
	JoinTime      time.Time
	LeaveTime     time.Time
}

type GroupRoom struct {
	gorm.Model
	UserID           string
	ConferenceID     string `gorm:"unique;not null"`
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
	ConferenceID     string `gorm:"unique;not null"`
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

type ScheduleConference struct {
	gorm.Model
	UserID           string
	ScheduleID       string
	Title            string
	Description      string
	Interest         string
	ParticipantLimit uint
	Time             time.Time
	Duration         time.Duration
	Status           string
}

type ScheduleGroupConference struct {
	gorm.Model
	GroupID     string
	ScheduleID  string
	Title       string
	Description string
	Interest    string
	Month       time.Month
	Day         time.Weekday
	Time        time.Time
	Duration    time.Duration
}

type SchedulePublicConference struct {
	gorm.Model
	UserID      string
	ScheduleID  string
	Title       string
	Description string
	Interest    string
	Month       time.Month
	Day         time.Weekday
	Time        time.Time
	Duration    time.Duration
}

type ScheduleEmail struct {
	Subject     string
	Content     string
	To          []string
	Cc          []string
	Bcc         []string
	AttachFiles []string
}
