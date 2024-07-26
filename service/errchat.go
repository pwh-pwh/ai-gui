package service

import "wailsdemo/types"

type ErrChat struct {
	errMsg string
}

func (e *ErrChat) Dochat(messages []types.Message) string {
	return e.errMsg
}
