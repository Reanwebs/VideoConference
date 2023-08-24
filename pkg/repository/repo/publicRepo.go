package repo

import (
	"conference/pkg/common/utility"

	"gorm.io/gorm"
)

func NewPublicConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreatePublicRoom(input utility.PublicRoom) error {
	return nil
}
func (c *conferenceRepo) AddParticipantInPublicRoom(input utility.PublicRoomParticipants) error {
	return nil
}
