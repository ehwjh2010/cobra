package setting

type Config struct {
	Application string `yaml:"application" json:"application"`
	ServerPort  int    `yaml:"serverPort" json:"serverPort"`
	Debug       bool   `yaml:"debug" json:"debug"`
	Log         logC   `yaml:"log" json:"log"`
	Mysql       mysqlC `yaml:"mysql" json:"mysql"`
	Redis       redisC `yaml:"redis" json:"redis"`
}

type mysqlC struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

type redisC struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Pwd  string `yaml:"pwd" json:"pwd"`
}

type logC struct {
	LogPath          string `yaml:"logPath" json:"logPath"`
	Level            string `yaml:"level" json:"level"`
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"`
}
