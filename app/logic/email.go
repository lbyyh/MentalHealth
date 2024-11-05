package logic

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/tools"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v2"
)

const (
	SMTPHost     = "smtp.qq.com"
	SMTPPort     = 587 // SMTP服务器的端口
	SMTPUsername = "412213958@qq.com"
	SMTPPassword = "bgutozutckrwbihh"
)

// CaptchaEmailRequest 用于绑定请求体
type CaptchaEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// sendEmail 函数负责发送实际的邮箱验证码
func sendCaptchaEmail(address string, captcha string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", SMTPUsername)
	m.SetHeader("To", address)
	m.SetHeader("Subject", "您的验证码")
	m.SetBody("text/plain", "您的验证代码是："+captcha)

	d := gomail.NewDialer(SMTPHost, SMTPPort, SMTPUsername, SMTPPassword)

	return d.DialAndSend(m)
}

// SendEmailCaptcha 是我们的Gin处理函数
func SendEmailCaptcha(c *gin.Context) {
	var req CaptchaEmailRequest

	// 绑定JSON到我们的请求结构体类型
	if err := c.BindJSON(&req); err != nil {
		// 如果JSON绑定失败或邮箱验证失败则返回错误
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "无效的电子邮件地址",
		})
		return
	}

	// 生成验证码
	rand.Seed(time.Now().UnixNano())
	captcha := strconv.Itoa(100000 + rand.Intn(899999))

	// 生成的验证码保存到Redis
	expiration := 10 * time.Minute // 设置验证码的失效时间，这里假设为10分钟
	err := model.Redis.Set(c, req.Email, captcha, expiration).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    10004,
			"message": "保存验证码失败",
		})
		return
	}

	// 发送验证码
	if err := sendCaptchaEmail(req.Email, captcha); err != nil {
		// 如果发送邮件出错
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    10005,
			"message": "发送验证码邮件失败",
		})
		return
	}

	// 此处应有将验证码保存到某个存储的逻辑，以便后续验证

	// 发送成功
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "验证码发送成功",
	})
}

// VerifyEmailCaptchaRequest 用于绑定验证请求体
type VerifyEmailCaptchaRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Captcha string `json:"captcha" binding:"required,len=6"`
}

// VerifyEmailCaptcha 验证用户提供的验证码
func VerifyEmailCaptcha(c *gin.Context, req VerifyEmailCaptchaRequest) bool {
	// 从Redis中获取验证码
	storedCaptcha, err := model.Redis.Get(c, req.Email).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "验证码过期或找不到电子邮件",
		})
		return false
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "检索验证码错误",
		})
		return false
	}

	if req.Captcha == storedCaptcha {
		// 成功验证逻辑
		// 删除已验证的验证码
		model.Redis.Del(c, req.Email)
		return true
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    1,
		"message": "无效的验证码",
	})
	return false

}

func EmailLogin(c *gin.Context) {
	fmt.Println("-------------------------")
	var users VerifyEmailCaptchaRequest

	if err := c.ShouldBind(&users); err != nil {
		c.JSON(http.StatusBadRequest, tools.ECode{ // 修改为 HTTP 400
			Code:    10001,
			Message: "输入参数有误", // 更改为通用错误消息，避免敏感信息泄露
		})
		return
	}
	if !VerifyEmailCaptcha(c, users) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "无效的验证码",
		})
		return
	}
	ret := model.GetUserbyE(users.Email)
	// 生成TOKEN
	token, err := model.GetJwt(int64(ret.Id), ret.Name)
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

	return
}
