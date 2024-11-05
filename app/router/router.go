package router

import (
	"MentalHealth-Platform/app/logic"
	"MentalHealth-Platform/app/middleware"
	"MentalHealth-Platform/app/model"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("app/view/*")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//初始登陆界面
	index := r.Group("")
	//静态页面
	{
		index.GET("/index", logic.Index)                                       //静态页面
		index.GET("/HealthData", logic.HealthData1)                            //静态页面
		index.GET("/DataMap", logic.DataMap)                                   //静态页面
		index.GET("/figure", logic.Figure)                                     //静态页面
		index.GET("/wxack", logic.Wxack)                                       //微信扫码登陆手机端确认登陆页面
		index.GET("/homepage", logic.HomePage)                                 //主页
		index.GET("/survey", logic.Survey)                                     //用户问卷界面
		index.GET("/surveys", logic.Surveys)                                   //用户问卷界面
		index.GET("/UserManage", logic.UserManage)                             //用户管理界面
		index.GET("/TrendAnalysis", logic.TrendAnalysis)                       //数据趋势分析界面
		index.GET("/HorizontalPrediction", logic.RiskPrediction)               //健康预测界面
		index.GET("/Img", logic.Img)                                           //主页海报界面
		index.GET("/app/images/avatars/img.jpg", logic.ImagejpgHandler)        //海报图片文件地址
		index.GET("/adminPersonalInformation", logic.AdminPersonalInformation) //管理员个人信息界面
		index.GET("/RiskForecast", logic.RiskForecast)                         //预测界面

	}

	r.Static("/images", "app/images")
	// 上传接口
	r.StaticFile("/favicon.ico", "./app/images/images/图书(1).png") //
	//css样式
	r.Static("/app/css", "./app/css")
	{
		//用户user
		user := r.Group("/user")
		user.POST("/uploadAvatar", logic.UploadAvatar)
		user.POST("/login", logic.Login)
		user.POST("/email-login", logic.EmailLogin)
		user.POST("/update-login-status", logic.Updateloginstatus)
		user.GET("/GetToken", model.GetToken)
		user.POST("/SendEmailCaptcha", logic.SendEmailCaptcha)
		user.POST("/SendSMSCaptcha", logic.SendSMSCaptcha)
		user.POST("/VerifySMSCaptcha", logic.VerifySMSCaptcha)
		user.GET("/wechat", logic.CheckSignature)
		user.GET("/wechat/login", logic.Redirect)
		user.GET("/wechat/Callback", logic.Callback)
		user.GET("/wechat/check_login", logic.CheckLogin)
		user.Use(middleware.CheckUser)

		user.GET("/UserLogin", logic.UserLogin)                      //用户界面
		user.POST("/submitTeenSurvey", logic.SubmitTeenSurvey)       //提交青少年问卷表单
		user.POST("/submitCollegeSurvey", logic.SubmitCollegeSurvey) //提交大学生问卷表单
		user.POST("/submitWorkerSurvey", logic.SubmitWorkerSurvey)   //提交社会工作者问卷表单
		user.GET("/SurveysFind", logic.SurveysFind)                  //根据用户id查找所有问卷
		user.GET("/csvData", logic.GetPaginatedRecords)              //传递csv数据
		user.GET("/getCountries", logic.GetCountries)                //传递csv国家数据
		user.GET("/getCountriesLatLong", logic.GetCountriesLatLong)  //传递csv国家经纬度数据
		user.POST("/countColumns", logic.CountColumns)               //传递csv数据
	}
	{
		//问卷获取
		surveys := r.Group("/surveys")
		surveys.GET("/SurveysTeenList", logic.SurveysTeenList)                             //获取青少年问卷列表
		surveys.GET("/SurveysCollegeList", logic.SurveysCollegeList)                       //获取大学生问卷列表
		surveys.GET("/SurveysWorkerList", logic.SurveysWorkerList)                         //获取社会工作者问卷列表
		surveys.GET("/SurveysTeenAllList", logic.SurveysTeenAllList)                       //获取青少年表所有信息
		surveys.GET("/SurveysCollegeAllList", logic.SurveysCollegeAllList)                 //获取大学生表所有信息
		surveys.GET("/SurveysWorkerAllList", logic.SurveysWorkerAllList)                   //获取社会工作者表所有信息
		surveys.GET("/SurveysTeenAllListByGender", logic.SurveysTeenAllListByGender)       //根据性别获取青少年表所有信息
		surveys.GET("/SurveysCollegeAllListByGender", logic.SurveysCollegeAllListByGender) //根据性别获取大学生表所有信息
		surveys.GET("/SurveysWorkerAllListByGender", logic.SurveysWorkerAllListByGender)   //根据性别获取社会工作者表所有信息
		surveys.GET("/teenContent/:id", logic.SurveysTeenContent)                          //根据青少年问卷id获取内容
		surveys.GET("/collegeContent/:id", logic.SurveysCollegeContent)                    //根据大学生问卷id获取内容
		surveys.GET("/workerContent/:id", logic.SurveysWorkerContent)                      //根据社会工作者问卷id获取内容
		surveys.DELETE("/teenDelete/:id", model.DeleteSurveyTeenByID)                      //根据社会工作者问卷id删除内容
		surveys.DELETE("/collegeDelete/:id", model.DeleteSurveyCollegeByID)                //根据社会工作者问卷id删除内容
		surveys.DELETE("/workerDelete/:id", model.DeleteSurveyWorkerByID)                  //根据社会工作者问卷id删除内容
		surveys.POST("/teenRemarkAdd/:id", model.TeenRemarkAdd)                            //根据社会工作者id添加问卷备注
		surveys.POST("/collegeRemarkAdd/:id", model.CollegeRemarkAdd)                      //根据社会工作者id添加问卷备注
		surveys.POST("/workerRemarkAdd/:id", model.WorkerRemarkAdd)                        //根据社会工作者id添加问卷备注

	}
	{
		//管理员admin
		admin := r.Group("/admin")
		admin.POST("/login", logic.AdminLogin)
		admin.Use(middleware.CheckAdmin)
		admin.GET("/logout", logic.AdminLogout)
		admin.GET("/AdminLogin", logic.AdminLoginS)                               //管理员界面
		admin.GET("/surveys", logic.AdminSurveys)                                 //管理员界面
		admin.GET("/GetUsersList", logic.GetUsersList)                            //管理员界面
		admin.POST("/UpdateUser", model.UpdateUser)                               //更新用户
		admin.POST("/DeleteUser", model.DeleteUser)                               //删除用户
		admin.POST("/SuspendUser", model.SuspendUser)                             //挂起用户
		admin.POST("/UnSuspendUser", model.UnSuspendUser)                         //取消挂起用户
		admin.POST("/AddUser", model.AddUser)                                     //添加用户
		admin.GET("/AllPersonalInformation", logic.AllPersonalInformation)        //查询管理员所有信息
		admin.GET("/PersonalInformation", logic.PersonalInformation)              //查询管理员个人信息
		admin.POST("/UpdatePersonalInformation", logic.UpdatePersonalInformation) //根据id修改管理员个人信息
		admin.POST("/UpdatePassword", logic.UpdatePassWord)                       //根据id修改管理员个人信息
	}
	{
		//健康预测
		predict := r.Group("/pre")
		predict.POST("/predict", logic.Predict1)
		predict.GET("/teenPredict", logic.TeenPredict)
		predict.GET("/collegePredict", logic.CollegePredict)
		predict.GET("/workerPredict", logic.WorkerPredict)
		predict.GET("/TeenRiskForecast/:id", logic.TeenRiskForecast)
		predict.GET("/CollegeRiskForecast/:id", logic.CollegeRiskForecast)
		predict.GET("/WorkerRiskForecast/:id", logic.WorkerRiskForecast)
	}
	{
		//图片验证码
		r.GET("/captcha", logic.Captcha)
		r.POST("/captcha/verify", logic.CaptchaVerify)
	}

	if err := r.Run(":8087"); err != nil {
		panic(err)
	}
}
