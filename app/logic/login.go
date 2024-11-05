package logic

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Users struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" form:"password"`
	CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

// 刷新图片验证码
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "获取上传文件失败",
		})
		return
	}

	// 确保avatars目录存在
	avatarDirectory := "app/images/avatars"
	if _, err := os.Stat(avatarDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(avatarDirectory, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    1,
				"message": "创建目录失败",
			})
			return
		}
	}

	// 创建唯一文件名
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timestamp, filepath.Base(file.Filename))
	filePath := filepath.Join(avatarDirectory, fileName)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    2,
			"message": "保存文件失败",
		})
		return
	}

	// 返回文件路径
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"message":  "文件上传成功",
		"filePath": fmt.Sprintf("/images/avatars/%s", fileName),
	})
}

// Login 执行用户登录
func Login(c *gin.Context) {
	fmt.Println("-------------------------")
	var users Users
	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: "输入参数有误", // 更改为通用错误消息，避免敏感信息泄露
		})
		c.Abort() // 终止请求处理
		return
	}

	if !tools.CaptchaVerify(tools.CaptchaData{
		CaptchaId: users.CaptchaId,
		Data:      users.CaptchaValue,
	}) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "验证码校验失败", // 更改为通用错误消息，避免敏感信息泄露
		})
		c.Abort() // 终止请求处理
		return
	}

	ret := model.GetUser(users.Name)
	fmt.Printf("ret.Password:%v\n", ret.Password)
	fmt.Printf("tools.EncryptV1(user.Password):%v\n", tools.EncryptV1(users.Password))
	if ret.Id < 1 || ret.Password != tools.EncryptV1(users.Password) {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: "帐号密码错误",
		})
		c.Abort() // 终止请求处理
		return
	}

	// 生成TOKEN
	token, err := model.GetJwt(int64(ret.Id), users.Name)
	c.SetCookie("token", token, 3600, "/user", "", true, false)
	if err != nil {
		c.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "登录失败，无法生成token",
		})
		c.Abort() // 终止请求处理
		return
	}

	// 将token发送给客户端
	c.JSON(http.StatusOK, tools.ECode{
		Code:    0, // 一般来说成功的响应代码是 0
		Message: "登录成功",
		Data:    token,
	})

	// 不需要再次调用 c.JSON，因为token已经在上面发送过了
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })

	return
}
