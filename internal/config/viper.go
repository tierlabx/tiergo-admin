package config

import (
	"fmt"
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
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Config struct {
	DB     DbConfig  `mapstructure:"DB"`
	WebApi WebConfig `mapstructure:"WebApi"`
}

func (d *Config) InitConfig() {
	// 加载 .env 环境变量文件
	_ = godotenv.Load()

	viper.SetConfigFile("tier-up/config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")              // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")                // 如果配置文件的名称中没有扩展名，则需要配置此项

	viper.AddConfigPath(".") // 还可以在工作目录中查找配置
	viper.WatchConfig()      //监控配置文件更新

	err := viper.ReadInConfig() // 查找并读取配置文件

	viper.AutomaticEnv()                                   // 自动查找环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // app.port → APP_PORT
	// viper 会自动映射 同名变量
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file %v", err))
	}
	// 读取配置文件
	autoCreateTable := viper.GetBool("DB.AutoCreateTable")
	user := viper.GetString("DB.User")
	password := viper.GetString("DB.Password")
	host := viper.GetString("DB.Host")
	port := viper.GetString("DB.Port")
	driverName := viper.GetString("DB.DriverName")
	charset := viper.GetString("DB.Charset")

	d.DB = DbConfig{
		AutoCreateTable: autoCreateTable,
		User:            user,
		Password:        password,
		Port:            port,
		Host:            host,
		DriverName:      driverName,
		Charset:         charset,
	}
	d.WebApi = WebConfig{
		Host: viper.GetString("WebApi.Host"),
		Port: viper.GetString("WebApi.Port"),
	}

}

func Load() Config {
	var c Config
	c.InitConfig()
	return c
}
