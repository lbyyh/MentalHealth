package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 向数据库获取用户的函数
func GetUser(name string) *Users {
	var ret Users
	err := MySQL.Table("users").Where("name=?", name).First(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return &ret
	}
	return &ret
}

// GetUserbyE 向数据库获取用户的函数(通过email)
func GetUserbyE(email string) *Users {
	var ret Users
	err := MySQL.Table("users").Where("contact_info=?", email).First(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return &ret
	}
	return &ret
}

// CreateUser 向数据库添加新用户的函数
func CreateUser(name, password, email, gender string, age int32) (*Users, error) {
	// 创建User实例
	newUser := Users{
		Name:             name,
		Password:         password,
		Age:              age,
		Gender:           gender,
		RegistrationDate: time.DateTime,
		ContactInfo:      email,
	}

	// 插入新创建的用户记录到数据库
	if err := MySQL.Table("users").Create(&newUser).Error; err != nil {
		fmt.Printf("Create users error: %s", err.Error())
		return nil, err
	}
	return &newUser, nil
}

// GetAllUsers 获取所有用户的函数
func GetAllUsers(c *gin.Context) []Users {
	var users []Users
	err := MySQL.Table("users").Find(&users).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	return users
}

// UpdateUser 修改用户信息的函数
func UpdateUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the data to update
	data := make(map[string]interface{})
	data["Name"] = user.Name
	data["Age"] = user.Age
	data["Gender"] = user.Gender
	data["ContactInfo"] = user.ContactInfo

	// 定位到需要更新的用户，并更新指定的字段
	if err := MySQL.Model(&Users{}).Where("id = ?", user.Id).Updates(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 响应更新成功
	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功！", "user": user})
}

// DeleteUser 根据id删除用户信息的函数
func DeleteUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 删除 TeenSurvey 中与用户关联的行
	if err := MySQL.Where("WrittenBy = ?", user.Id).Delete(&TeenSurvey{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除 CollegeSurvey 中与用户关联的行
	if err := MySQL.Where("WrittenBy = ?", user.Id).Delete(&CollegeSurvey{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除 WorkerSurvey 中与用户关联的行
	if err := MySQL.Where("WrittenBy = ?", user.Id).Delete(&WorkerSurvey{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 删除用户
	if err := MySQL.Where("id = ?", user.Id).Delete(&Users{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功！"})
}

// SuspendUser 用于处理挂起用户的请求
func SuspendUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the data to update
	data := make(map[string]interface{})
	data["Status"] = 1

	// 定位到需要更新的用户，并更新指定的字段
	if err := MySQL.Model(&Users{}).Where("id = ?", user.Id).Updates(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 响应更新成功
	c.JSON(http.StatusOK, gin.H{"message": "用户已成功挂起！", "user": user})
}

// UnSuspendUser 用于处理取消挂起用户的请求
func UnSuspendUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the data to update
	data := make(map[string]interface{})
	data["Status"] = 0

	// 定位到需要更新的用户，并更新指定的字段
	if err := MySQL.Model(&Users{}).Where("id = ?", user.Id).Updates(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 响应更新成功
	c.JSON(http.StatusOK, gin.H{"message": "用户已取消挂起！", "user": user})
}

// AddUser - 添加新用户信息的函数
func AddUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 设置用户注册日期为当前时间，并格式化为 "2006-01-02"
	user.RegistrationDate = time.Now().Format("2006-01-02")
	// 当传入的JSON绑定到用户结构时，我们将创建新用户信息
	if err := MySQL.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 响应操作成功
	c.JSON(http.StatusOK, gin.H{"message": "新用户信息已经成功创建！", "user": user})
}
