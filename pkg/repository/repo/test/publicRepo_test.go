package test

import (
	"conference/pkg/common/utility"
	"conference/pkg/repository/repo"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreatePublicRoom(t *testing.T) {

	type args struct {
		input utility.PublicRoom
	}
	tests := []struct {
		name    string
		args    args
		input   utility.PublicRoom
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "public room created",
			input: utility.PublicRoom{
				UserID:           "yourUserID",
				ConferenceID:     "conferenceUID",
				Type:             "publicType",
				Title:            "Public Title",
				Description:      "Public Description",
				Interest:         "Public Interest",
				Recording:        true,
				Chat:             true,
				Broadcast:        true,
				Participantlimit: 200,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO public_rooms(.+)$`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs("yourUserID", "conferenceUID", "publicType", "Public Title", "Public Description", "Public Interest", true, true, true, 200, time.Time{}, time.Time{}).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			err := publicRepo.CreatePublicRoom(tt.input)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_AddParticipantInPublicRoom(t *testing.T) {
	tests := []struct {
		name    string
		input   utility.PublicRoomParticipants
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant added to public room",
			input: utility.PublicRoomParticipants{
				UserID:       "userUID",
				ConferenceID: "conferenceUID",
				Permission:   true,
				CamStatus:    "active",
				MicStatus:    "active",
				JoinTime:     time.Now(),
				ExitTime:     time.Now(),
				Role:         "participant",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO public_room_participants(.+)$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("userUID", "conferenceUID", true, "active", "active", sqlmock.AnyArg(), sqlmock.AnyArg(), "participant").
					WillReturnResult(mockResult)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			err := publicRepo.AddParticipantInPublicRoom(tt.input)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckPublicLimit(t *testing.T) {
	tests := []struct {
		name                 string
		conferenceID         string
		stub                 func(sqlmock.Sqlmock)
		wantParticipantLimit uint
		wantErr              error
	}{
		{
			name:         "check public room participant limit",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT participantlimit FROM public_rooms WHERE conference_id = ?`
				mockRows := sqlmock.NewRows([]string{"participantlimit"}).AddRow(100)
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs("conferenceUID").
					WillReturnRows(mockRows)
			},
			wantParticipantLimit: 100,
			wantErr:              nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			participantLimit, err := publicRepo.CheckPublicLimit(tt.conferenceID)

			assert.Equal(t, tt.wantParticipantLimit, participantLimit)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CountPublicParticipants(t *testing.T) {
	tests := []struct {
		name                 string
		conferenceID         string
		stub                 func(sqlmock.Sqlmock)
		wantParticipantCount uint
		wantErr              error
	}{
		{
			name:         "count public room participants",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT COUNT\(\*\) FROM public_room_participants WHERE conference_id = ?`
				mockRows := sqlmock.NewRows([]string{"count"}).AddRow(50)
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs("conferenceUID").
					WillReturnRows(mockRows)
			},
			wantParticipantCount: 50,
			wantErr:              nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			participantCount, err := publicRepo.CountPublicParticipants(tt.conferenceID)

			assert.Equal(t, tt.wantParticipantCount, participantCount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckPublicParticipantPermission(t *testing.T) {
	tests := []struct {
		name           string
		conferenceID   string
		userID         string
		stub           func(sqlmock.Sqlmock)
		wantPermission bool
		wantErr        error
	}{
		{
			name:         "check public room participant permission - user exists",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT permission FROM public_room_participants WHERE conference_id = \$1 AND user_id = \$2$`
				mockRows := sqlmock.NewRows([]string{"permission"}).AddRow(true)
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs("conferenceUID", "userUID").
					WillReturnRows(mockRows)
			},
			wantPermission: true,
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			permission, err := publicRepo.CheckPublicParticipantPermission(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantPermission, permission)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UpdatePublicParticipantExitTime(t *testing.T) {
	participantInput := utility.PublicRoomParticipants{
		UserID:       "yourUserID",
		ConferenceID: "conf122",
		ExitTime:     time.Now(),
	}

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant exit time updated",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^UPDATE public_room_participants SET exit_time = ?`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs(participantInput.ExitTime, participantInput.UserID, participantInput.ConferenceID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			err := publicRepo.UpdatePublicParticipantExitTime(participantInput)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemovePublicParticipant(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		userID       string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "remove public room participant",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM public_room_participants WHERE conference_id = \$1 AND user_id = \$2$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("conferenceUID", "userUID").
					WillReturnResult(mockResult)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			err := publicRepo.RemovePublicParticipant(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_BlockPublicParticipant(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		userID       string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "block public room participant",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^UPDATE public_room_participants SET permission = false WHERE conference_id = \$1 AND user_id = \$2$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("conferenceUID", "userUID").
					WillReturnResult(mockResult)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)
			err := publicRepo.BlockPublicParticipant(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemovePublicRoom(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "remove public room",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM public_rooms WHERE conference_id = ?`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("conferenceUID").
					WillReturnResult(mockResult)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			publicRepo := repo.NewPublicConferenceRepo(gormDB)

			err := publicRepo.RemovePublicRoom(tt.conferenceID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

//
