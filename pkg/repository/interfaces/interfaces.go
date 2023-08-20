package interfaces

import "conference/pkg/common/utility"

type ConferenceRepo interface {
	CreateRoom(utility.ConferenceRoom) (uint, error)
	AddParticipant()
	CheckLimit()
	CheckType()
	CheckInterest()
	RemoveParticipant()
	RemoveRoom()
}
