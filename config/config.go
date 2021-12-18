package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/lexkong/log"
)

// 项目配置
type Config struct {
	Name        string
	ViperConfig *ViperConfig
}

var C Config

// 配置文件的配置项
type ViperConfig struct {
	RunMode      string
	Addr         string
	Name         string
	PingUrl      string
	MaxPingCount int
	Log          LogConfg
	Mysql        MySqlConfig
}

type LogConfg struct {
	Writers        string
	LoggerLevel    string
	LoggerFile     string
	LogFormatText  bool
	RollingPolicy  string
	LogRotateDate  int
	LogBackupCount int
	LogRotateSize  int
}
type MySqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// initConfig function init config info
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GOAPI")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c.ViperConfig); err != nil {
		log.Fatal("配置文件转换为 结构体失败", err)
	}
	return nil
}

// watchConfig funvc watch config file change then hot reload
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// log.Printf("Config file changed: %s", in.Name)
	})
}

func (c *Config) initLog() {
	log.InitWithConfig(&log.PassLagerCfg{
		Writers:        C.ViperConfig.Log.Writers,
		LoggerLevel:    C.ViperConfig.Log.LoggerLevel,
		LoggerFile:     C.ViperConfig.Log.LoggerFile,
		LogFormatText:  C.ViperConfig.Log.LogFormatText,
		LogRotateDate:  C.ViperConfig.Log.LogRotateDate,
		LogRotateSize:  C.ViperConfig.Log.LogRotateSize,
		LogBackupCount: C.ViperConfig.Log.LogBackupCount,
		RollingPolicy:  C.ViperConfig.Log.RollingPolicy,
	})
}

// func init() {}

func Init(cfg string) error {
	C = Config{
		Name:        cfg,
		ViperConfig: &ViperConfig{},
	}
	if err := C.initConfig(); err != nil {
		return err
	}
	// 初始化 日志
	C.initLog()
	C.watchConfig()
	return nil
}
