package serviceconfig

type AppConfig struct {
	BnBotFt BnBotFt `mapstructure:"bn-bot-ft" json:"bn-bot-ft"`
	Queue   Queue   `mapstructure:"queue" json:"queue"`
	Server  Server  `mapstructure:"server" json:"server"`
}

type BnBotFt struct {
	BaseUrl          string
	NewOrderEndpoint string
}

type Queue struct {
	Size int
}

type Server struct {
	Port int
	Env  string
}
