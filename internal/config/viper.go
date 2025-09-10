package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// DbConfig MySQL 链接配置
// 字段 类型 员信息
type DbConfig struct {
	AutoCreateTable bool   `yaml:"AutoCreateTable"`
	Driver          string `yaml:"Driver"`
	DriverName      string `yaml:"DriverName"`
	Host            string `yaml:"Host"`
	User            string `yaml:"User"`
	Password        string `yaml:"Password"`
	Port            string `yaml:"Port"`
	Charset         string `yaml:"Charset"`
}

type WebConfig struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
type Config struct {
	DB     DbConfig  `mapstructure:"DB"`
	WebApi WebConfig `mapstructure:"WebApi"`
}

func (d *Config) InitConfig() {
	_ = godotenv.Load() // 忽略错误，生产可能没有 .env

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 环境变量优先
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件没有找到: %v", err))
	}

	// 绑定配置
	if err := viper.Unmarshal(&d); err != nil {
		panic(fmt.Errorf("无法解码配置: %v", err))
	}

	// 优先环境变量覆盖敏感字段
	if pwd := os.Getenv("DB_PASSWORD"); pwd != "" {
		d.DB.Password = pwd
	}
}

func Load() Config {
	var c Config
	c.InitConfig()
	return c
}
