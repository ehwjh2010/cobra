package time

import (
	"github.com/ehwjh2010/viper/client/enums"
	"strings"
	"sync"
	"time"
)

var locationMap sync.Map

// GetBJLocation 东八区
func GetBJLocation() *time.Location {
	location, _ := GetLocationByName(enums.BJ)
	return location
}

// GetUTCLocation 获取UTC时区
func GetUTCLocation() *time.Location {
	return time.UTC
}

// GetLocationByName 根据名字获取时区
func GetLocationByName(name string) (*time.Location, error) {
	name = strings.ToUpper(name)
	if loc, ok := locationMap.Load(name); ok {
		location := loc.(*time.Location)
		return location, nil
	}

	if location, err := time.LoadLocation(name); err == nil {
		locationMap.LoadOrStore(name, location)
	} else {
		return nil, err
	}

	loc, _ := locationMap.Load(name)
	location := loc.(*time.Location)
	return location, nil
}
