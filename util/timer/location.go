package timer

import (
	"strings"
	"time"
)

const BJ = "Asia/Shanghai"

var locationMap = map[string]*time.Location{"UTC": time.UTC}

//GetBJLocation 东八区
func GetBJLocation() *time.Location {
	location, _ := GetLocationByName(BJ)
	return location
}

//GetUTCLocation 获取UTC时区
func GetUTCLocation() *time.Location {
	return time.UTC
}

//GetLocationByName 根据名字获取时区
func GetLocationByName(name string) (*time.Location, error) {
	name = strings.ToUpper(name)
	if location, ok := locationMap[name]; ok {
		return location, nil
	}

	if location, err := time.LoadLocation(name); err == nil {
		locationMap[name] = location
	} else {
		return nil, err
	}

	location, _ := locationMap[name]
	return location, nil
}
