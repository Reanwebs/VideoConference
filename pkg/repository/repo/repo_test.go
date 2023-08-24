package repo

import (
	"conference/pkg/common/utility"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateRoom(t *testing.T) {

	type args struct {
		input utility.PrivateRoom
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    uint
		wantErr error
	}{
		{
			name: "conference room created",
			args: args{
				input: utility.PrivateRoom{
					UserID:           "yourUserID",
					ConferenceID:     "conferenceUID",
					Type:             "conferenceType",
					Title:            "Conference Title",
					Description:      "Conference Description",
					Interest:         "Conference Interest",
					Recording:        true,
					Chat:             true,
					Broadcast:        true,
					Participantlimit: 100},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO conference_rooms(.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("yourUserID", "conferenceUID", "conferenceType", "Conference Title", "Conference Description", "Conference Interest", true, true, true, 100, time.Time{}, time.Time{}).
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CreatePrivateRoom(tt.args.input)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_CheckLimit(t *testing.T) {

	conferenceID := "conf102"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    uint
		wantErr error
	}{
		{
			name: "participant limit retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT participantlimit FROM conference_rooms WHERE conference_id = ?`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(conferenceID).
					WillReturnRows(sqlmock.NewRows([]string{"participantlimit"}).AddRow(100))
			},
			want:    100,
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CheckLimit(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CountParticipants(t *testing.T) {
	conferenceID := "conf102"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    uint
		wantErr error
	}{
		{
			name: "participant count retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT COUNT\(\*\) FROM conference_participants WHERE conference_id = $?`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(conferenceID).
					WillReturnRows(sqlmock.NewRows([]string{"participantcount"}).AddRow(0))
			},
			want:    0,
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CountParticipants(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckParticipantPermission(t *testing.T) {
	conferenceID := "conf102"
	userID := "user102"
	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wantErr error
	}{
		{
			name: "participant permission retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `SELECT permission FROM conference_participants`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(conferenceID, userID).
					WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(true))
			},
			want:    true,
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CheckParticipantPermission(conferenceID, userID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_AddParticipant(t *testing.T) {

	participantInput := utility.PrivateRoomParticipants{
		UserID:       "yourUserID",
		ConferenceID: "conf102",
		CamStatus:    "on",
		MicStatus:    "off",
		JoinTime:     time.Now(),
		ExitTime:     time.Now(),
		Role:         "participant",
	}

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant added",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO conference_participants(.+)$`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs("yourUserID", "conf102", "on", "off", sqlmock.AnyArg(), sqlmock.AnyArg(), "participant", sqlmock.AnyArg(), sqlmock.AnyArg()).
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

			repo := NewConferenceRepo(gormDB)

			err := repo.AddParticipantInPrivateRoom(participantInput)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_BlockParticipant(t *testing.T) {

	conferenceID := "conf202"
	userID := "UserID"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant blocked",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^UPDATE conference_participants SET permission = false`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs(conferenceID, userID).
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

			repo := NewConferenceRepo(gormDB)

			err := repo.BlockParticipant(conferenceID, userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UpdateParticipantExitTime(t *testing.T) {

	participantInput := utility.PrivateRoomParticipants{
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
				expectedQuery := `^UPDATE conference_participants SET exit_time = ?`
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

			repo := NewConferenceRepo(gormDB)

			err := repo.UpdateParticipantExitTime(participantInput)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemoveParticipant(t *testing.T) {

	conferenceID := "conf102"
	userID := "UserID"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "participant removed",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM conference_participants`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs(conferenceID, userID).
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

			repo := NewConferenceRepo(gormDB)

			err := repo.RemoveParticipant(conferenceID, userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckType(t *testing.T) {

	conferenceID := "conf111"
	conferenceType := "private"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "conference type retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT type FROM conference_rooms WHERE conference_id = ?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(conferenceID).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow(conferenceType))
			},
			want:    conferenceType,
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CheckType(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckInterest(t *testing.T) {

	conferenceID := "conf102"
	interest := "Conference Interest"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "interest retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT interest FROM conference_rooms WHERE conference_id = ?`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(conferenceID).
					WillReturnRows(sqlmock.NewRows([]string{"interest"}).AddRow(interest))
			},
			want:    interest,
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

			repo := NewConferenceRepo(gormDB)

			got, err := repo.CheckInterest(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemoveRoom(t *testing.T) {

	conferenceID := "conf102"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "room removed",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM conference_rooms WHERE conference_id = ?`
				mockSQL.ExpectExec(expectedQuery).
					WithArgs(conferenceID).
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

			repo := NewConferenceRepo(gormDB)

			err := repo.RemoveRoom(conferenceID)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
