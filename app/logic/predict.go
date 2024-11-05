package logic

import (
	"MentalHealth-Platform/app/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type PredictRequest struct {
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	Mood           string `json:"mood"`
	WorkStress     string `json:"workStress"`
	SocialActivity string `json:"socialActivity"`
	SleepQuality   string `json:"sleepQuality"`
}

type PredictResponse struct {
	Result string `json:"result"`
}

// Predict1 简单预测
func Predict1(c *gin.Context) {
	var req PredictRequest
	//var resp PredictResponse

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	fmt.Printf("req%v\n", req)
	// 构建参数切片
	args := []string{
		"D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\predict.py",
		req.Name,
		strconv.Itoa(req.Age),
		req.Gender,
		req.Mood,
		req.WorkStress,
		req.SocialActivity,
		req.SleepQuality,
	}

	// 使用os/exec包执行Python脚本
	cmd := exec.Command("python", args...)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %v, output: %s", err, string(outputBytes))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   string(outputBytes),
		})
		return
	}
	output := string(outputBytes)
	log.Printf("Python Output: %s", output)

	// 直接将字符串输出发送给客户端
	c.JSON(http.StatusOK, gin.H{
		"result": output,
	})
}

// TeenPredict 根据问卷id查询青少年问卷表的数据并进行预测
func TeenPredict(c *gin.Context) {
	var req PredictRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	fmt.Printf("req%v\n", req)
	// 构建参数切片
	args := []string{
		"D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\predict.py",
		req.Name,
		strconv.Itoa(req.Age),
		req.Gender,
		req.Mood,
		req.WorkStress,
		req.SocialActivity,
		req.SleepQuality,
	}

	// 使用os/exec包执行Python脚本
	cmd := exec.Command("python", args...)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %v, output: %s", err, string(outputBytes))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   string(outputBytes),
		})
		return
	}
	output := string(outputBytes)
	log.Printf("Python Output: %s", output)

	// 直接将字符串输出发送给客户端
	c.JSON(http.StatusOK, gin.H{
		"result": output,
	})
}

type PredictRequest1 struct {
	HomeTown                   int `form:"home_town" json:"home_town" binding:"required"`
	Expectations               int `form:"expectations" json:"expectations" binding:"required"`
	SingleChild                int `form:"single_child" json:"single_child" binding:"required"`
	FutureJobExpectation       int `form:"future_job_expectation" json:"future_job_expectation" binding:"required"`
	RelationshipWithClassmates int `form:"relationship_with_classmates" json:"relationship_with_classmates" binding:"required"`
	ExamTasks                  int `form:"exam_tasks" json:"exam_tasks" binding:"required"`
	AbilityToHandle            int `form:"ability_to_handle" json:"ability_to_handle" binding:"required"`
	CareAboutOthers            int `form:"care_about_others" json:"care_about_others" binding:"required"`
	SelfRequirement            int `form:"self_requirement" json:"self_requirement" binding:"required"`
	ImpactByGrade              int `form:"impact_by_grade" json:"impact_by_grade" binding:"required"`
	ImpactByGender             int `form:"impact_by_gender" json:"impact_by_gender" binding:"required"`
}

func CollegePredict(c *gin.Context) {
	// 使用c.Query()获取URL参数
	id, err := strconv.Atoi(c.Query("surveyId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid survey ID"})
		return
	}

	var survey model.CollegeSurvey
	if err := model.MySQL.Table("college_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	homeTown, _ := strconv.Atoi(survey.HomeTown)
	expectations, _ := strconv.Atoi(survey.Expectations)
	singleChild, _ := strconv.Atoi(survey.SingleChild)
	futureJobExpectation, _ := strconv.Atoi(survey.FutureJobExpectation)
	relationshipWithClassmates, _ := strconv.Atoi(survey.RelationshipWithClassmates)
	examTasks, _ := strconv.Atoi(survey.ExamTasks)
	abilityToHandle, _ := strconv.Atoi(survey.AbilityToHandle)
	careAboutOthers, _ := strconv.Atoi(survey.CareAboutOthers)
	selfRequirement, _ := strconv.Atoi(survey.SelfRequirement)
	impactByGrade, _ := strconv.Atoi(survey.ImpactByGrade)
	impactByGender, _ := strconv.Atoi(survey.ImpactByGender)

	req := PredictRequest1{
		HomeTown:                   homeTown,
		Expectations:               expectations,
		SingleChild:                singleChild,
		FutureJobExpectation:       futureJobExpectation,
		RelationshipWithClassmates: relationshipWithClassmates,
		ExamTasks:                  examTasks,
		AbilityToHandle:            abilityToHandle,
		CareAboutOthers:            careAboutOthers,
		SelfRequirement:            selfRequirement,
		ImpactByGrade:              impactByGrade,
		ImpactByGender:             impactByGender,
	}

	fmt.Println("req--------------------------", req)
	// 将预测请求转换成JSON格式
	jsonReq, _ := json.Marshal(req)

	// 构建参数切片
	args := []string{
		"D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\normalpredict.py",
		string(jsonReq),
	}

	// 使用os/exec包执行Python脚本
	cmd := exec.Command("python", args...)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8", "PYTHONPATH=D:\\pycode\\pythonProject\\.venv\\Lib\\site-packages")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %v, output: %s", err, string(outputBytes))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   string(outputBytes),
		})
		return
	}
	output := string(outputBytes)
	log.Printf("Python Output: %s", output)

	// 直接将字符串输出发送给客户端
	c.JSON(http.StatusOK, gin.H{
		"result": output,
	})
}

// WorkerPredict 根据问卷id查询社会工作者表的数据并进行预测
func WorkerPredict(c *gin.Context) {
	var req PredictRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	fmt.Printf("req%v\n", req)
	// 构建参数切片
	args := []string{
		"D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\predict.py",
		req.Name,
		strconv.Itoa(req.Age),
		req.Gender,
		req.Mood,
		req.WorkStress,
		req.SocialActivity,
		req.SleepQuality,
	}

	// 使用os/exec包执行Python脚本
	cmd := exec.Command("python", args...)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to execute command: %v, output: %s", err, string(outputBytes))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   string(outputBytes),
		})
		return
	}
	output := string(outputBytes)
	log.Printf("Python Output: %s", output)

	// 直接将字符串输出发送给客户端
	c.JSON(http.StatusOK, gin.H{
		"result": output,
	})
}

// TeenForecast 根据问卷id预测青少年问卷表的数据并进行预测
func TeenForecast(c *gin.Context) {
	// 使用c.Query()获取URL参数
	id, err := strconv.Atoi(c.Query("surveyId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid survey ID"})
		return
	}

	var survey model.TeenSurvey
	if err := model.MySQL.Table("teen_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将预测结果输出发送给客户端
	c.JSON(http.StatusOK, gin.H{
		"result": survey,
	})
}

// TeenPredict1 是一个用于预测的假设函数
func TeenPredict1(survey model.TeenSurvey) string {
	// 在这里实现模型的逻辑...
	// 根据survey的字段做一些计算，返回预测结果

	return "预测结果"
}
