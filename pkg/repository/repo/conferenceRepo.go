package repo

import (
	"conference/pkg/repository/interfaces"
	"fmt"
)

type ConferenceRepo struct {
	interfaces.Conference
}

func (ConferenceRepo) CreateRoom(s string) {
	fmt.Println(s)
}
