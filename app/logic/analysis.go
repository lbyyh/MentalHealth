package logic

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/tools"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// SurveysTeenAllList 获取数据库青少年表的的所有信息并整理后传递
func SurveysTeenAllList(c *gin.Context) {
	surveys := model.GetSurveysTeen(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
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
			"remark":          survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysCollegeAllList 获取数据库大学生表的的所有信息并整理后传递
func SurveysCollegeAllList(c *gin.Context) {
	surveys := model.GetSurveysCollege(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，获取所有字段
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":                         survey.Id,
			"grade":                      survey.Grade,
			"gender":                     survey.Gender,
			"birthplace":                 survey.Birthplace,
			"monthlyExpense":             survey.MonthlyExpense,
			"homeTown":                   survey.HomeTown,
			"expectations":               survey.Expectations,
			"singleChild":                survey.SingleChild,
			"futureJobExpectation":       survey.FutureJobExpectation,
			"relationshipWithClassmates": survey.RelationshipWithClassmates,
			"examTasks":                  survey.ExamTasks,
			"abilityToHandle":            survey.AbilityToHandle,
			"careAboutOthers":            survey.CareAboutOthers,
			"selfRequirement":            survey.SelfRequirement,
			"impactByGrade":              survey.ImpactByGrade,
			"impactByGender":             survey.ImpactByGender,
			"writtenBy":                  survey.WrittenBy,
			"creationTime":               survey.CreationTime,
			"remark":                     survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysWorkerAllList 获取数据库社会工作者表的的所有信息并整理后传递
func SurveysWorkerAllList(c *gin.Context) {
	surveys := model.GetSurveysWorker(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，获取所有字段
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":                survey.Id,
			"workUnit":          survey.WorkUnit,
			"gender":            survey.Gender,
			"age":               survey.Age,
			"maritalStatus":     survey.MaritalStatus,
			"educationLevel":    survey.EducationLevel,
			"professionalTitle": survey.ProfessionalTitle,
			"department":        survey.Department,
			"workYears":         survey.WorkYears,
			"enterpriseNature":  survey.EnterpriseNature,
			"writtenBy":         survey.WrittenBy,
			"creationTime":      survey.CreationTime,
			"remark":            survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysTeenAllListByGender 根据性别获取数据库青少年表的的所有信息并整理后传递
func SurveysTeenAllListByGender(c *gin.Context) {
	surveys := model.GetSurveysTeenByGender(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
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
			"remark":          survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysCollegeAllListByGender 根据性别获取数据库大学生表的的所有信息并整理后传递
func SurveysCollegeAllListByGender(c *gin.Context) {
	surveys := model.GetSurveysCollegeByGender(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":                         survey.Id,
			"grade":                      survey.Grade,
			"gender":                     survey.Gender,
			"birthplace":                 survey.Birthplace,
			"monthlyExpense":             survey.MonthlyExpense,
			"homeTown":                   survey.HomeTown,
			"expectations":               survey.Expectations,
			"singleChild":                survey.SingleChild,
			"futureJobExpectation":       survey.FutureJobExpectation,
			"relationshipWithClassmates": survey.RelationshipWithClassmates,
			"examTasks":                  survey.ExamTasks,
			"abilityToHandle":            survey.AbilityToHandle,
			"careAboutOthers":            survey.CareAboutOthers,
			"selfRequirement":            survey.SelfRequirement,
			"impactByGrade":              survey.ImpactByGrade,
			"impactByGender":             survey.ImpactByGender,
			"writtenBy":                  survey.WrittenBy,
			"creationTime":               survey.CreationTime,
			"remark":                     survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

// SurveysWorkerAllListByGender 根据性别获取数据库社会工作者表的的所有信息并整理后传递
func SurveysWorkerAllListByGender(c *gin.Context) {
	surveys := model.GetSurveysWorkerByGender(c)

	// 创建一个新的切片来存储处理后的数据
	var processedSurveys []map[string]interface{}

	// 遍历获取到的调查问卷数据，仅获取'id'和'WrittenBy'
	for _, survey := range surveys {
		processedSurvey := map[string]interface{}{
			"id":                survey.Id,
			"workUnit":          survey.WorkUnit,
			"gender":            survey.Gender,
			"age":               survey.Age,
			"maritalStatus":     survey.MaritalStatus,
			"educationLevel":    survey.EducationLevel,
			"professionalTitle": survey.ProfessionalTitle,
			"department":        survey.Department,
			"workYears":         survey.WorkYears,
			"enterpriseNature":  survey.EnterpriseNature,
			"writtenBy":         survey.WrittenBy,
			"creationTime":      survey.CreationTime,
			"remark":            survey.Remark,
		}
		processedSurveys = append(processedSurveys, processedSurvey)
	}

	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, processedSurveys)
}

type AllFind struct {
	Id           string    `json:"id"`
	WrittenBy    string    `json:"WrittenBy"`
	CreationTime time.Time `gorm:"column:Creation_time;default:NULL"`
	Remark       string    `json:"remark"`

	// 增加一个用来识别问卷种类的字段
	SurveyType string `json:"survey_type"`
}

// SurveysFind 根据用户id获取所有问卷表的数据
func SurveysFind(c *gin.Context) {
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
		fmt.Println("无效的token")
		return
	}
	id := userData.Id
	fmt.Println("id----------------", id)
	if id == 0 {
		// 将处理后的错误消息返回到前端
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var result []AllFind // 用来保存最终结果的数组

	// 在查询时增加一个条件
	var teenSurveys []AllFind // 存储青少年问卷结果的数组
	if err := model.MySQL.Table("teen_survey").Where("WrittenBy = ?", id).Find(&teenSurveys).Error; err != nil {
		fmt.Printf("Retrieve teen_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 设置问卷种类为"teen_survey"
	for i := range teenSurveys {
		teenSurveys[i].SurveyType = "青少年"
	}
	result = append(result, teenSurveys...) // 将结果添加到最终结果数组

	// 在查询时增加一个条件
	var collegeSurveys []AllFind // 存储大学生问卷结果的数组
	if err := model.MySQL.Table("college_survey").Where("WrittenBy = ?", id).Find(&collegeSurveys).Error; err != nil {
		fmt.Printf("Retrieve college_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 设置问卷种类为"college_survey"
	for i := range collegeSurveys {
		collegeSurveys[i].SurveyType = "大学生"
	}
	result = append(result, collegeSurveys...) // 将结果添加到最终结果数组
	var workerSurveys []AllFind                // 存储工人问卷结果的数组
	if err := model.MySQL.Table("worker_survey").Where("WrittenBy = ?", id).Find(&workerSurveys).Error; err != nil {
		fmt.Printf("Retrieve worker_survey error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 设置问卷种类为"worker_survey"
	for i := range workerSurveys {
		workerSurveys[i].SurveyType = "社会工作者"
	}
	result = append(result, workerSurveys...) // 将结果添加到最终结果数组
	// 将处理后的数据返回到前端
	c.JSON(http.StatusOK, result)
}

// CsvData 处理器函数
func CsvData(c *gin.Context) {

	// 打开CSV文件
	file, err := os.Open("D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\Mental Health Dataset.csv") // 请替换为实际的文件路径
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open CSV file"})
		return
	}
	defer file.Close()

	// 读取CSV文件
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read CSV file"})
		return
	}

	// 解析CSV文件
	var data []map[string]string
	headers := lines[0]              // header row
	for _, line := range lines[1:] { // skip header row
		row := make(map[string]string)
		for i, header := range headers {
			row[header] = line[i]
		}
		data = append(data, row)
	}

	// 将数据转换为JSON格式并返回
	c.JSON(http.StatusOK, data)
}

// 健康记录的结构
//type HealthRecord struct {
//	// 这里根据你的csv文件设置字段
//}

func GetPaginatedRecords(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "20")

	// 转换成 int类型
	intPage, _ := strconv.Atoi(page)
	intRowsPerPage, _ := strconv.Atoi(rowsPerPage)

	file, err := os.Open("D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\Mental Health Dataset.csv") // 请替换为实际的文件路径
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open CSV file"})
		return
	}
	defer file.Close()

	// 读取CSV文件
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read CSV file"})
		return
	}

	// 解析CSV文件
	var data []map[string]string
	headers := lines[0]              // 第一行是header
	for _, line := range lines[1:] { // 不包括header
		row := make(map[string]string)
		for i, header := range headers {
			row[header] = line[i]
		}
		data = append(data, row)
	}

	// 根据页码和每页行数进行切片
	startIndex := (intPage - 1) * intRowsPerPage
	endIndex := intPage * intRowsPerPage
	if endIndex > len(data) {
		endIndex = len(data)
	}

	c.JSON(http.StatusOK, data[startIndex:endIndex])
}

// Record 结构体，此处将 CSV 文件中的每一行定义为一个 Record
type Record struct {
	Country string `json:"Country"`
}

// GetCountries 获取csv中的国家信息
func GetCountries(c *gin.Context) {
	// 打开CSV文件
	file, err := os.Open("D:\\code\\GO\\MentalHealth-Platform\\app\\pymodel\\Mental Health Dataset.csv") // 请替换为实际的文件路径
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 创建一个新的 CSV reader
	reader := csv.NewReader(bufio.NewReader(file))

	var out []Record
	_, _ = reader.Read() // 将 CSV 文件的第一行(header)抛掉
	for {
		row, err := reader.Read() // 读取每一行
		if err != nil {
			break
		}
		country := strings.TrimSpace(row[2]) // 第三列为 "Country"
		out = append(out, Record{
			Country: country,
		})
	}
	c.JSON(http.StatusOK, out)
}

type Location struct {
	Country   string  `json:"Country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetCountriesLatLong(c *gin.Context) {
	// 用你提供的经纬度创建一个Location数组
	locations := []Location{
		{Country: "United States", Latitude: 38.9072, Longitude: -77.0369},
		{Country: "United Kingdom", Latitude: 51.5074, Longitude: -0.1278},
		{Country: "Canada", Latitude: 45.4215, Longitude: -75.6972},
		{Country: "Australia", Latitude: -35.2809, Longitude: 149.1300},
		{Country: "Sweden", Latitude: 59.3293, Longitude: 18.0686},
		{Country: "Ireland", Latitude: 53.3498, Longitude: -6.2603},
		{Country: "Poland", Latitude: 52.2297, Longitude: 21.0122},
		{Country: "South Africa", Latitude: -25.7449, Longitude: 28.1878},
		{Country: "New Zealand", Latitude: -41.2866, Longitude: 174.7756},
		{Country: "Netherlands", Latitude: 52.3702, Longitude: 4.8952},
		{Country: "India", Latitude: 28.6139, Longitude: 77.2090},
		{Country: "Belgium", Latitude: 50.8503, Longitude: 4.3517},
	}

	locationsJSON, err := json.Marshal(locations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, string(locationsJSON))
}
