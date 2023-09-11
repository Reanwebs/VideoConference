package repo

import (
	"conference/pkg/common/utility"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

func NewPublicConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreatePublicSchedule(input utility.SchedulePublicConference) (uint, error) {
	return 1, nil
}

func (c *conferenceRepo) CreatePublicRoom(input utility.PublicRoom) error {
	query := `
        INSERT INTO public_rooms (user_id, conference_id, type, title, description, interest, recording, chat, broadcast, participantlimit, created_at, updated_at)
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
		return err
	}
	return nil

}

func (c *conferenceRepo) AddParticipantInPublicRoom(input utility.PublicRoomParticipants) error {
	query := `
        INSERT INTO public_room_participants (user_id, conference_id, permission, cam_status, mic_status, join_time, exit_time, role)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	err := c.DB.Exec(query,
		input.UserID,
		input.ConferenceID,
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

func (c *conferenceRepo) CheckPublicLimit(conferenceID string) (uint, error) {
	query := `
        SELECT participantlimit
        FROM public_rooms
        WHERE conference_id = ?`

	var participantLimit uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantLimit)
	if err != nil {
		return 0, err
	}

	return participantLimit, nil
}

func (c *conferenceRepo) CountPublicParticipants(conferenceID string) (uint, error) {
	query := `
        SELECT COUNT(*)
        FROM public_room_participants
        WHERE conference_id = ?`

	var participantCount uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantCount)
	if err != nil {
		return 0, err
	}

	return participantCount, nil
}

func (c *conferenceRepo) CheckPublicParticipantPermission(conferenceID string, userID string) (bool, error) {
	query := `
        SELECT permission
        FROM public_room_participants
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

func (c *conferenceRepo) UpdatePublicParticipantExitTime(input utility.PublicRoomParticipants) error {
	query := `
        UPDATE public_room_participants
        SET exit_time = ?
        WHERE user_id = ? AND conference_id = ?`

	result := c.DB.Exec(query, input.ExitTime, input.UserID, input.ConferenceID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) RemovePublicParticipant(conferenceID string, userID string) error {
	query := `
        DELETE FROM public_room_participants
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) BlockPublicParticipant(conferenceID string, userID string) error {
	query := `
        UPDATE public_room_participants
        SET permission = false
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) RemovePublicRoom(conferenceID string) error {
	query := `
        DELETE FROM public_rooms
        WHERE conference_id = ?`

	result := c.DB.Exec(query, conferenceID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
