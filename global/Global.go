package global

import (
	"context"
	"github.com/Deansquirrel/goEncodeToolsWs/object"
)

const (
	//PreVersion = "0.0.1 Build20190531"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
