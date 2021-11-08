package utils

import (
	"fmt"
	"time"
)

const BJ = "Asia/Shanghai"

var BJZone *time.Location

func GetBJLocation() *time.Location {
	if BJZone != nil {
		return BJZone
	}

	location, err := time.LoadLocation(BJ)
	if err != nil {
		Panic(fmt.Sprintf("Access beijing location failed, err: %v", err))
	}

	BJZone = location

	return location
}
