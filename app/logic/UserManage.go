package logic

import (
	"MentalHealth-Platform/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsersList 处理获取所有用户的请求
func GetUsersList(c *gin.Context) {
	users := model.GetAllUsers(c)

	processedUsers := make([]map[string]interface{}, len(users))

	for i, user := range users {
		processedUsers[i] = map[string]interface{}{
			"Id":               user.Id,
			"Name":             user.Name,
			"Password":         user.Password,
			"Age":              user.Age,
			"Gender":           user.Gender,
			"RegistrationDate": user.RegistrationDate,
			"ContactInfo":      user.ContactInfo,
			"Status":           user.Status,
			"LastLandingTime":  user.LastLandingTime,
		}
	}
	fmt.Println("processedUsers:=================", processedUsers)
	// 将处理后的用户数据返回到前端
	c.JSON(http.StatusOK, processedUsers)
}
