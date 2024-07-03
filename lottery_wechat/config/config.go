package config

import (
	rlog "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	globalConfig GlobalConfig
	once         sync.Once
)

type GlobalConfig struct {
	AppConfig AppConf `yaml:"app" mapstructure:"app"`
	LogConfig LogConf `yaml:"log" mapstructure:"log"`
}

type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"`
	Version string `yaml:"version" mapstructure:"version"`
	RunMod  string `yaml:"run_mod" mapstructure:"run_mod"`
	Port    int    `yaml:"port" mapstructure:"port"`
}
type LogConf struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"`
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`
	SaveDays   int    `yaml:"save_days" mapstructure:"save_days"`
	Level      string `yaml:"level" mapstructure:"level"`
}

func GetGlobalConfig() *GlobalConfig {
	once.Do(func() {
		readConf()
	})
	return &globalConfig
}
func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read config error")
	}
	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		panic("viper Unmarshal config error")
	}
}
func InitGlobalConfig() {
	conf := GetGlobalConfig()
	level, err := log.ParseLevel(conf.LogConfig.Level)
	if err != nil {
		panic("parse level err")
	}
	log.SetLevel(level)
	log.SetFormatter(&LogFormatter{
		log.TextFormatter{
			DisableColors:   true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	})
	log.SetReportCaller(true)
	switch conf.LogConfig.LogPattern {
	case "stdout":
		log.SetOutput(os.Stdout)
		break
	case "stderr":
		log.SetOutput(os.Stderr)
		break
	case "file":
		rl, err := rlog.New(conf.LogConfig.LogPath+".%Y%m%d", rlog.WithRotationTime(time.Hour*24), rlog.WithRotationCount(conf.LogConfig.SaveDays))
		if err != nil {
			panic("rotate log create fail")
		}
		log.SetOutput(rl)
		break
	default:
		panic("no out")
	}
}
