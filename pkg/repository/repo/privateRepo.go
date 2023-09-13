package repo

import (
	"conference/pkg/common/utility"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type conferenceRepo struct {
	DB *gorm.DB
}

func NewPrivateConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

func (c *conferenceRepo) CreatePrivateSchedule(input utility.ScheduleConference) (uint, error) {
	query := `
        INSERT INTO schedule_conferences (user_id, schedule_id, title, description, interest, time, duration, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, NULL)
        RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.UserId,
		input.ScheduleID,
		input.Title,
		input.Description,
		input.Interest,
		input.Time,
		input.Duration,
		time.Now(),
	).Row().Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *conferenceRepo) GetPrivateSchedules(userID string) ([]utility.ScheduleConference, error) {
	query := `
        SELECT id, schedule_id, title, description, interest, time, duration
        FROM schedule_conferences
        WHERE user_id = ?`

	var schedules []utility.ScheduleConference
	err := c.DB.Raw(query, userID).Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (c *conferenceRepo) GetCompletedSchedules(userID string) ([]utility.ScheduleConference, error) {
	return nil, nil
}

func (c *conferenceRepo) CreatePrivateRoom(input utility.PrivateRoom) (uint, error) {
	query := `
        INSERT INTO private_rooms (user_id,conference_id,sdp_offer,ice_candidate, type, title, description, interest, recording, chat, broadcast, participantlimit, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.UserID,
		input.ConferenceID,
		input.SdpOffer,
		input.IceCandidate,
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

func (c *conferenceRepo) CheckPrivateLimit(conferenceID string) (uint, error) {
	query := `
        SELECT participantlimit
        FROM private_rooms
        WHERE conference_id = ?`

	var participantLimit uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantLimit)
	if err != nil {
		return 0, err
	}

	return participantLimit, nil
}

func (c *conferenceRepo) CountPrivateParticipants(conferenceID string) (uint, error) {
	query := `
        SELECT COUNT(*)
        FROM private_room_participants
        WHERE conference_id = ?`

	var participantCount uint
	err := c.DB.Raw(query, conferenceID).Row().Scan(&participantCount)
	if err != nil {
		return 0, err
	}
	fmt.Println(participantCount)
	return participantCount, nil
}

func (c *conferenceRepo) CheckPrivateParticipantPermission(conferenceID string, userID string) (bool, error) {
	query := `
        SELECT permission
        FROM private_room_participants
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

func (c *conferenceRepo) GetSdpOffer(conferenceID string) (string, error) {
	query := `
        SELECT sdp_offer
        FROM private_rooms
        WHERE conference_id = ?`

	var sdpOffer string
	err := c.DB.Raw(query, conferenceID).Row().Scan(&sdpOffer)
	if err != nil {
		return "", err
	}

	return sdpOffer, nil
}

func (c *conferenceRepo) AddParticipantInPrivateRoom(input utility.PrivateRoomParticipants) error {
	query := `
        INSERT INTO private_room_participants (user_id, conference_id,permission, cam_status, mic_status, join_time, exit_time, role, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result := c.DB.Exec(query,
		input.UserID,
		input.ConferenceID,
		input.Permission,
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

func (c *conferenceRepo) BlockPrivateParticipant(conferenceID string, userID string) error {
	query := `
        UPDATE private_room_participants
        SET permission = false
        WHERE conference_id = ? AND user_id = ?`

	result := c.DB.Exec(query, conferenceID, userID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (c *conferenceRepo) UpdatePrivateParticipantExitTime(input utility.PrivateRoomParticipants) error {
	query := `
        UPDATE private_room_participants
        SET exit_time = ?
        WHERE user_id = ? AND conference_id = ?`

	result := c.DB.Exec(query, input.ExitTime, input.UserID, input.ConferenceID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (c *conferenceRepo) GetJoinTime(conferenceID string, userID string) (time.Time, error) {
	query := `
        SELECT join_time
        FROM private_room_participants
        WHERE conference_id = ? AND user_id = ?`

	var joinTime time.Time
	err := c.DB.Raw(query, conferenceID, userID).Row().Scan(&joinTime)
	if err != nil {
		return time.Time{}, err
	}

	return joinTime, nil
}

func (c *conferenceRepo) RemovePrivateParticipant(conferenceID string, userID string) error {
	query := `
        DELETE FROM private_room_participants
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
        FROM private_rooms
        WHERE conference_id = ?`

	var conferenceType string
	err := c.DB.Raw(query, conferenceID).Row().Scan(&conferenceType)
	if err != nil {
		return "", err
	}

	return conferenceType, nil
}

func (c *conferenceRepo) CheckPrivateInterest(conferenceID string) (string, error) {
	query := `
        SELECT interest
        FROM private_rooms
        WHERE conference_id = ?`

	var interest string
	err := c.DB.Raw(query, conferenceID).Row().Scan(&interest)
	if err != nil {
		return "", err
	}

	return interest, nil
}

func (c *conferenceRepo) RemovePrivateRoom(conferenceID string) error {
	query := `
        DELETE FROM private_rooms
        WHERE conference_id = ?`

	result := c.DB.Exec(query, conferenceID)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
