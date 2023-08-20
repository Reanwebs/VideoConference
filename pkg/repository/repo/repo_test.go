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
		input utility.ConferenceRoom
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
				input: utility.ConferenceRoom{
					UserID:           "yourUserID",
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
				mockSQL.ExpectQuery(expectedQuery).WithArgs("yourUserID", "conferenceType", "Conference Title", "Conference Description", "Conference Interest", true, true, true, 100, time.Time{}, time.Time{}).
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

			u := NewConferenceRepo(gormDB)

			got, err := u.CreateRoom(tt.args.input)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}
