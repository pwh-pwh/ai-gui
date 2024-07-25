package service

import "wailsdemo/types"

type EchoChat struct {
}

func (e *EchoChat) Dochat(messages []types.Message) string {
	return messages[len(messages)-1].Content
}
