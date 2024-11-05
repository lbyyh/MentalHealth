package main

import (
	"MentalHealth-Platform/app"
)

func main() {
	app.Start()
}

//type TeenSurvey struct {
//	Id              int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
//	Gender          string    `gorm:"column:gender;NOT NULL"`
//	Character       string    `gorm:"column:character;NOT NULL"`
//	Hobby           string    `gorm:"column:hobby;NOT NULL"`
//	ParentsDivorced string    `gorm:"column:parents_divorced;NOT NULL"`
//	Loneliness      string    `gorm:"column:loneliness;NOT NULL"`
//	EarlyLoveImpact string    `gorm:"column:early_love_impact;NOT NULL"`
//	Alcohol         string    `gorm:"column:alcohol;NOT NULL"`
//	ProblemSolving  string    `gorm:"column:problem_solving;NOT NULL"`
//	Failure         string    `gorm:"column:failure;NOT NULL"`
//	EducationImpact string    `gorm:"column:education_impact;NOT NULL"`
//	FutureJob       string    `gorm:"column:future_job;NOT NULL"`
//	WrittenBy       string    `gorm:"column:WrittenBy;NOT NULL"`
//	CreationTime    time.Time `gorm:"column:Creation_time;default:NULL"`
//	Remark          string    `gorm:"column:remark;default:NULL"`
//}
//
//var questionArr = [][]string{
//	{"单位A", "单位B", "单位C", "单位D", "单位E"}, // 用具体的工作单位名称替换默认值
//	{"男", "女"},
//	{"25岁及以下", "26-30", "31-40", "40-50", "51岁及以上"},
//	{"单身", "已婚"},
//	{"中专及以下", "大专", "本科", "硕士及以上"},
//	{"无职称", "初级", "中级", "副高级", "正高级"},
//	{"技术研发部", "生产部", "质检部", "工程部", "其他"},
//	{"1年及以下", "2-3年", "4-6年", "7-10年", "11年及以上", "21年及以上"},
//	{"国有企业", "集体所有制企业", "私营企业", "股份制企业", "联营企业", "外商投资企业", "港、澳、台投资企业", "股份合作企业"},
//}
//
//func main() {
//	dsn := "health:dLE3xr43k4Pak22i@tcp(192.168.30.30:3306)/health?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rand.Seed(time.Now().UnixNano()) // Initialize global pseudo random generator
//
//	for i := 0; i < 100; i++ {
//		insertData(db)
//	}
//}
//func insertData(db *gorm.DB) {
//	var survey model.WorkerSurvey
//	survey.WorkUnit = questionArr[0][rand.Intn(len(questionArr[0]))]
//	survey.Gender = questionArr[1][rand.Intn(len(questionArr[1]))]
//	survey.Age = questionArr[2][rand.Intn(len(questionArr[2]))]
//	survey.MaritalStatus = questionArr[3][rand.Intn(len(questionArr[3]))]
//	survey.EducationLevel = questionArr[4][rand.Intn(len(questionArr[4]))]
//	survey.ProfessionalTitle = questionArr[5][rand.Intn(len(questionArr[5]))]
//	survey.Department = questionArr[6][rand.Intn(len(questionArr[6]))]
//	survey.WorkYears = questionArr[7][rand.Intn(len(questionArr[7]))]
//	survey.EnterpriseNature = questionArr[8][rand.Intn(len(questionArr[8]))]
//
//	// 设置 WrittenBy 字段为 1-20的随机数
//	survey.WrittenBy = fmt.Sprintf("%d", rand.Intn(20)+1)
//
//	// 将当前时间作为创建时间
//	survey.CreationTime = time.Now()
//
//	// 若有需要，添加备注
//	survey.Remark = ""
//
//	// 创建记录
//	result := db.Table("worker_survey").Save(&survey)
//	if result.Error != nil {
//		log.Fatal(result.Error)
//	}
//}
