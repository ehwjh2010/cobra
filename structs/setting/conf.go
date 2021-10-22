package setting

type Config struct {
	Application      string `yaml:"application" json:"application"`
	ServerPort       int    `yaml:"serverPort" json:"serverPort"`
	Debug            bool   `yaml:"debug" json:"debug"`
	Log              log    `yaml:"log" json:"log"`
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"`
	Mysql            mysql  `yaml:"mysql" json:"mysql"`
	Redis            redis  `yaml:"redis" json:"redis"`
}

type mysql struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

type redis struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Pwd  string `yaml:"pwd" json:"pwd"`
}

type log struct {
	LogPath          string `yaml:"logPath" json:"logPath"`
	Level            string `yaml:"level" json:"level"`
	EnableLogConsole bool   `yaml:"enableLogConsole" json:"enableLogConsole"`
}
