package setting

type Config struct {
	Application string      `yaml:"application" json:"application"`
	ServerPort  int         `yaml:"serverPort" json:"serverPort"`
	Debug       bool        `yaml:"debug" json:"debug"`
	LogConfig   LogConfig   `yaml:"log" json:"log"`
	MysqlConfig MysqlConfig `yaml:"mysql" json:"mysql"`
	RedisConfig RedisConfig `yaml:"redis" json:"redis"`
}

type MysqlConfig struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
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
