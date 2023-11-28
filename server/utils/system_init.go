package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	log2 "server/log"
	"time"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client
)

func newLogger() *log2.Logger {
	return log2.NewLogger()
}

var myLog = newLogger()

func InitConfig() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		myLog.Fatal("Failed to init" + err.Error())
		return err
	}
	return nil
}

func InitMySQL() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	_, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		myLog.Error("Failed to init the MySQL table" + err.Error())
		return err
	}
	myLog.Info("Init the mySql at")
	return nil
}

func InitRedis() error {
	REDIS = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	ctx := context.Background()
	_, err := REDIS.Ping(ctx).Result()
	if err != nil {
		myLog.Error("redis initilization...." + err.Error())
		return err
	} else {
		myLog.Info("success to init redis at ")

	}
	return nil

}

const (
	PublishKey = "websocket"
)

// 发送消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = REDIS.Publish(ctx, channel, msg).Err()
	return err
}

// 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := REDIS.Subscribe(ctx, channel)
	fmt.Println("Subscribe the message......", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		myLog.Error(err.Error())
		return "", err
	}

	return msg.Payload, err
}
