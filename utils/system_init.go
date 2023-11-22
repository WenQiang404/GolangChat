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
	"time"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
}

func InitRedis() {
	REDIS = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})

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
	sub := REDIS.PSubscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println(msg.Payload)
	return msg.Payload, err
}
