package types

type Config struct {
	BotType string    `json:"botType"`
	Gpt     GptConfig `json:"gpt"`
}

type GptConfig struct {
	BaseUrl string `json:"baseUrl"`
	ApiKey  string `json:"apiKey"`
	Model   string `json:"model"`
}
