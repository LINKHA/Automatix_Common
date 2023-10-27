package config

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	// 设置配置文件名和路径
	viper.SetConfigFile("config.yaml")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("无法读取配置文件:", err)
	}

	// 将配置解析到结构体中
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("无法解析配置:", err)
	}
	fmt.Println("---------------------------")
	fmt.Println(config)
}
