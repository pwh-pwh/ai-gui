package service

import "wailsdemo/types"

type Chat interface {
	Dochat([]types.Message) string
}

func GetChatByConfig(conf types.Config) Chat {
	switch conf.BotType {
	case "echo":
		return &EchoChat{}
	case "gpt":
		gptClient, err := NewGptChat(conf.Gpt)
		if err != nil {
			return &ErrChat{errMsg: err.Error()}
		}
		return gptClient
	default:
		return &ErrChat{errMsg: "botType配置不正确"}
	}
}
