package service

import (
	"context"
	"errors"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"github.com/pwh-pwh/ai-gui/types"
)

type QfChat struct {
	accessKey string
	secretKey string
	model     string
}

func NewQfChat(conf types.QfConfig) (*QfChat, error) {
	if conf.AccessKey == "" || conf.SecretKey == "" {
		return nil, errors.New("accessKey或secretKey为空")
	}
	qfClient := &QfChat{
		accessKey: conf.AccessKey,
		secretKey: conf.SecretKey,
		model:     conf.Model,
	}
	return qfClient, nil
}

func (c *QfChat) Dochat(messages []types.Message) string {
	qianfan.GetConfig().AccessKey = c.accessKey
	qianfan.GetConfig().SecretKey = c.secretKey
	chat := qianfan.NewChatCompletion(
		qianfan.WithModel(c.model))
	var msgList []qianfan.ChatCompletionMessage
	for _, message := range messages {
		msgList = append(msgList, qianfan.ChatCompletionMessage{
			Role:    message.UserType,
			Content: message.Content,
		})
	}
	resp, err := chat.Do(context.Background(), &qianfan.ChatCompletionRequest{
		Messages: msgList,
	})
	if err != nil {
		return err.Error()
	}
	return resp.Result
}
