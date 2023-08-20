package repo

import (
	"conference/pkg/common/utility"
	"fmt"

	"gorm.io/gorm"
)

type conferenceRepo struct {
	DB *gorm.DB
}

func NewConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreateRoom(input utility.ConferenceRoom) (uint, error) {
	query := `
        INSERT INTO conference_rooms (user_id, type, title, description, interest, recording, chat, broadcast, participantlimit, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.UserID,
		input.Type,
		input.Title,
		input.Description,
		input.Interest,
		input.Recording,
		input.Chat,
		input.Broadcast,
		input.Participantlimit,
		input.CreatedAt,
		input.UpdatedAt,
	).Row().Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *conferenceRepo) AddParticipant() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckLimit() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckType() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckInterest() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) RemoveParticipant() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) RemoveRoom() {
	fmt.Println("sssssssssssssssssssss")
}
