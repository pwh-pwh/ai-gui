package types

type Config struct {
	BotType string    `json:"botType"`
	Gpt     GptConfig `json:"gpt"`
	Qf      QfConfig  `json:"qf"`
}

type GptConfig struct {
	BaseUrl string `json:"baseUrl"`
	ApiKey  string `json:"apiKey"`
	Model   string `json:"model"`
}

type QfConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Model     string `json:"model"`
}
