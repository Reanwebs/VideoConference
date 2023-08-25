package interfaces

import "conference/pkg/common/utility"

type PrivateRepo interface {
	CreatePrivateRoom(utility.PrivateRoom) (uint, error)
	CheckPrivateLimit(string) (uint, error)
	CountPrivateParticipants(string) (uint, error)
	CheckPrivateParticipantPermission(string, string) (bool, error)
	AddParticipantInPrivateRoom(utility.PrivateRoomParticipants) error
	BlockPrivateParticipant(string, string) error
	RemovePrivateParticipant(string, string) error
	UpdatePrivateParticipantExitTime(utility.PrivateRoomParticipants) error
	CheckType(string) (string, error)
	CheckPrivateInterest(string) (string, error)
	RemovePrivateRoom(string) error
}

type GroupRepo interface {
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
