package router

import (
	"fmt"
	"github.com/Deansquirrel/goEncodeToolsWs/global"
	"github.com/Deansquirrel/goEncodeToolsWs/object"
	"github.com/kataras/iris"
)

import log "github.com/Deansquirrel/goToolLog"

type base struct {
	app *iris.Application
	c   common
}

func NewRouterBase(app *iris.Application) *base {
	return &base{
		app: app,
		c:   common{},
	}
}

func (base *base) AddBase() {
	base.app.Get("/version", base.version)
	base.app.Post("/code", base.code)
}

func (base *base) version(ctx iris.Context) {
	v := object.VersionResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Version: global.Version,
	}
	base.c.WriteResponse(ctx, v)
}

func (base *base) code(ctx iris.Context) {
	var r object.MessageRequest
	err := ctx.ReadJSON(&r)
	if err != nil {
		errMsg := fmt.Sprintf("非法的请求内容,body: %s", base.c.GetRequestBody(ctx))
		log.Error(fmt.Sprintf("获取请求内容时发生错误，err: %s", err.Error()))
		log.Error(errMsg)
		ctx.StatusCode(iris.StatusBadRequest)
		base.c.WriteResponse(ctx, &object.MessageResponse{
			ErrCode: iris.StatusBadRequest,
			ErrMsg:  errMsg,
		})
		return
	}
	switch r.OprType {
	case 1:
		sObj := object.ObjSecret{}
		s, err := sObj.Encrypt(r.RequestText, r.RequestKey)
		if err != nil {
			errMsg := fmt.Sprintf("加密时遇到错误，err: %s", err.Error())
			log.Error(errMsg)
			ctx.StatusCode(iris.StatusInternalServerError)
			base.c.WriteResponse(ctx, &object.MessageResponse{
				ErrCode: iris.StatusInternalServerError,
				ErrMsg:  errMsg,
			})
		}
		base.c.WriteResponse(ctx, &object.MessageResponse{
			ErrCode:      int(object.ErrTypeCodeNoError),
			ErrMsg:       string(object.ErrTypeMsgNoError),
			ResponseText: s,
		})
	case 2:
		sObj := object.ObjSecret{}
		resp := object.MessageResponse{}
		s, err := sObj.Decrypt(r.RequestText, r.RequestKey)
		if err != nil {
			switch err.Error() {
			case "解密失败。（密码非法）":
				resp.ErrCode = int(object.ErrTypeCodeWrongPassWord)
				resp.ErrMsg = string(object.ErrTypeMsgWrongPassWord)
			case "解密失败。（校验错误）":
				resp.ErrCode = int(object.ErrTypeCodeVerificationFailed)
				resp.ErrMsg = string(object.ErrTypeMsgVerificationFailed)
			default:
				errMsg := fmt.Sprintf("解密时遇到错误，err: %s", err.Error())
				resp.ErrCode = iris.StatusInternalServerError
				resp.ErrMsg = errMsg
			}
			log.Error(fmt.Sprintf("解密错误,code: %d,err: %s", resp.ErrCode, resp.ErrMsg))
		} else {
			resp.ErrCode = int(object.ErrTypeCodeNoError)
			resp.ErrMsg = string(object.ErrTypeMsgNoError)
			resp.ResponseText = s
		}
		base.c.WriteResponse(ctx, &resp)
	default:
		log.Error(fmt.Sprintf("非法的操作类型，exp: 1 or 2，act %d", r.OprType))
		ctx.StatusCode(iris.StatusInternalServerError)
		base.c.WriteResponse(ctx, &object.MessageResponse{
			ErrCode: int(object.ErrTypeCodeWrongOprType),
			ErrMsg:  string(object.ErrTypeMsgWrongOprType),
		})
		return
	}
}
