package interfaces

import "conference/pkg/common/utility"

type ConferenceRepo interface {
	CreateRoom(utility.ConferenceRoom) (uint, error)
	CheckLimit(string) (uint, error)
	CountParticipants(string) (uint, error)
	CheckParticipantPermission(string, string) (bool, error)
	AddParticipant(utility.ConferenceParticipants) error
	BlockParticipant(string, string) error
	RemoveParticipant(string, string) error
	UpdateParticipantExitTime(utility.ConferenceParticipants) error
	CheckType(string) (string, error)
	CheckInterest(string) (string, error)
	RemoveRoom(string) error
}
