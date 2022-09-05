package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// TomlConfig represent the root of configuration
type SequoiaConfig struct {
	Title            string
	Server           Server
	Logging          Logging
	Fs               Freeswitch
	Redis            Redis
	Heartbeat        Heartbeat
	RecordingService RecordingService
	Rating           Rating
	Numbers          Numbers
	SipEndpoint      SipEndpoint
	Kamgo            Kamgo
	Sls              SLS
}

type Heartbeat struct {
	BaseUrl  string
	UserName string
	Secret   string
}

type RecordingService struct {
	BaseUrl  string
	UserName string
	Secret   string
}

type Rating struct {
	BaseUrl  string
	UserName string
	Secret   string
	Region   string
}
type Numbers struct {
	BaseUrl  string
	UserName string
	Secret   string
}
type SipEndpoint struct {
	BaseUrl  string
	UserName string
	Secret   string
}

type Kamgo struct {
	BaseUrl  string
	UserName string
	Secret   string
}

// Server 服务启动信息
type Server struct {
	Port          string
	QueueLen      int
	ErrorQueueLen int
	GinMode       string
}

// Logging 日志信息
type Logging struct {
	Facility string
	Level    string
	Tag      string
	Syslog   string
	Sentry   string
	Path     string
	Day      int
	Hour     int
}

// FreeSwitch SIP 链接信息
type Freeswitch struct {
	FsHost     string
	FsPort     string
	FsPassword string
	FsTimeout  int
}

//存储Redis信息
type Redis struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

//SLS 对接阿里SLS日志系统
type SLS struct {
	Enable       int
	Endpoint     string
	AccessKey    string
	AccessSecret string
	Project      string
	LogStor      string
	Topic        string
	Source       string
}

var Config SequoiaConfig

func InitConfig() error {
	var err error

	configFile := os.Getenv("SEQUOIA_CONFIG")
	if len(configFile) == 0 {
		configFile = "./config.toml"
	}

	if _, err = toml.DecodeFile(configFile, &Config); err != nil {
		return err
	}

	return nil
}
