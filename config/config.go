package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/lexkong/log"
)

type Config struct {
	Name string
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
		Writers:        viper.GetString("writers"),
		LoggerLevel:    viper.GetString("logger_level"),
		LoggerFile:     viper.GetString("logger_file"),
		LogFormatText:  viper.GetBool("log_format_text"),
		LogRotateDate:  viper.GetInt("log_rotate_date"),
		LogRotateSize:  viper.GetInt("log_rotate_size"),
		LogBackupCount: viper.GetInt("log_backup_count"),
		RollingPolicy:  viper.GetString("rollingPolicy"),
	})
}

// func init() {}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	if err := c.initConfig(); err != nil {
		return err
	}
	// 初始化 日志
	c.initLog()
	c.watchConfig()
	return nil
}
