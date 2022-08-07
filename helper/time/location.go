package time

import (
	"strings"
	"sync"
	"time"

	"github.com/ehwjh2010/viper/enums"
)

var locationMap sync.Map

// GetBJLocation 东八区.
func GetBJLocation() *time.Location {
	location, _ := GetLocationByName(enums.BJ)
	return location
}

// GetUTCLocation 获取UTC时区.
func GetUTCLocation() *time.Location {
	return time.UTC
}

// GetLocationByName 根据名字获取时区.
func GetLocationByName(name string) (*time.Location, error) {
	name = strings.ToUpper(name)
	if loc, ok := locationMap.Load(name); ok {
		location := loc.(*time.Location) //nolint:errcheck
		return location, nil
	}

	location, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	locationMap.Store(name, location)
	return location, nil
}
