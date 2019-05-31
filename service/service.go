package service

import (
	"github.com/Deansquirrel/goEncodeToolsWs/global"
	"github.com/Deansquirrel/goEncodeToolsWs/webServer"
	log "github.com/Deansquirrel/goToolLog"
)

//启动服务内容
func StartService() error {
	log.Debug("StartService")
	go func() {
		ws := webServer.NewWebServer(global.SysConfig.Iris.Port)
		ws.StartWebService()
	}()
	return nil
}
