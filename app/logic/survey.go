package logic

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// SurveyData 用于解析前端JSON数据
type SurveyData struct {
	Gender          string `json:"gender"`
	Character       string `json:"character"`
	Hobby           string `json:"hobby"`
	ParentsDivorced string `json:"parents_divorced"`
	Loneliness      string `json:"loneliness"`
	EarlyLoveImpact string `json:"early_love_impact"`
	Alcohol         string `json:"alcohol"`
	ProblemSolving  string `json:"problem_solving"`
	Failure         string `json:"failure"`
	EducationImpact string `json:"education_impact"`
	FutureJob       string `json:"future_job"`
	WrittenBy       string `json:"written_by"`
}

// SubmitSurvey 处理提交问卷的请求
func SubmitTeenSurvey(c *gin.Context) {
	// 获取用户信息
	cookie, err := c.Request.Cookie("token")
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}
	jwt := cookie.Value
	userData, err := model.CheckJwt(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		return
	}
	userid := userData.Id

	var data SurveyData

	// 解析前端发送的JSON数据
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求数据失败。处理步骤: " + err.Error()})
		return
	}

	// 插入数据库
	_, err = model.AddTeenSurvey(data.Gender, data.Character, data.Hobby, data.ParentsDivorced, data.Loneliness, data.EarlyLoveImpact,
		data.Alcohol, data.ProblemSolving, data.Failure, data.EducationImpact, data.FutureJob, strconv.FormatInt(userid, 10), time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储调查数据失败。处理步骤: " + err.Error()})
		return
	}

	// 成功返回
	c.JSON(http.StatusOK, gin.H{"status": "成功提交调查报告!"})
}

// CollegeSurveyData 用于解析从前端传过来的JSON数据
type CollegeSurveyData struct {
	Grade                      string `json:"grade"`
	Gender                     string `json:"gender"`
	Birthplace                 string `json:"birthplace"`
	MonthlyExpense             string `json:"monthly_expense"`
	HomeTown                   string `json:"home_town"`
	Expectations               string `json:"expectations"`
	SingleChild                string `json:"single_child"`
	FutureJobExpectation       string `json:"future_job_expectation"`
	RelationshipWithClassmates string `json:"relationship_with_classmates"`
	ExamTasks                  string `json:"exam_tasks"`
	AbilityToHandle            string `json:"ability_to_handle"`
	CareAboutOthers            string `json:"care_about_others"`
	SelfRequirement            string `json:"self_requirement"`
	ImpactByGrade              string `json:"impact_by_grade"`
	ImpactByGender             string `json:"impact_by_gender"`
	WrittenBy                  string `json:"written_by"`
}

// SubmitCollegeSurvey 处理提交大学生问卷的请求
func SubmitCollegeSurvey(c *gin.Context) {
	// 获取用户信息
	cookie, err := c.Request.Cookie("token")
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}
	jwt := cookie.Value
	userData, err := model.CheckJwt(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		return
	}
	userid := userData.Id

	var data CollegeSurveyData

	// 解析前端发送的JSON数据
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求数据失败。处理步骤: " + err.Error()})
		return
	}

	// 插入数据库
	_, err = model.AddCollegeSurvey(data.Grade, data.Gender, data.Birthplace, data.MonthlyExpense, data.HomeTown, data.Expectations, data.SingleChild,
		data.FutureJobExpectation, data.RelationshipWithClassmates, data.ExamTasks, data.AbilityToHandle, data.CareAboutOthers, data.SelfRequirement,
		data.ImpactByGrade, data.ImpactByGender, strconv.FormatInt(userid, 10), time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储调查数据失败。处理步骤: " + err.Error()})
		return
	}

	// 成功返回
	c.JSON(http.StatusOK, gin.H{"status": "成功提交调查报告!"})
}

// WorkerSurveyData 用于解析从前端发来的JSON数据
type WorkerSurveyData struct {
	WorkUnit          string `json:"work_unit"`
	Gender            string `json:"gender"`
	Age               string `json:"age"`
	MaritalStatus     string `json:"marital_status"`
	EducationLevel    string `json:"education_level"`
	ProfessionalTitle string `json:"professional_title"`
	Department        string `json:"department"`
	WorkYears         string `json:"work_years"`
	EnterpriseNature  string `json:"enterprise_nature"`
	WrittenBy         string `json:"written_by"`
}

