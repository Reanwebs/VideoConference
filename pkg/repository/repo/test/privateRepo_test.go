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

func Test_CreatePrivateRoom(t *testing.T) {

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
					SdpOffer:         "sdpOfferString",
					IceCandidate:     "iceCandidateString",
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

				expectedQuery := `^INSERT INTO private_rooms(.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("yourUserID", "conferenceUID", "sdpOfferString", "iceCandidateString", "conferenceType", "Conference Title", "Conference Description", "Conference Interest", true, true, true, 100, time.Time{}, time.Time{}).
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CreatePrivateRoom(tt.args.input)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_CheckPrivateLimit(t *testing.T) {

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
				expectedQuery := `^SELECT participantlimit FROM private_rooms WHERE conference_id = ?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CheckPrivateLimit(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CountPrivateParticipants(t *testing.T) {
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
				expectedQuery := `^SELECT COUNT\(\*\) FROM private_room_participants WHERE conference_id = $?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CountPrivateParticipants(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckPrivateParticipantPermission(t *testing.T) {
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
				expectedQuery := `SELECT permission FROM private_room_participants WHERE conference_id = \$1 AND user_id = \$2$`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CheckPrivateParticipantPermission(conferenceID, userID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetSdpOffer(t *testing.T) {
	conferenceID := "conf102"
	sdpOffer := "sample_sdp_offer"

	tests := []struct {
		name    string
		stub    func(mockSQL sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "SDP offer retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT sdp_offer FROM private_rooms WHERE conference_id = ?`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(conferenceID).
					WillReturnRows(sqlmock.NewRows([]string{"sdp_offer"}).AddRow(sdpOffer))
			},
			want:    sdpOffer,
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.GetSdpOffer(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_AddParticipantInPrivateRoom(t *testing.T) {

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
				expectedQuery := `^INSERT INTO private_room_participants(.+)$`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			err := privateRepo.AddParticipantInPrivateRoom(participantInput)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_BlockPrivateParticipant(t *testing.T) {

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
				expectedQuery := `^UPDATE private_room_participants SET permission = false`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			err := privateRepo.BlockPrivateParticipant(conferenceID, userID)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_UpdatePrivateParticipantExitTime(t *testing.T) {

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
				expectedQuery := `^UPDATE private_room_participants SET exit_time = ?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			err := privateRepo.UpdatePrivateParticipantExitTime(participantInput)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetJoinTime(t *testing.T) {

	conferenceID := "conf122"
	userID := "yourUserID"
	joinTime := time.Now()

	tests := []struct {
		name    string
		stub    func(mockSQL sqlmock.Sqlmock)
		want    time.Time
		wantErr error
	}{
		{
			name: "join time retrieved",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^SELECT join_time FROM private_room_participants WHERE conference_id = \$1 AND user_id = \$2`
				mockSQL.ExpectQuery(expectedQuery).
					WithArgs(conferenceID, userID).
					WillReturnRows(sqlmock.NewRows([]string{"join_time"}).AddRow(joinTime))
			},
			want:    joinTime,
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.GetJoinTime(conferenceID, userID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemovePrivateParticipant(t *testing.T) {

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
				expectedQuery := `^DELETE FROM private_room_participants`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			err := privateRepo.RemovePrivateParticipant(conferenceID, userID)

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
				expectedQuery := `^SELECT type FROM private_rooms WHERE conference_id = ?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CheckType(conferenceID)

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
				expectedQuery := `^SELECT interest FROM private_rooms WHERE conference_id = ?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			got, err := privateRepo.CheckPrivateInterest(conferenceID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemovePrivateRoom(t *testing.T) {

	conferenceID := "conf102"

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "room removed",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^DELETE FROM private_rooms WHERE conference_id = ?`
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

			privateRepo := repo.NewPrivateConferenceRepo(gormDB)

			err := privateRepo.RemovePrivateRoom(conferenceID)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
