package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"wailsdemo/service"
	"wailsdemo/types"
	"wailsdemo/utils"
)

const CONFIG_FILE_NAME = "app.json"

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
	InitWithApp()
	utils.InitConfigPath()
	fullFilePath := filepath.Join(utils.GetConfigPath(), CONFIG_FILE_NAME)
	exists := utils.FileExists(fullFilePath)
	conf := types.Config{}
	if !exists {
		jsonData, _ := json.Marshal(conf)
		err := utils.WriteFile(fullFilePath, jsonData)
		if err != nil {
			a.status = fmt.Sprintf("无法在目录:%s%c创建%s 配置文件，请手动创建", utils.GetConfigPath(), os.PathSeparator, CONFIG_FILE_NAME)
		} else {
			a.status = fmt.Sprintf("应用初始化，请配置:%s%c%s 配置文件", utils.GetConfigPath(), os.PathSeparator, CONFIG_FILE_NAME)
		}
		utils.OpenFolder(utils.GetConfigPath())
	} else {
		a.status = fmt.Sprintf("应用启动，成功读取路径:%s%c%s 配置文件", utils.GetConfigPath(), os.PathSeparator, CONFIG_FILE_NAME)
		bytesData, _ := utils.ReadFile(CONFIG_FILE_NAME)
		err := json.Unmarshal(bytesData, &conf)
		if err != nil {
			a.status = fmt.Sprintf("应用启动，读取配置文件失败:%s", err.Error())
		}
	}
	a.chat = service.GetChatByConfig(conf)
}

// 对接 /loadconf
func (a *App) ReloadConf() string {
	conf := types.Config{}
	bytesData, _ := utils.ReadFile(filepath.Join(utils.GetConfigPath(), CONFIG_FILE_NAME))
	err := json.Unmarshal(bytesData, &conf)
	if err != nil {
		a.status = fmt.Sprintf("重载配置文件失败:%s", err.Error())
	}
	a.chat = service.GetChatByConfig(conf)
	return "配置重载完成"
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

// 对接 /opconf
func (a *App) OpenConfigFolder() {
	utils.OpenFolder(utils.GetConfigPath())
}
