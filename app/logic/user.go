package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "homepage.tmpl", nil) //http.statusOK == 200
	fmt.Println("用户主页")
}

func AdminLoginS(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "AdminLogin.tmpl", nil) //http.statusOK == 200
	fmt.Println("管理员主页")
}

func VisitorLoginS(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "VisitorLogin.tmpl", nil) //http.statusOK == 200
	fmt.Println("游客主页")
}

func AdminSurveys(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "AdminSurvey.tmpl", nil) //http.statusOK == 200
	fmt.Println("管理员问卷界面")
}
