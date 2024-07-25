package service

import "wailsdemo/types"

type Chat interface {
	Dochat([]types.Message) string
}
