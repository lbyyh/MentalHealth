package logic

import (
	"MentalHealth-Platform/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RiskResult struct {
	RiskScore   float64
	RiskLevel   string
	RiskFactors []string
	Suggestions []string
}

func TeenRiskForecast(c *gin.Context) {
	var survey model.TeenSurvey

	// 获取URL中的ID参数
	id := c.Param("id")

	if err := model.MySQL.Table("teen_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := TeenRiskPredict(survey)

	c.JSON(http.StatusOK, gin.H{
		"riskResult": result,
	})
}

func TeenRiskPredict(survey model.TeenSurvey) RiskResult {
	// 计算风险值
	riskScore, riskFactors, suggestions := calculateTeenRiskScore(survey)

	riskLevel := "low"

	// 根据风险值，设置风险等级
	if riskScore >= 150 && riskScore < 200 {
		riskLevel = "medium"
	} else if riskScore >= 200 {
		riskLevel = "high"
	}

	// 返回风险分数和等级
	return RiskResult{
		RiskScore:   riskScore,
		RiskLevel:   riskLevel,
		RiskFactors: riskFactors,
		Suggestions: suggestions,
	}
}

func calculateTeenRiskScore(survey model.TeenSurvey) (float64, []string, []string) {
	var riskScore float64
	var riskFactors []string
	var suggestions []string

	// 为 "gender" 基于性别设定风险值
	if survey.Gender == "男" {
		riskScore += 10
	} else if survey.Gender == "女" {
		riskScore += 20
	}

	// 为 "character" 基于性格设定风险值
	if survey.Character == "内向型" {
		riskScore += 20
		riskFactors = append(riskFactors, "你的内向性格可能会增加你的心理健康风险。")
		suggestions = append(suggestions, "尝试更多地参与社交活动，与朋友、家人分享你的思想和感情。")
	} else if survey.Character == "外向型" {
		riskScore += 10
	} else if survey.Character == "双向型" {
		riskScore += 15
	}

	// 为 "hobby" 基于爱好设定风险值
	if survey.Hobby == "没有什么特别的爱好" {
		riskScore += 20
		riskFactors = append(riskFactors, "没有特别的爱好可能会增加你的心理健康风险。")
		suggestions = append(suggestions, "尝试发现并培养一些兴趣爱好，这样可以帮助你缓解压力，愉快心情。")
	} else if survey.Hobby == "有，但是没坚持下来" {
		riskScore += 30
		riskFactors = append(riskFactors, "未能坚持爱好可能会增加你的心理健康风险。")
		suggestions = append(suggestions, "试着重新拾起你曾经喜欢的爱好，或者找到一些新的兴趣来尝试。")
	} else if survey.Hobby == "有，而且坚持下来了" {
		riskScore += 10
	}

	// 父母离异也可能影响风险值
	if survey.ParentsDivorced == "是" {
		riskScore += 30
		riskFactors = append(riskFactors, "父母离异可能会增加你的心理压力。")
		suggestions = append(suggestions, "这是一个生活中的重大变化，寻求专业咨询可能会有所帮助。")
	} else if survey.ParentsDivorced == "否" {
		riskScore += 10
	}

	// 为 "loneliness" 基于孤独感设定风险值
	if survey.Loneliness == "是" {
		riskScore += 30
		riskFactors = append(riskFactors, "感到孤独可能会对你的心理健康产生影响。")
		suggestions = append(suggestions, "试图增加社交活动，与朋友和家人建立更紧密的联系。如果感到难以处理，可以寻求专业帮助。")
	} else if survey.Loneliness == "否" {
		riskScore += 10
	}

	// 为 "early_love_impact" 基于早恋对学习的影响设定风险值
	if survey.EarlyLoveImpact == "有" {
		riskScore += 20
		riskFactors = append(riskFactors, "早恋可能会对你的学习和生活产生影响。")
		suggestions = append(suggestions, "确保你对自己的生活有全面的控制，并确保任何的人际关系都不会对你的个人发展产生负面影响。")
	} else if survey.EarlyLoveImpact == "无" {
		riskScore += 10
	}

	// 为 "alcohol" 基于喜欢喝酒解愁设定风险值
	if survey.Alcohol == "是" {
		riskScore += 30
		riskFactors = append(riskFactors, "喜欢喝酒解愁可能会对你的心理健康产生影响。")
		suggestions = append(suggestions, "寻找更健康的方式来应对压力，如运动、冥想或与朋友交流。如果需要的话，寻求专业的心理辅导。")
	} else if survey.Alcohol == "否" {
		riskScore += 10
	}

	// 对于问题 "problem_solving"，为解决问题的方式设定风险值
	if survey.ProblemSolving == "选择逃避" {
		riskScore += 20
		riskFactors = append(riskFactors, "逃避问题可能会增加你的心理压力并延迟问题的解决。")
		suggestions = append(suggestions, "试图学习和练习更积极的问题解决策略，例如寻求帮助或者倾诉，以更健康的方式应对压力。")
	} else if survey.ProblemSolving == "自己解决" {
		riskScore += 10
	} else if survey.ProblemSolving == "寻求帮助" {
		riskScore += 5
	}
	// 对于问题 "failure"，为做错事情后的应对设定风险值
	if survey.Failure == "觉得自己好笨" {
		riskScore += 20
		riskFactors = append(riskFactors, "对失败的消极看法可能会对你的心理健康产生影响。")
		suggestions = append(suggestions, "试图以更积极的方式看待失败，看它作为一个学习和成长的机会。")
	} else if survey.Failure == "很后悔，希望时间重来" {
		riskScore += 10
	} else if survey.Failure == "错了就错了，下次避免再犯" {
		riskScore += 5
	}

	// 为 "education_impact" 问题，根据读书对未来影响的理解来设定风险值
	if survey.EducationImpact == "不会" {
		riskScore += 20
		riskFactors = append(riskFactors, "对教育影响的负面认知可能会影响你的心理健康。")
		suggestions = append(suggestions, "尝试建立积极的学习态度，并了解教育的长远影响。")
	} else if survey.EducationImpact == "会" {
		riskScore += 10
	} else if survey.EducationImpact == "某种程度上" {
		riskScore += 15
	}

	// 以 "future_job" 为例，根据未来职业设定风险值
	if survey.FutureJob == "不确定" {
		riskScore += 30
		riskFactors = append(riskFactors, "对未来职业的不确定感可能会增加你的心理压力。")
		suggestions = append(suggestions, "寻求职业咨询，为自己的未来制定明确的职业规划，有助于减轻心理压力。")
	} else if survey.FutureJob == "已确定" {
		riskScore += 15
	} else if survey.FutureJob == "正在探索" {
		riskScore += 20
	}
	return riskScore, riskFactors, suggestions
}

type Question struct {
	// 原始的问题描述
	Description string
	// 问题的可能答案及对应的风险分数
	Choices map[string]int
}

func CollegeRiskForecast(c *gin.Context) {
	var survey model.CollegeSurvey

	// 获取URL中的ID参数
	id := c.Param("id")

	if err := model.MySQL.Table("college_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := CollegeRiskPredict(survey)

	c.JSON(http.StatusOK, gin.H{
		"riskResult": result,
	})
}

func CollegeRiskPredict(survey model.CollegeSurvey) RiskResult {
	// 初始化风险分数
	riskScore := 0
	riskFactors := []string{} // 存储影响风险的因素
	suggestions := []string{} // 存储减少风险的建议

	// 使用 Question 结构体替代原字符串描述
	questions := map[string]Question{
		"Grade":                      {Description: "你的年级", Choices: map[string]int{"大一": 1, "大二": 2, "大三": 3, "大四": 4}},
		"Gender":                     {Description: "你的性别", Choices: map[string]int{"男": 1, "女": 2}},
		"Birthplace":                 {Description: "你的出生地", Choices: map[string]int{"城市": 1, "乡镇": 2, "农村": 3, "山区": 4}},
		"MonthlyExpense":             {Description: "你的每月支出", Choices: map[string]int{"1000元以下": 1, "1000到1500元": 2, "1500到2000元": 3, "2000到3000元": 4, "3000元以上": 5}},
		"HomeTown":                   {Description: "你的生长地及家乡的风俗习惯", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"Expectations":               {Description: "对父母对您的期望过高或过低", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"SingleChild":                {Description: "觉得是否为独生子女", Choices: map[string]int{"否": 1, "是": 5}}, // 假设此题“是”为5，”否“为1
		"FutureJobExpectation":       {Description: "对未来工作的要求水平", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"RelationshipWithClassmates": {Description: "对和同学的关系的好坏", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"ExamTasks":                  {Description: "觉得学校组织的考试及各种学习任务", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"AbilityToHandle":            {Description: "对调节处理各种事情对自己心理影响的能力", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"CareAboutOthers":            {Description: "对是否在乎其他人的看法、想法", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"SelfRequirement":            {Description: "对自己要求的高低", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"ImpactByGrade":              {Description: "对所在年级", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		"ImpactByGender":             {Description: "性别", Choices: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
	}

	answers := map[string]*string{
		"Grade":                      &survey.Grade,
		"Gender":                     &survey.Gender,
		"Birthplace":                 &survey.Birthplace,
		"MonthlyExpense":             &survey.MonthlyExpense,
		"HomeTown":                   &survey.HomeTown,
		"Expectations":               &survey.Expectations,
		"SingleChild":                &survey.SingleChild,
		"FutureJobExpectation":       &survey.FutureJobExpectation,
		"RelationshipWithClassmates": &survey.RelationshipWithClassmates,
		"ExamTasks":                  &survey.ExamTasks,
		"AbilityToHandle":            &survey.AbilityToHandle,
		"CareAboutOthers":            &survey.CareAboutOthers,
		"SelfRequirement":            &survey.SelfRequirement,
		"ImpactByGrade":              &survey.ImpactByGrade,
		"ImpactByGender":             &survey.ImpactByGender,
	}

	for key, _ := range questions {
		value := *answers[key]                              // 获取问题的答案
		if score, ok := questions[key].Choices[value]; ok { // 从 Choices 获取答案的评分
			riskScore += score
			if score >= 5 { // 假设5为高风险阈值
				riskFactors = append(riskFactors, "问题 "+questions[key].Description+" 的影响程度为 "+strconv.Itoa(score))
				suggestions = append(suggestions, "对于问题 "+questions[key].Description+" ，需要特别注意，寻求专业的帮助和建议。")
			}
		}
	}

	// 风险等级判断
	var riskLevel string
	if riskScore < 30 {
		riskLevel = "低风险"
	} else if riskScore < 60 {
		riskLevel = "中等风险"
	} else {
		riskLevel = "高风险"
	}

	// 返回风险结果
	return RiskResult{
		RiskScore:   float64(riskScore),
		RiskLevel:   riskLevel,
		RiskFactors: riskFactors,
		Suggestions: suggestions,
	}
}

func WorkerRiskForecast(c *gin.Context) {
	var survey model.WorkerSurvey

	// 获取URL中的ID参数
	id := c.Param("id")

	if err := model.MySQL.Table("worker_survey").Where("id = ?", id).First(&survey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := WorkerRiskPredict(survey)

	c.JSON(http.StatusOK, gin.H{
		"riskResult": result,
	})
}

func WorkerRiskPredict(survey model.WorkerSurvey) RiskResult {

	// 初始化风险分数
	riskScore := 0
	riskFactors := []string{} // 存储影响风险的因素
	suggestions := []string{} // 存储减少风险的建议

	// 根据 Age 计算风险分数，并提供相关的建议
	if survey.Age == "51岁及以上" {
		riskScore += 10
		riskFactors = append(riskFactors, "年龄于51岁及以上可能增加工作压力。")
		suggestions = append(suggestions, "维持健康生活方式，进行适当的体育活动，保持良好的饮食习惯。")
	}

	// 根据 MaritalStatus 计算风险分数，并提供相关的建议
	if survey.MaritalStatus == "单身" {
		riskScore += 5
		riskFactors = append(riskFactors, "单身可能需要承受更多生活压力。")
		suggestions = append(suggestions, "寻找社区或者兴趣小组参与，扩大社交网络。")
	}

	// 根据 WorkYears 计算风险分数，并提供相关的建议
	if survey.WorkYears == "21年及以上" {
		riskScore += 15
		riskFactors = append(riskFactors, "工作年龄超过21年可能带来的职业疲倦。")
		suggestions = append(suggestions, "定期进行职业技能培训，提高自身实力")
	}
	// 根据 EducationLevel 计算风险分数，并提供相关的建议
	if survey.EducationLevel == "中专及以下" {
		riskScore += 10
		riskFactors = append(riskFactors, "受教育程度较低可能增加工作压力。")
		suggestions = append(suggestions, "考虑进一步的教育或培训机会，提高自己的职业技能。")
	}

	// 根据 ProfessionalTitle 计算风险分数，并提供相关的建议
	if survey.ProfessionalTitle == "无职称" {
		riskScore += 5
		riskFactors = append(riskFactors, "没有专业职称可能限制了你的工作机会。")
		suggestions = append(suggestions, "进一步提高自己的专业技能，考虑获取专业认证或职称。")
	}

	// 根据 Department 计算风险分数，并提供相关的建议
	if survey.Department == "技术研发部" || survey.Department == "生产部" {
		riskScore += 10
		riskFactors = append(riskFactors, "你所在的部门可能带来较大的工作压力。")
		suggestions = append(suggestions, "确保工作和生活的平衡，定期进行身体检查，保持健康的生活方式。")
	}

	// 根据 EnterpriseNature 计算风险分数，并提供相关的建议
	if survey.EnterpriseNature == "股份制企业" {
		riskScore += 5
		riskFactors = append(riskFactors, "股份制企业带来的工作压力可能较大。")
		suggestions = append(suggestions, "学习和应用压力管理技巧，寻找健康的应对压力的方法。")
	}

	// 风险等级判断
	var riskLevel string
	if riskScore < 30 {
		riskLevel = "低风险"
	} else if riskScore < 60 {
		riskLevel = "中等风险"
	} else {
		riskLevel = "高风险"
	}

	// 返回风险结果
	return RiskResult{
		RiskScore:   float64(riskScore),
		RiskLevel:   riskLevel,
		RiskFactors: riskFactors,
		Suggestions: suggestions,
	}
}
