package timer

import (
	"github.com/ehwjh2010/viper/global"
	"strings"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

var locationMap = map[string]*time.Location{"UTC": time.UTC}

//storeLoc 存储时区
func storeLoc(name string, loc *time.Location) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if _, exist := locationMap[name]; exist {
		return
	}

	if loc != nil {
		locationMap[strings.ToUpper(name)] = loc
	}
}

//queryLoc 查询时区
func queryLoc(name string) (*time.Location, bool) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	loc, exist := locationMap[name]
	return loc, exist
}

//GetBJLocation 东八区
func GetBJLocation() *time.Location {
	location, _ := GetLocationByName(global.BJ)
	return location
}

//GetUTCLocation 获取UTC时区
func GetUTCLocation() *time.Location {
	return time.UTC
}

//GetLocationByName 根据名字获取时区
func GetLocationByName(name string) (*time.Location, error) {
	name = strings.ToUpper(name)
	if location, ok := queryLoc(name); ok {
		return location, nil
	}

	if location, err := time.LoadLocation(name); err == nil {
		storeLoc(name, location)
	} else {
		return nil, err
	}

	location, _ := queryLoc(name)
	return location, nil
}
