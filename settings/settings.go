package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}


type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}


type RedisConfig struct {
	Host     string `mapstructure:"host"`
	PassWord string `mapstructure:""`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath) // 指定配置文件路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {               // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed: %s \n", err)
		return err
	}
	// Unmarshal -> Conf
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed: %s \n", err)
		return err
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了~~~")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed: %s \n", err)

		}
	})
	return
}
