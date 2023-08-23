package repo

import (
	"conference/pkg/common/utility"
	"fmt"
	"time"

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
        INSERT INTO conference_rooms (user_id,conference_id, type, title, description, interest, recording, chat, broadcast, participantlimit, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.UserID,
		input.ConferenceID,
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

func (c *conferenceRepo) CheckLimit(conferenceID string) (uint, error) {
	query := `
        SELECT participantlimit
        FROM conference_rooms
        WHERE id = ?`

	var participantLimit uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantLimit)
	if err != nil {
		return 0, err
	}

	return participantLimit, nil
}

func (c *conferenceRepo) CountParticipants(conferenceID string) (uint, error) {
	query := `
        SELECT COUNT(*)
        FROM conference_participants
        WHERE conference_id = ?`

	var participantCount uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantCount)
	if err != nil {
		return 0, err
	}
	fmt.Println(participantCount)
	return participantCount, nil
}

func (c *conferenceRepo) CheckParticipantPermission(conferenceID string, userID string) (bool, error) {
	query := `
        SELECT permission
        FROM conference_participants
        WHERE conference_id = ? AND user_id = ?`

	var permission bool
	err := c.DB.Raw(query, conferenceID, userID).Row().Scan(&permission)
	if err != nil {
		return false, err
	}

	return permission, nil
}

func (c *conferenceRepo) AddParticipant(input utility.ConferenceParticipants) error {
	query := `
        INSERT INTO conference_participants (user_id, conference_id, cam_status, mic_status, join_time, exit_time, role, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result := c.DB.Exec(query,
		input.UserID,
		input.ConferenceID,
		input.CamStatus,
		input.MicStatus,
		input.JoinTime,
		input.ExitTime,
		input.Role,
		time.Now(),
		time.Now(),
	)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}

}

func (c *conferenceRepo) BlockParticipant(conferenceID string, userID string) error {
	query := `
        UPDATE conference_participants
        SET permission = false
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (c *conferenceRepo) UpdateParticipantExitTime(input utility.ConferenceParticipants) error {
	query := `
        UPDATE conference_participants
        SET exit_time = ?
        WHERE user_id = ? AND conference_id = ?`

	result := c.DB.Exec(query, input.ExitTime, input.UserID, input.ConferenceID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (c *conferenceRepo) RemoveParticipant(conferenceID string, userID string) error {
	query := `
        DELETE FROM conference_participants
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (c *conferenceRepo) CheckType(conferenceID string) (string, error) {
	query := `
        SELECT type
        FROM conference_rooms
        WHERE id = ?`

	var conferenceType string
	err := c.DB.Raw(query, conferenceID).Row().Scan(&conferenceType)
	if err != nil {
		return "", err
	}

	return conferenceType, nil
}

func (c *conferenceRepo) CheckInterest(conferenceID string) (string, error) {
	query := `
        SELECT interest
        FROM conference_rooms
        WHERE id = ?`

	var interest string
	err := c.DB.Raw(query, conferenceID).Row().Scan(&interest)
	if err != nil {
		return "", err
	}

	return interest, nil
}

func (c *conferenceRepo) RemoveRoom(conferenceID string) error {
	query := `
        DELETE FROM conference_rooms
        WHERE id = ?`

	result := c.DB.Exec(query, conferenceID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
