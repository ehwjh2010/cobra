package setting

import "fmt"

type Config struct {
	Application string       `yaml:"application" json:"application"`
	ServerPort  int          `yaml:"serverPort" json:"serverPort"`
	Debug       bool         `yaml:"debug" json:"debug"`
	LogConfig   *LogConfig   `yaml:"log" json:"log"`
	MysqlConfig *MysqlConfig `yaml:"mysql" json:"mysql"`
	RedisConfig *RedisConfig `yaml:"redis" json:"redis"`
}

type MysqlConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
}

func (c *MysqlConfig) Uri() string {
	uri := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		c.User, c.Password, c.Host, c.Port, c.Database)
	return uri
}

type RedisConfig struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Pwd  string `yaml:"pwd" json:"pwd"`
}

type LogConfig struct {
	LogPath          string `yaml:"logPath" json:"logPath"`
	Level            string `yaml:"level" json:"level"`
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"`
	AccessMethodRow  bool   `yaml:"accessMethodRow" json:"accessMethodRow"`
}
