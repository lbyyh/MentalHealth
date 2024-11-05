package model

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	MySQL *gorm.DB
	Redis *redis.Client
)

func Init(mysqlConfig MysqlConfig, redisConfig RedisConfig) error {
	var err error
	MySQL, err = initMySQL(mysqlConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize MySQL: %s", err)
	}
	Redis, store = initRedis(redisConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize MySQL: %s", err)
	}
	return nil
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app")
	// 读取配置数据
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 从配置文件中读取数据库配置
	redisConf := RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetInt("redis.port"),
		Password: viper.GetString("redis.password"),
	}
	mysqlConf := MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	}

	// 初始化数据库连接
	err := Init(mysqlConf, redisConf)
	fmt.Printf("mysqlConf: %+v\n", mysqlConf)
	fmt.Printf("redisConf: %+v\n", redisConf)
	if err != nil {
		return
	}
	//defer Close() // 确保在程序退出时关闭数据库连接
}

func initMySQL(c MysqlConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %s", err)
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %s", err)
	}
	return db, nil
}

func initRedis(c RedisConfig) (*redis.Client, *redisstore.RedisStore) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       0,
	})
	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	store, err := redisstore.NewRedisStore(ctx, client)
	if err != nil {
		log.Fatalf("Failed to create redis store: %v", err)
	}
	return client, store
}

func Close() {
	if MySQL != nil {
		sqlDB, err := MySQL.DB()
		if err != nil {
			log.Fatalf("Failed to get SQL DB from GORM: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close MySQL connection: %v", err)
		}
	}
	if Redis != nil {
		if err := Redis.Close(); err != nil {
			log.Fatalf("Failed to close Redis connection: %v", err)
		}
	}
}