// SubmitWorkerSurvey 处理提交社会工作者问卷的请求
func SubmitWorkerSurvey(c *gin.Context) {
	// 获取用户信息
	cookie, err := c.Request.Cookie("token")
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}
	jwt := cookie.Value
	userData, err := model.CheckJwt(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		return
	}
	userid := userData.Id

	var data WorkerSurveyData

	//解析前端提交的 JSON 数据
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求数据失败。处理步骤: " + err.Error()})
		return
	}

	//将解析后的数据保存至数据库
	_, err = model.AddWorkerSurvey(data.WorkUnit, data.Gender, data.Age, data.MaritalStatus, data.EducationLevel, data.ProfessionalTitle, data.Department, data.WorkYears, data.EnterpriseNature, strconv.FormatInt(userid, 10), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "存储调查数据失败。处理步骤: " + err.Error()})
		return
	}

	//成功返回
	c.JSON(http.StatusOK, gin.H{"status": "成功提交调查报告!"})
}

// SurveysTeenList 处理提交青少年问卷的请求
func SurveysTeenList(c *gin.Context) {
	surveys := model.GetSurveysTeen(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":           survey.Id,
			"WrittenBy":    survey.WrittenBy,
			"CreationTime": survey.CreationTime,
			"Remark":       survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysCollegeList 处理提交大学生问卷的请求
func SurveysCollegeList(c *gin.Context) {
	surveys := model.GetSurveysCollege(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":           survey.Id,
			"WrittenBy":    survey.WrittenBy,
			"CreationTime": survey.CreationTime,
			"Remark":       survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysWorkerList 处理社会工作者问卷的请求
func SurveysWorkerList(c *gin.Context) {
	surveys := model.GetSurveysWorker(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'和'CreationTime'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":           survey.Id,
			"WrittenBy":    survey.WrittenBy,
			"CreationTime": survey.CreationTime,
			"Remark":       survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

func SurveysTeenContent(c *gin.Context) {
	result := model.GetSurveyTeenByID(c)
	fmt.Println("result:", result)
	survey, ok := result.(*model.TeenSurvey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No survey found"})
		return
	}

	// 创建一个新的映射来存储处理后的数据
	var processedSurvey = map[string]interface{}{
		"id":              survey.Id,
		"gender":          survey.Gender,
		"character":       survey.Character,
		"hobby":           survey.Hobby,
		"parentsDivorced": survey.ParentsDivorced,
		"loneliness":      survey.Loneliness,
		"earlyLoveImpact": survey.EarlyLoveImpact,
		"alcohol":         survey.Alcohol,
		"problemSolving":  survey.ProblemSolving,
		"failure":         survey.Failure,
		"educationImpact": survey.EducationImpact,
		"futureJob":       survey.FutureJob,
		"writtenBy":       survey.WrittenBy,
		"creationTime":    survey.CreationTime,
		"Remark":          survey.Remark,
	}
	fmt.Println("processedSurvey------------:", processedSurvey)
	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurvey)
}

func SurveysCollegeContent(c *gin.Context) {
	result := model.GetSurveyCollegeByID(c)
	survey, ok := result.(*model.CollegeSurvey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No survey found"})
		return
	}

	// 创建一个新的映射来存储处理后的数据
	var processedSurvey = map[string]interface{}{
		"id":                           survey.Id,
		"grade":                        survey.Grade,
		"gender":                       survey.Gender,
		"birthplace":                   survey.Birthplace,
		"monthly_expense":              survey.MonthlyExpense,
		"home_town":                    survey.HomeTown,
		"expectations":                 survey.Expectations,
		"single_child":                 survey.SingleChild,
		"future_job_expectation":       survey.FutureJobExpectation,
		"relationship_with_classmates": survey.RelationshipWithClassmates,
		"exam_tasks":                   survey.ExamTasks,
		"ability_to_handle":            survey.AbilityToHandle,
		"care_about_others":            survey.CareAboutOthers,
		"self_requirement":             survey.SelfRequirement,
		"impact_by_grade":              survey.ImpactByGrade,
		"impact_by_gender":             survey.ImpactByGender,
		"writtenBy":                    survey.WrittenBy,
		"creationTime":                 survey.CreationTime,
		"Remark":                       survey.Remark,
	}
	fmt.Println("processedSurvey------------:", processedSurvey)
	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurvey)
}

func SurveysWorkerContent(c *gin.Context) {
	result := model.GetSurveyWorkerByID(c)
	survey, ok := result.(*model.WorkerSurvey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No survey found"})
		return
	}

	// 创建一个新的映射来存储处理后的数据
	var processedSurvey = map[string]interface{}{
		"id":                 survey.Id,
		"work_unit":          survey.WorkUnit,
		"gender":             survey.Gender,
		"age":                survey.Age,
		"marital_status":     survey.MaritalStatus,
		"education_level":    survey.EducationLevel,
		"professional_title": survey.ProfessionalTitle,
		"department":         survey.Department,
		"work_years":         survey.WorkYears,
		"enterprise_nature":  survey.EnterpriseNature,
		"writtenBy":          survey.WrittenBy,
		"creationTime":       survey.CreationTime,
		"Remark":             survey.Remark,
	}
	fmt.Println("processedSurvey------------:", processedSurvey)
	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurvey)
}
