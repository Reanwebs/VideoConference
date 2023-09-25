package interfaces

import (
	"conference/pkg/common/utility"
	"time"
)

type PrivateRepo interface {
	CreatePrivateSchedule(utility.ScheduleConference) (uint, error)
	GetPrivateSchedules(string) ([]utility.ScheduleConference, error)
	GetCompletedSchedules(string) ([]utility.ScheduleConference, error)
	CreatePrivateRoom(utility.PrivateRoom) (uint, error)
	CheckPrivateLimit(string) (uint, error)
	CountPrivateParticipants(string) (uint, error)
	CheckPrivateParticipantPermission(string, string) (bool, error)
	GetSdpOffer(string) (string, error)
	AddParticipantInPrivateRoom(utility.PrivateRoomParticipants) error
	BlockPrivateParticipant(string, string) error
	RemovePrivateParticipant(string, string) error
	GetJoinTime(string, string) (time.Time, error)
	UpdatePrivateParticipantExitTime(utility.PrivateRoomParticipants) error
	CheckType(string) (string, error)
	CheckPrivateInterest(string) (string, error)
	RemovePrivateRoom(string) error
}

type GroupRepo interface {
	CreateGroupSchedule(utility.ScheduleGroupConference) (uint, error)
	CreateGroupRoom(utility.GroupRoom) error
	AddParticipantInGroupRoom(utility.GroupRoomParticipants) error
	CheckGroupLimit(string) (uint, error)
	CheckGroupParticipantPermission(string, string) (bool, error)
	CountGroupParticipants(string) (uint, error)
	UpdateGroupParticipantExitTime(utility.GroupRoomParticipants) error
	RemoveGroupParticipant(string, string) error
	BlockGroupParticipant(string, string) error
	RemoveGroupRoom(string) error
}

type PublicRepo interface {
	CreateStreamRoom(utility.StreamRoom) error
	GetStream(string) (utility.StreamRoom, error)
	GetStreamList() ([]utility.StreamRoom, error)
	GetSortedStreamList(string) ([]utility.StreamRoom, error)
	UpdateStreamRoom(string, string, string) error
	FindStream(string) error
	AddStreamParticipants(utility.StreamRoomParticipants) error
	UpdateStreamParticipants(utility.StreamRoomParticipants) error
	GetStreamJoinTime(string, string) (time.Time, error)
	CreatePublicSchedule(utility.SchedulePublicConference) (uint, error)
	CreatePublicRoom(utility.PublicRoom) error
	AddParticipantInPublicRoom(utility.PublicRoomParticipants) error
	CheckPublicLimit(string) (uint, error)
	CountPublicParticipants(string) (uint, error)
	CheckPublicParticipantPermission(string, string) (bool, error)
	UpdatePublicParticipantExitTime(utility.PublicRoomParticipants) error
	RemovePublicParticipant(string, string) error
	BlockPublicParticipant(string, string) error
	RemovePublicRoom(string) error
}
