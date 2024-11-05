package logic

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

func AdminLogin(c *gin.Context) {
	var admin Admin
	if err := c.ShouldBind(&admin); err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(), //这里有风险
		})
	}

	fmt.Printf("admin:%+v\n", admin)

	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: admin.CaptchaId,
		Data:      admin.CaptchaValue,
	}) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "验证码校验失败！", //这里有风险
		})
		return
	}

	ret := model.GetAdmin(admin.Name)
	if ret.AdminId < 1 || ret.Password != tools.EncryptV1(admin.Password) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: "帐号密码错误！",
		})
		return
	}

	_ = model.SetSession(c, admin.Name, int64(ret.AdminId))
	c.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
	})
}

// AdminLogout 管理员退出登录
func AdminLogout(c *gin.Context) {
	_ = model.FlushSession(c)
	c.JSON(http.StatusUnauthorized, tools.ECode{
		Code:    0,
		Message: "您已退出登录",
	})
}

// AllPersonalInformation 获取管理员信息列表的函数
func AllPersonalInformation(c *gin.Context) {
	var admins []model.Admin // 注意这里改为了切片

	// 执行id的管理员信息查询
	if err := model.MySQL.Table("admin").Find(&admins).Error; err != nil {
		//如果查询不成功，则返回错误
		c.JSON(http.StatusNotFound, gin.H{"error": "管理员信息未找到！"})
		return
	}

	// 响应操作成功
	c.JSON(http.StatusOK, gin.H{"message": "查询操作成功！", "admins": admins}) // 注意这里改为了 admins
}

// PersonalInformation 获取指定id的管理员信息的函数
func PersonalInformation(c *gin.Context) {
	var admin model.Admin
	// 获取用户信息
	cookie, err := c.Request.Cookie("session-name")
	fmt.Println("cookie----", cookie)
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}

	jwt := cookie.Value
	fmt.Println("jwt----", jwt)
	adminData := model.GetSession(c)
	fmt.Println("adminData----", adminData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		fmt.Println("无效的token")
		return
	}
	id := adminData["id"]

	// 执行id的管理员信息查询
	if err := model.MySQL.Table("admin").Where("admin_id = ?", id).First(&admin).Error; err != nil {
		//如果查询不成功，则返回错误
		c.JSON(http.StatusNotFound, gin.H{"error": "管理员信息未找到！"})
		return
	}

	// 响应操作成功
	c.JSON(http.StatusOK, gin.H{"message": "查询操作成功！", "admin": admin})
}

// UpdatePersonalInformation 根据id修改管理员自身信息的函数
func UpdatePersonalInformation(c *gin.Context) {
	var admin model.Admin
	// 获取用户信息
	cookie, err := c.Request.Cookie("session-name")
	fmt.Println("cookie----", cookie)
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}

	jwt := cookie.Value
	fmt.Println("jwt----", jwt)
	adminData := model.GetSession(c)
	fmt.Println("adminData----", adminData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		fmt.Println("无效的token")
		return
	}
	id := adminData["id"]
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.MySQL.Model(&model.Admin{}).Where("admin_id = ?", id).Updates(admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	// 响应操作成功
	c.JSON(http.StatusOK, gin.H{"message": "更新管理员信息成功！", "admin": admin})
}

// UpdatePassWord 根据id验证和修改管理员的信息
func UpdatePassWord(c *gin.Context) {
	var admin model.Admin
	var adminInput struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	// 获取用户信息
	cookie, err := c.Request.Cookie("session-name")
	if err != nil || cookie.Value == "" {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10000,
			Message: "无法获取token或token为空",
		})
		return
	}

	adminData := model.GetSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    10001,
			Message: "无效的token",
		})
		return
	}
	id := adminData["id"]

	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取旧密码并匹配
	if err := model.MySQL.Where("admin_id = ?", id).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "管理员信息未找到！"})
		return
	}

	if tools.EncryptV1(adminInput.OldPassword) != admin.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "旧密码错误！"})
		return
	}

	// 设置新密码并更新
	admin.Password = tools.EncryptV1(adminInput.NewPassword)
	if err := model.MySQL.Model(&model.Admin{}).Where("admin_id = ?", id).Updates(admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	// 响应操作成功
	c.JSON(http.StatusOK, gin.H{"message": "密码更新成功！"})
}
