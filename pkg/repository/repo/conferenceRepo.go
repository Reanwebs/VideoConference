package repo

import (
	"fmt"

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

func (c *conferenceRepo) CreateRoom(s string) {
	fmt.Println("sssssssssssssssssssss", s)
}

func (c *conferenceRepo) AddParticipant() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckLimit() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckType() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) CheckInterest() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) RemoveParticipant() {
	fmt.Println("sssssssssssssssssssss")
}
func (c *conferenceRepo) RemoveRoom() {
	fmt.Println("sssssssssssssssssssss")
}
