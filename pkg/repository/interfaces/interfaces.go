package interfaces

type Conference interface {
	CreateRoom(string)
	AddParticipant()
	CheckLimit()
	CheckType()
	CheckInterest()
	RemoveParticipant()
	RemoveRoom()
}
