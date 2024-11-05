package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "index.tmpl", nil) //http.statusOK == 200
	fmt.Println("主页")
}

func HealthData1(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "HealthData.tmpl", nil) //http.statusOK == 200
}

func DataMap(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "DataMap.tmpl", nil) //http.statusOK == 200
}

func Figure(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "figure.tmpl", nil) //http.statusOK == 200
}

func Wxack(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "wxack.tmpl", nil) //http.statusOK == 200
}
func HomePage(context *gin.Context) {
	//fmt.Println("--------------------------------------")
	context.HTML(http.StatusOK, "homepage.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func Survey(context *gin.Context) {
	context.HTML(http.StatusOK, "Survey.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func Surveys(context *gin.Context) {
	context.HTML(http.StatusOK, "AdminSurvey.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func UserManage(context *gin.Context) {
	context.HTML(http.StatusOK, "UserManage.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func TrendAnalysis(context *gin.Context) {
	context.HTML(http.StatusOK, "TrendAnalysis.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func RiskPrediction(context *gin.Context) {
	context.HTML(http.StatusOK, "RiskPrediction.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func Img(context *gin.Context) {
	context.HTML(http.StatusOK, "img.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func AdminPersonalInformation(context *gin.Context) {
	context.HTML(http.StatusOK, "adminPersonalInformation.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func RiskForecast(context *gin.Context) {
	context.HTML(http.StatusOK, "RiskForecast.tmpl", nil) //http.statusOK == 200
	fmt.Println("")
}

func ImagejpgHandler(context *gin.Context) {
	context.File("D:\\code\\GO\\MentalHealth-Platform\\app\\images\\avatars\\img.jpg")
}
