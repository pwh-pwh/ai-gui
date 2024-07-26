package service

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"wailsdemo/types"
)

type GptChat struct {
	token   string
	baseUrl string
	model   string
	client  *openai.Client
}

func NewGptChat(conf types.GptConfig) (*GptChat, error) {
	if conf.ApiKey == "" {
		return nil, errors.New("apiKey为空")
	}
	baseUrl := "https://api.openai.com/v1"
	if conf.BaseUrl != "" {
		baseUrl = conf.BaseUrl
	}
	model := "gpt-3.5-turbo"
	if conf.Model != "" {
		model = conf.Model
	}
	gptChat := &GptChat{
		token:   conf.ApiKey,
		baseUrl: baseUrl,
		model:   model,
	}
	cfg := openai.DefaultConfig(gptChat.token)
	cfg.BaseURL = gptChat.baseUrl
	client := openai.NewClientWithConfig(cfg)
	gptChat.client = client
	return gptChat, nil
}

func (g *GptChat) Dochat(messages []types.Message) string {
	var msgList []openai.ChatCompletionMessage
	for _, message := range messages {
		msgList = append(msgList, openai.ChatCompletionMessage{
			Role:    message.UserType,
			Content: message.Content,
		})
	}
	req := openai.ChatCompletionRequest{
		Model:    g.model,
		Messages: msgList,
	}
	response, err := g.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return err.Error()
	}
	return response.Choices[0].Message.Content
}
