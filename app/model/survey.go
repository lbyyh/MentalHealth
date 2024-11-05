package model

import (
	//"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AddTeenSurvey 青少年表添加数据
func AddTeenSurvey(gender, character, hobby, parentsDivorced, loneliness, earlyLoveImpact, alcohol, problemSolving, failure, educationImpact, futureJob, writtenBy string, creationTime time.Time) (*TeenSurvey, error) {
	// 青少年问卷实例创建
	survey := TeenSurvey{
		Gender:          gender,
		Character:       character,
		Hobby:           hobby,
		ParentsDivorced: parentsDivorced,
		Loneliness:      loneliness,
		EarlyLoveImpact: earlyLoveImpact,
		Alcohol:         alcohol,
		ProblemSolving:  problemSolving,
		Failure:         failure,
		EducationImpact: educationImpact,
		FutureJob:       futureJob,
		WrittenBy:       writtenBy,
		CreationTime:    creationTime,
	}
	// 插入新创建的用户记录到数据库
	if err := MySQL.Table("teen_survey").Create(&survey).Error; err != nil {
		fmt.Printf("Create teen_survey error: %s", err.Error())
		return nil, err
	}
	return &survey, nil
}

// AddCollegeSurvey 大学生表添加数据
func AddCollegeSurvey(grade, gender, birthplace, monthlyExpense, homeTown, expectations, singleChild, futureJobExpectation, relationshipWithClassmates, examTasks, abilityToHandle, careAboutOthers, selfRequirement, impactByGrade, impactByGender, writtenBy string, creationTime time.Time) (*CollegeSurvey, error) {
	// 大学生问卷实例创建
	survey := CollegeSurvey{
		Grade:                      grade,
		Gender:                     gender,
		Birthplace:                 birthplace,
		MonthlyExpense:             monthlyExpense,
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
		WrittenBy:                  writtenBy,
		CreationTime:               creationTime,
	}

	// 插入新创建的问卷记录到数据库
	if err := MySQL.Table("college_survey").Create(&survey).Error; err != nil {
		fmt.Printf("Create college_survey error: %s", err.Error())
		return nil, err
	}
	return &survey, nil
}

// AddWorkerSurvey 社会工作者表添加数据
func AddWorkerSurvey(workUnit, gender, age, maritalStatus, educationLevel, professionalTitle, department, workYears, enterpriseNature, writtenBy string, creationTime time.Time) (*WorkerSurvey, error) {
	// 社会工作者问卷实例创建
	survey := WorkerSurvey{
		WorkUnit:          workUnit,
		Gender:            gender,
		Age:               age,
		MaritalStatus:     maritalStatus,
		EducationLevel:    educationLevel,
		ProfessionalTitle: professionalTitle,
		Department:        department,
		WorkYears:         workYears,
		EnterpriseNature:  enterpriseNature,
		WrittenBy:         writtenBy,
		CreationTime:      creationTime,
	}

	// 插入新创建的问卷记录到数据库
	if err := MySQL.Table("worker_survey").Create(&survey).Error; err != nil {
		fmt.Printf("Create worker_survey error: %s", err.Error())
		return nil, err
	}
	return &survey, nil
}

// GetSurveysTeen 获取青少年表所有数据
func GetSurveysTeen(c *gin.Context) []TeenSurvey {
	var surveys []TeenSurvey
	if err := MySQL.Table("teen_survey").Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve teen_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}

// GetSurveysCollege 获取大学生表所有数据
func GetSurveysCollege(c *gin.Context) []CollegeSurvey {
	var surveys []CollegeSurvey
	if err := MySQL.Table("college_survey").Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve college_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}

// GetSurveysWorker 获取社会工作者表所有数据
func GetSurveysWorker(c *gin.Context) []WorkerSurvey {
	var surveys []WorkerSurvey
	if err := MySQL.Table("worker_survey").Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve worker_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}

// GetSurveyTeenByID 获取指定id的青少年调查数据
func GetSurveyTeenByID(c *gin.Context) interface{} {
	// 获取URL中的ID参数
	id := c.Param("id")
	fmt.Println("id:------------------", id)

	var survey TeenSurvey
	if err := MySQL.Table("teen_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	return &survey
}

// GetSurveyCollegeByID 获取指定id的大学生调查数据
func GetSurveyCollegeByID(c *gin.Context) interface{} {
	// 获取URL中的ID参数
	id := c.Param("id")

	var survey CollegeSurvey
	if err := MySQL.Table("college_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	return &survey
}

// GetSurveyWorkerByID 获取指定id的社会工作者调查数据
func GetSurveyWorkerByID(c *gin.Context) interface{} {
	// 获取URL中的ID参数
	id := c.Param("id")

	var survey WorkerSurvey
	if err := MySQL.Table("worker_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	return &survey
}

// DeleteSurveyTeenByID 删除指定id的青少年调查数据
func DeleteSurveyTeenByID(c *gin.Context) {
	// 获取URL中的ID参数
	id := c.Param("id")

	var survey TeenSurvey
	if err := MySQL.Table("teen_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除指定的记录
	if err := MySQL.Table("teen_survey").Delete(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 成功删除后，返回删除的数据
	c.JSON(http.StatusOK, gin.H{"data": survey})
}

// DeleteSurveyCollegeByID 删除指定id的大学生调查数据
func DeleteSurveyCollegeByID(c *gin.Context) {
	// 获取URL中的ID参数
	id := c.Param("id")

	var survey CollegeSurvey
	if err := MySQL.Table("college_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除指定的记录
	if err := MySQL.Table("college_survey").Delete(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 成功删除后，返回删除的数据
	c.JSON(http.StatusOK, gin.H{"data": survey})
}

// DeleteSurveyWorkerByID 删除指定id的社会工作者调查数据
func DeleteSurveyWorkerByID(c *gin.Context) {
	// 获取URL中的ID参数
	id := c.Param("id")

	var survey WorkerSurvey
	if err := MySQL.Table("worker_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除指定的记录
	if err := MySQL.Table("worker_survey").Delete(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 成功删除后，返回删除的数据
	c.JSON(http.StatusOK, gin.H{"data": survey})
}

// Remark 定义一个结构体接收id 和 remark
type Remark struct {
	Id     string `json:"id"`
	Remark string `json:"remark"`
}

func TeenRemarkAdd(c *gin.Context) {
	var remark_info Remark
	if err := c.ShouldBindJSON(&remark_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MySQL.Table("teen_survey").Where("id = ?", remark_info.Id).Updates(TeenSurvey{Remark: remark_info.Remark}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "成功更新备注"})
}

func CollegeRemarkAdd(c *gin.Context) {
	var remark_info Remark
	if err := c.ShouldBindJSON(&remark_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MySQL.Table("college_survey").Where("id = ?", remark_info.Id).Updates(TeenSurvey{Remark: remark_info.Remark}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "成功更新备注"})
}

func WorkerRemarkAdd(c *gin.Context) {
	var remark_info Remark
	if err := c.ShouldBindJSON(&remark_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MySQL.Table("worker_survey").Where("id = ?", remark_info.Id).Updates(TeenSurvey{Remark: remark_info.Remark}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "成功更新备注"})
}

// GetSurveysTeenByGender 根据性别获取青少年表所有数据
func GetSurveysTeenByGender(c *gin.Context) []TeenSurvey {
	var surveys []TeenSurvey
	gender := c.Query("gender") // 获取查询参数

	// 在查询时增加一个条件
	if err := MySQL.Table("teen_survey").Where("gender = ?", gender).Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve teen_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}

// GetSurveysCollegeByGender 根据性别获取大学生表所有数据
func GetSurveysCollegeByGender(c *gin.Context) []CollegeSurvey {
	var surveys []CollegeSurvey
	gender := c.Query("gender") // 获取查询参数

	if err := MySQL.Table("college_survey").Where("gender = ?", gender).Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve college_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}

// GetSurveysWorkerByGender 根据性别获取社会工作者表所有数据
func GetSurveysWorkerByGender(c *gin.Context) []WorkerSurvey {
	var surveys []WorkerSurvey
	gender := c.Query("gender") // 获取查询参数

	if err := MySQL.Table("worker_survey").Where("gender = ?", gender).Find(&surveys).Error; err != nil {
		fmt.Printf("Retrieve worker_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return surveys
	}
	return surveys
}
