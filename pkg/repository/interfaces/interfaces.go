package interfaces

import "conference/pkg/common/utility"

type PrivateRepo interface {
	CreatePrivateRoom(utility.PrivateRoom) (uint, error)
	CheckLimit(string) (uint, error)
	CountParticipants(string) (uint, error)
	CheckParticipantPermission(string, string) (bool, error)
	AddParticipantInPrivateRoom(utility.PrivateRoomParticipants) error
	BlockParticipant(string, string) error
	RemoveParticipant(string, string) error
	UpdateParticipantExitTime(utility.PrivateRoomParticipants) error
	CheckType(string) (string, error)
	CheckInterest(string) (string, error)
	RemoveRoom(string) error
}

type GroupRepo interface {
}

type PublicRepo interface {
}
