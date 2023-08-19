package interfaces

type Conference interface {
	CreateRoom()
	AddParticipant()
	CheckLimit()
	CheckType()
	CheckInterest()
	RemoveParticipant()
	RemoveRoom()
}
