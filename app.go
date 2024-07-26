package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"wailsdemo/service"
	"wailsdemo/types"
	"wailsdemo/utils"
)

// App struct
type App struct {
	ctx    context.Context
	chat   service.Chat
	status string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	exists := utils.FileExists("app.json")
	if !exists {
		conf := types.Config{}
		jsonData, _ := json.Marshal(conf)
		utils.WriteFile("app.json", jsonData)
		a.status = fmt.Sprintf("应用初始化，请配置:%s%capp.json 配置文件", utils.GetCurrentPath(), os.PathSeparator)
	}
	a.status = fmt.Sprintf("应用启动，成功读取路径:%s%capp.json 配置文件", utils.GetCurrentPath(), os.PathSeparator)
	a.chat = &service.EchoChat{}
	InitWithApp()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetStatus() string {
	return a.status
}

func (a *App) DoChat(msg []types.Message) string {
	return a.chat.Dochat(msg)
}
