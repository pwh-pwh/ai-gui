package service

import "github.com/pwh-pwh/ai-gui/types"

type EchoChat struct {
}

func (e *EchoChat) Dochat(messages []types.Message) string {
	return messages[len(messages)-1].Content
}
