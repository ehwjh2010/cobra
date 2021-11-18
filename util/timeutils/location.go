package timeutils

import (
	"time"
)

const BJ = "Asia/Shanghai"

var BJZone *time.Location

//GetBJLocation 获取北京时区
func GetBJLocation() (*time.Location, error) {
	if BJZone != nil {
		return BJZone, nil
	}

	location, err := time.LoadLocation(BJ)
	BJZone = location

	return BJZone, err
}

//GetUtcLocation 获取UTC时区
func GetUtcLocation() *time.Location {
	return time.UTC
}

//GetLocationByName 根据名字获取时区
func GetLocationByName(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
