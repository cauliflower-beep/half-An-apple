package controller

import (
	"github.com/gin-gonic/gin"
)

//解析用户模块请求

type MemberController struct{

}

//注册路由
func (mc *MemberController)Router(engine *gin.Context){
	//发送手机验证码
	engine.GET("/api/sendcode",mc.sendSmsCode)
}

func (mc *MemberController)sendSmsCode(ctx *gin.Context){

}

