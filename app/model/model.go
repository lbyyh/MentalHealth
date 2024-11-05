package model

import "time"

type Admin struct {
	AdminId      int32  `gorm:"column:admin_id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Username     string `gorm:"column:username;NOT NULL"`
	Password     string `gorm:"column:password;NOT NULL"`
	WechatId     string `gorm:"column:wechat_id;default:NULL"`
	PhoneNumber  string `gorm:"column:phone_number;default:NULL"`
	EmailAddress string `gorm:"column:email_address;default:NULL"`
}

func (a *Admin) TableName() string {
	return "admin"
}

type BookInfo struct {
	Id                 uint32 `gorm:"column:id;primary_key;NOT NULL;comment:'书的id'"`
	Uid                int64  `gorm:"column:uid;default:NULL;comment:'书的uid'"`
	BookName           string `gorm:"column:book_name;default:NULL;comment:'书名'"`
	Author             string `gorm:"column:author;default:NULL;comment:'作者'"`
	PublishingHouse    string `gorm:"column:publishing_house;default:NULL;comment:'出版社'"`
	Translator         string `gorm:"column:translator;default:NULL;comment:'译者'"`
	Num                int32  `gorm:"column:num;default:NULL;comment:'书的数量'"`
	PublishDate        string `gorm:"column:publish_date;default:NULL;comment:'出版时间'"`
	Pages              int32  `gorm:"column:pages;default:100;comment:'页数'"`
	ISBN               string `gorm:"column:ISBN;default:NULL;comment:'ISBN号码'"`
	Price              string `gorm:"column:price;default:1;comment:'价格'"`
	BriefIntroduction  string `gorm:"column:brief_introduction;default:;comment:'内容简介'"`
	AuthorIntroduction string `gorm:"column:author_introduction;default:;comment:'作者简介'"`
	ImgUrl             string `gorm:"column:img_url;default:NULL;comment:'封面地址'"`
	DelFlg             int32  `gorm:"column:del_flg;default:0;comment:'删除标识'"`
}

func (b *BookInfo) TableName() string {
	return "book_info"
}

// User undefined
//type User struct {
//	ID          int64     `json:"id" gorm:"id"`
//	Name        string    `json:"name" gorm:"name"`
//	Password    string    `json:"password" gorm:"password"`
//	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
//	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
//}
//
//// TableName 表名称
//func (*User) TableName() string {
//	return "user"
//}

type Users struct {
	Id               int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name             string `gorm:"column:name;NOT NULL"`
	Password         string `gorm:"column:password;default:NULL"`
	Age              int32  `gorm:"column:age;default:NULL"`
	Gender           string `gorm:"column:gender;default:NULL"`
	RegistrationDate string `gorm:"column:registration_date;default:NULL"`
	ContactInfo      string `gorm:"column:contact_info;default:NULL"`
	Status           int32  `gorm:"column:status;default:NULL"`
	LastLandingTime  string `gorm:"column:last_landing_time;default:NULL"`
}

func (u *Users) TableName() string {
	return "users"
}

// BookUser undefined
type BookUser struct {
	ID          int64     `json:"id" gorm:"id"`
	UserId      int64     `json:"user_id" gorm:"user_id"`
	BookId      int64     `json:"book_id" gorm:"book_id"`
	Status      int64     `json:"status" gorm:"status"`
	Time        int64     `json:"time" gorm:"time"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*BookUser) TableName() string {
	return "book_user"
}

type MysqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type Config struct {
	Mysql MysqlConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
}

// TeenSurvey 青少年问卷表
type TeenSurvey struct {
	Id              int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Gender          string    `gorm:"column:gender;NOT NULL"`
	Character       string    `gorm:"column:character;NOT NULL"`
	Hobby           string    `gorm:"column:hobby;NOT NULL"`
	ParentsDivorced string    `gorm:"column:parents_divorced;NOT NULL"`
	Loneliness      string    `gorm:"column:loneliness;NOT NULL"`
	EarlyLoveImpact string    `gorm:"column:early_love_impact;NOT NULL"`
	Alcohol         string    `gorm:"column:alcohol;NOT NULL"`
	ProblemSolving  string    `gorm:"column:problem_solving;NOT NULL"`
	Failure         string    `gorm:"column:failure;NOT NULL"`
	EducationImpact string    `gorm:"column:education_impact;NOT NULL"`
	FutureJob       string    `gorm:"column:future_job;NOT NULL"`
	WrittenBy       string    `gorm:"column:WrittenBy;NOT NULL"`
	CreationTime    time.Time `gorm:"column:Creation_time;default:NULL"`
	Remark          string    `gorm:"column:remark;default:NULL"`
}

func (t *TeenSurvey) TableName() string {
	return "teen_survey"
}

// 大学生问卷表
type CollegeSurvey struct {
	Id                         int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Grade                      string    `gorm:"column:grade;NOT NULL"`
	Gender                     string    `gorm:"column:gender;NOT NULL"`
	Birthplace                 string    `gorm:"column:birthplace;NOT NULL"`
	MonthlyExpense             string    `gorm:"column:monthly_expense;NOT NULL"`
	HomeTown                   string    `gorm:"column:home_town;NOT NULL"`
	Expectations               string    `gorm:"column:expectations;NOT NULL"`
	SingleChild                string    `gorm:"column:single_child;NOT NULL"`
	FutureJobExpectation       string    `gorm:"column:future_job_expectation;NOT NULL"`
	RelationshipWithClassmates string    `gorm:"column:relationship_with_classmates;NOT NULL"`
	ExamTasks                  string    `gorm:"column:exam_tasks;NOT NULL"`
	AbilityToHandle            string    `gorm:"column:ability_to_handle;NOT NULL"`
	CareAboutOthers            string    `gorm:"column:care_about_others;NOT NULL"`
	SelfRequirement            string    `gorm:"column:self_requirement;NOT NULL"`
	ImpactByGrade              string    `gorm:"column:impact_by_grade;NOT NULL"`
	ImpactByGender             string    `gorm:"column:impact_by_gender;NOT NULL"`
	WrittenBy                  string    `gorm:"column:WrittenBy;NOT NULL"`
	CreationTime               time.Time `gorm:"column:Creation_time;default:NULL"`
	Remark                     string    `gorm:"column:remark;default:NULL"`
	Level                      int32     `gorm:"column:level;default:NULL"`
}

func (c *CollegeSurvey) TableName() string {
	return "college_survey"
}

// WorkerSurvey 社会工作者问卷表
type WorkerSurvey struct {
	Id                int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	WorkUnit          string    `gorm:"column:work_unit;NOT NULL"`
	Gender            string    `gorm:"column:gender;NOT NULL"`
	Age               string    `gorm:"column:age;NOT NULL"`
	MaritalStatus     string    `gorm:"column:marital_status;NOT NULL"`
	EducationLevel    string    `gorm:"column:education_level;NOT NULL"`
	ProfessionalTitle string    `gorm:"column:professional_title;NOT NULL"`
	Department        string    `gorm:"column:department;NOT NULL"`
	WorkYears         string    `gorm:"column:work_years;NOT NULL"`
	EnterpriseNature  string    `gorm:"column:enterprise_nature;NOT NULL"`
	WrittenBy         string    `gorm:"column:WrittenBy;NOT NULL"`
	CreationTime      time.Time `gorm:"column:Creation_time;default:NULL"`
	Remark            string    `gorm:"column:remark;default:NULL"`
}

func (w *WorkerSurvey) TableName() string {
	return "worker_survey"
}
