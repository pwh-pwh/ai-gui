package service

import "github.com/pwh-pwh/ai-gui/types"

type ErrChat struct {
	errMsg string
}

func (e *ErrChat) Dochat(messages []types.Message) string {
	return e.errMsg
}
