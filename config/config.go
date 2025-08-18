package serviceconfig

type AppConfig struct {
	BnBotFt BnBotFt `mapstructure:"bn-bot-ft" json:"bn-bot-ft"`
	Queue   Queue   `mapstructure:"queue" json:"queue"`
	Server  Server  `mapstructure:"server" json:"server"`
}

type BnBotFt struct {
	BaseUrl               string `mapstructure:"baseUrl" json:"baseUrl"`
	NewOrderEndpoint      string `mapstructure:"newOrderEndpoint" json:"newOrderEndpoint"`
	ActivateBotEndpoint   string `mapstructure:"activateBotEndpoint" json:"activateBotEndpoint"`
	DeactivateBotEndpoint string `mapstructure:"deactivateBotEndpoint" json:"deactivateBotEndpoint"`
	TvActivationEndpoint  string `mapstructure:"tvActivationEndpoint" json:"tvActivationEndpoint"`
}

type Queue struct {
	Size int `mapstructure:"size" json:"size"`
}

type Server struct {
	Port int    `mapstructure:"port" json:"port"`
	Env  string `mapstructure:"env" json:"env"`
}
