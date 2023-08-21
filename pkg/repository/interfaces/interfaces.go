package interfaces

import "conference/pkg/common/utility"

type ConferenceRepo interface {
	CreateRoom(utility.ConferenceRoom) (uint, error)
	CheckLimit(int32) (uint, error)
	CountParticipants(int32) (uint, error)
	CheckParticipantPermission(int32, string) (bool, error)
	AddParticipant(utility.ConferenceParticipants) error
	BlockParticipant(int32, string) error
	RemoveParticipant(int32, string) error
	UpdateParticipantExitTime(utility.ConferenceParticipants) error
	CheckType(int32) (string, error)
	CheckInterest(int32) (string, error)
	RemoveRoom(int32) error
}
