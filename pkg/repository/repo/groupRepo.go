package repo

import (
	"conference/pkg/common/utility"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func NewGroupConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreateGroupRoom(input utility.GroupRoom) error {
	query := `
        INSERT INTO group_rooms (user_id, conference_id, group_id, type, title, description, interest, recording, chat, broadcast, participantlimit, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.UserID,
		input.ConferenceID,
		input.GroupID,
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
		return err
	}
	return nil
}

func (c *conferenceRepo) AddParticipantInGroupRoom(input utility.GroupRoomParticipants) error {
	query := `
        INSERT INTO group_room_participants (user_id, conference_id, group_id, permission, cam_status, mic_status, join_time, exit_time, role)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	err := c.DB.Exec(query,
		input.UserID,
		input.ConferenceID,
		input.GroupID,
		input.Permission,
		input.CamStatus,
		input.MicStatus,
		input.JoinTime,
		input.ExitTime,
		input.Role,
	)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (c *conferenceRepo) CheckGroupLimit(conferenceID string) (uint, error) {
	query := `
        SELECT participantlimit
        FROM group_rooms
        WHERE conference_id = ?`

	var participantLimit uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantLimit)
	if err != nil {
		return 0, err
	}

	return participantLimit, nil
}

func (c *conferenceRepo) CountGroupParticipants(conferenceID string) (uint, error) {
	query := `
        SELECT COUNT(*)
        FROM group_room_participants
        WHERE conference_id = ?`

	var participantCount uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantCount)
	if err != nil {
		return 0, err
	}
	fmt.Println(participantCount)
	return participantCount, nil
}

func (c *conferenceRepo) CheckGroupParticipantPermission(conferenceID string, userID string) (bool, error) {
	query := `
        SELECT permission
        FROM group_room_participants
        WHERE conference_id = ? AND user_id = ?`

	var permission bool
	err := c.DB.Raw(query, conferenceID, userID).Row().Scan(&permission)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		} else {
			return false, err
		}
	}
	return permission, nil
}

func (c *conferenceRepo) UpdateGroupParticipantExitTime(input utility.GroupRoomParticipants) error {
	query := `
        UPDATE group_room_participants
        SET exit_time = ?
        WHERE user_id = ? AND conference_id = ?`

	result := c.DB.Exec(query, input.ExitTime, input.UserID, input.ConferenceID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) RemoveGroupParticipant(conferenceID string, userID string) error {
	query := `
        DELETE FROM group_room_participants
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) BlockGroupParticipant(conferenceID string, userID string) error {
	query := `
        UPDATE group_room_participants
        SET permission = false
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) RemoveGroupRoom(conferenceID string) error {
	query := `
        DELETE FROM group_rooms
        WHERE conference_id = ?`

	result := c.DB.Exec(query, conferenceID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
