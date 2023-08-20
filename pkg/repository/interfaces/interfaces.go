package interfaces

type ConferenceRepo interface {
	CreateRoom(string)
	AddParticipant()
	CheckLimit()
	CheckType()
	CheckInterest()
	RemoveParticipant()
	RemoveRoom()
}
