package repo

import (
	"conference/pkg/common/utility"

	"gorm.io/gorm"
)

func NewGroupConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreateGroupRoom(input utility.GroupRoom) error {
	return nil
}
func (c *conferenceRepo) AddParticipantInGroupRoom(input utility.GroupRoomParticipants) error {
	return nil
}
