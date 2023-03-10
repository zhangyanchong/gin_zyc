package config

type App struct {
	AppName  string      `json:"app_name"`
	AppModel string      `json:"app_mode"`
	AppPort  string      `json:"app_port"`
	Mysql    MysqlConfig `json:"mysql"`
}

type MysqlConfig struct {
	Ip       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}
