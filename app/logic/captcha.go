package logic

import (
	"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 验证码
func Captcha(context *gin.Context) {

	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: captcha,
	})
}

// 验证码
func CaptchaVerify(context *gin.Context) {
	var param tools.CaptchaData
	if err := context.ShouldBind(&param); err != nil {
		context.JSON(http.StatusOK, tools.ParamErr)
		return
	}

	fmt.Printf("参数为：%+v", param)
	if !tools.CaptchaVerify(param) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10008,
			Message: "验证失败",
		})
		return
	}
	context.JSON(http.StatusOK, tools.OK)
}
