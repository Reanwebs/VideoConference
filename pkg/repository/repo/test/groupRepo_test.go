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

func Test_CreateGroupRoom(t *testing.T) {
	type args struct {
		input utility.GroupRoom
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    uint
		wantErr error
	}{
		{
			name: "group room created",
			args: args{
				input: utility.GroupRoom{
					UserID:           "yourUserID",
					ConferenceID:     "conferenceUID",
					GroupID:          "groupUID",
					Type:             "groupType",
					Title:            "Group Title",
					Description:      "Group Description",
					Interest:         "Group Interest",
					Recording:        true,
					Chat:             true,
					Broadcast:        true,
					Participantlimit: 50,
				},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO group_rooms(.+)$`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs("yourUserID", "conferenceUID", "groupUID", "groupType", "Group Title", "Group Description", "Group Interest", true, true, true, 50, time.Time{}, time.Time{}).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    1,
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

			repo := repo.NewGroupConferenceRepo(gormDB)

			err := repo.CreateGroupRoom(tt.args.input)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_AddParticipantInGroupRoom(t *testing.T) {
	type args struct {
		input utility.GroupRoomParticipants
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant added to group room",
			args: args{
				input: utility.GroupRoomParticipants{
					UserID:       "yourUserID",
					ConferenceID: "conferenceUID",
					GroupID:      "groupUID",
					Permission:   true,
					CamStatus:    "on",
					MicStatus:    "on",
					JoinTime:     time.Now(),
					ExitTime:     time.Now(),
					Role:         "participant",
				},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO group_room_participants(.+)$`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("yourUserID", "conferenceUID", "groupUID", true, "on", "on", sqlmock.AnyArg(), sqlmock.AnyArg(), "participant").
					WillReturnResult(sqlmock.NewResult(1, 1))
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			err := groupRepo.AddParticipantInGroupRoom(tt.args.input)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckGroupLimit(t *testing.T) {
	tests := []struct {
		name                 string
		conferenceID         string
		stub                 func(sqlmock.Sqlmock)
		wantParticipantLimit uint
		wantErr              error
	}{
		{
			name:         "group room limit fetched",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT participantlimit FROM group_rooms WHERE conference_id = ?`
				mockRows := sqlmock.NewRows([]string{"participantlimit"}).AddRow(50)
				mockSQL.ExpectQuery(expectedQuery).WithArgs("conferenceUID").WillReturnRows(mockRows)
			},
			wantParticipantLimit: 50,
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			gotParticipantLimit, err := groupRepo.CheckGroupLimit(tt.conferenceID)

			assert.Equal(t, tt.wantParticipantLimit, gotParticipantLimit)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CountGroupParticipants(t *testing.T) {
	tests := []struct {
		name                 string
		conferenceID         string
		stub                 func(sqlmock.Sqlmock)
		wantParticipantCount uint
		wantErr              error
	}{
		{
			name:         "group participants count fetched",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT COUNT\(\*\) FROM group_room_participants WHERE conference_id = ?`
				mockRows := sqlmock.NewRows([]string{"count"}).AddRow(10)
				mockSQL.ExpectQuery(expectedQuery).WithArgs("conferenceUID").WillReturnRows(mockRows)
			},
			wantParticipantCount: 10,
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			gotParticipantCount, err := groupRepo.CountGroupParticipants(tt.conferenceID)

			assert.Equal(t, tt.wantParticipantCount, gotParticipantCount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckGroupParticipantPermission(t *testing.T) {
	tests := []struct {
		name           string
		conferenceID   string
		userID         string
		stub           func(sqlmock.Sqlmock)
		wantPermission bool
		wantErr        error
	}{
		{
			name:         "group participant permission checked",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT permission FROM group_room_participants WHERE conference_id = \$1 AND user_id = \$2$`
				mockRows := sqlmock.NewRows([]string{"permission"}).AddRow(true)
				mockSQL.ExpectQuery(expectedQuery).WithArgs("conferenceUID", "userUID").WillReturnRows(mockRows)
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			gotPermission, err := groupRepo.CheckGroupParticipantPermission(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantPermission, gotPermission)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UpdateGroupParticipantExitTime(t *testing.T) {
	tests := []struct {
		name    string
		input   utility.GroupRoomParticipants
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "group participant exit time updated",
			input: utility.GroupRoomParticipants{
				UserID:       "userUID",
				ConferenceID: "conferenceUID",
				ExitTime:     time.Now(),
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^UPDATE group_room_participants SET exit_time = \$1 WHERE user_id = \$2 AND conference_id = \$3$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).WithArgs(sqlmock.AnyArg(), "userUID", "conferenceUID").WillReturnResult(mockResult)
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			err := groupRepo.UpdateGroupParticipantExitTime(tt.input)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemoveGroupParticipant(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		userID       string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "group participant removed",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM group_room_participants WHERE conference_id = \$1 AND user_id = \$2$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).WithArgs("conferenceUID", "userUID").WillReturnResult(mockResult)
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			err := groupRepo.RemoveGroupParticipant(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_BlockGroupParticipant(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		userID       string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "group participant blocked",
			conferenceID: "conferenceUID",
			userID:       "userUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^UPDATE group_room_participants SET permission = false WHERE conference_id = \$1 AND user_id = \$2$`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).WithArgs("conferenceUID", "userUID").WillReturnResult(mockResult)
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			err := groupRepo.BlockGroupParticipant(tt.conferenceID, tt.userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemoveGroupRoom(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		stub         func(sqlmock.Sqlmock)
		wantErr      error
	}{
		{
			name:         "group room removed",
			conferenceID: "conferenceUID",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM group_rooms WHERE conference_id = ?`
				mockResult := sqlmock.NewResult(0, 1)
				mockSQL.ExpectExec(expectedQuery).WithArgs("conferenceUID").WillReturnResult(mockResult)
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

			groupRepo := repo.NewGroupConferenceRepo(gormDB)

			err := groupRepo.RemoveGroupRoom(tt.conferenceID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
