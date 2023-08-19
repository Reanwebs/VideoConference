package repo

import (
	"conference/pkg/repository/interfaces"
	"fmt"
)

type ConferenceRepo struct {
	interfaces.Conference
}

func (c *ConferenceRepo) CreateRoom(s string) {
	fmt.Println("sssssssssssssssssssss", s)
}
