package repo

import (
	"conference/pkg/common/utility"
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
)

func NewPublicConferenceRepo(db *gorm.DB) *conferenceRepo {
	return &conferenceRepo{
		DB: db,
	}
}

// Stream functions

func (c *conferenceRepo) CreateStreamRoom(input utility.StreamRoom) error {
	query := `
	INSERT INTO stream_rooms (host_id, stream_id, title, description, thumbnail_id, interest,status,avatar_id,user_name, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING id`

	var id uint
	err := c.DB.Raw(query,
		input.HostID,
		input.StreamID,
		input.Title,
		input.Description,
		input.ThumbnailID,
		input.Interest,
		input.Status,
		input.AvatarID,
		input.UserName,
		time.Now(),
		time.Now(),
	).Row().Scan(&id)

	if err != nil {
		return err
	}
	return nil
}

func (c *conferenceRepo) GetStream(streamID string) (utility.StreamRoom, error) {
	var stream utility.StreamRoom

	query := `
        SELECT host_id, stream_id, title, description, thumbnail_id, interest, status
        FROM stream_rooms
        WHERE stream_id = ?`

	result := c.DB.Raw(query, streamID).Row().Scan(
		&stream.HostID,
		&stream.StreamID,
		&stream.Title,
		&stream.Description,
		&stream.ThumbnailID,
		&stream.Interest,
		&stream.Status,
	)

	if result != nil {
		return utility.StreamRoom{}, result
	}

	return stream, nil
}

func (c *conferenceRepo) GetStreamList() ([]utility.StreamRoom, error) {
	var streams []utility.StreamRoom

	query := `
        SELECT host_id, stream_id, title, description, thumbnail_id, interest, status, avatar_id, user_name
        FROM stream_rooms
        WHERE status != 'Ended'`

	result := c.DB.Raw(query).Find(&streams)

	if result.Error != nil {
		return nil, result.Error
	}

	return streams, nil
}
func (c *conferenceRepo) GetSortedStreamList(filter string) ([]utility.StreamRoom, error) {
	var streams []utility.StreamRoom

	query := `
        SELECT host_id, stream_id, title, description, thumbnail_id, interest, status, avatar_id, user_name
        FROM stream_rooms
        WHERE status != 'Ended' AND interest = ?`

	result := c.DB.Raw(query, filter).Find(&streams)

	if result.Error != nil {
		return nil, result.Error
	}

	return streams, nil
}

func (c *conferenceRepo) UpdateStreamRoom(streamID string, hostID string, status string) error {
	query := `
	UPDATE stream_rooms
	SET status = ?, updated_at = ?
	WHERE stream_id = ?`

	result := c.DB.Exec(query,
		status,
		time.Now(),
		streamID,
	)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) AddStreamParticipants(input utility.StreamRoomParticipants) error {
	query := `
	INSERT INTO stream_room_participants (stream_id, participant_id, join_time, leave_time)
	VALUES (?, ?, ?, ?)`

	result := c.DB.Exec(query,
		input.StreamID,
		input.ParticipantID,
		input.JoinTime,
		input.LeaveTime,
	)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) UpdateStreamParticipants(input utility.StreamRoomParticipants) error {
	query := `
        UPDATE stream_room_participants
        SET leave_time = ?
        WHERE stream_id = ? AND participant_id = ?`

	result := c.DB.Exec(query,
		input.LeaveTime,
		input.StreamID,
		input.ParticipantID,
	)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *conferenceRepo) GetStreamJoinTime(streamID string, userID string) (time.Time, error) {
	var joinTime time.Time
	query := `
        SELECT join_time
        FROM stream_room_participants
        WHERE stream_id = ? AND participant_id = ?`

	if err := c.DB.Raw(query, streamID, userID).Row().Scan(&joinTime); err != nil {
		return time.Time{}, err
	}
	return joinTime, nil
}

// Conference functions

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
