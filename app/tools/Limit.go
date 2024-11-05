package tools

import (
	"MentalHealth-Platform/app/model"
	"context"
	"time"
)

var ctx = context.Background()

// IsRateLimited 检查用户是否达到了操作的频率限制
func IsRateLimited(userId, actionKey string, period time.Duration, limit int) (bool, error) {
	// 保证有一个有效的客户端连接

	key := "rate_limit:" + userId + ":" + actionKey

	// 使用Redis的INCR命令递增计数器，并检查错误
	val, err := model.Redis.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if val == 1 {
		// 如果是第一次计数则设置过期时间
		if _, err := model.Redis.Expire(ctx, key, period).Result(); err != nil {
			return false, err
		}
	}

	// 检查计数器值是否超过了限制
	return val > int64(limit), nil
}
