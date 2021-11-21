package timeutils

import (
	"time"
)

const BJ = "Asia/Shanghai"

var BJZone *time.Location

//GetBJLocation 东八区
func GetBJLocation() *time.Location {
	if BJZone != nil {
		return BJZone
	}

	location, _ := time.LoadLocation(BJ)
	BJZone = location

	return BJZone
}

//GetUTCLocation 获取UTC时区
func GetUTCLocation() *time.Location {
	return time.UTC
}

//GetLocationByName 根据名字获取时区
func GetLocationByName(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
