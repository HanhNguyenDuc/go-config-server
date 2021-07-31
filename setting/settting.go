package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App ... Define struct
type App struct {
	StatusSuccess int
	StatusError   int
}

// AppSetting ... Init object
var AppSetting = &App{}

// Paginator ... Define struct
type Paginator struct {
	PageSize int
}

// PaginatorSetting ... Init object
var PaginatorSetting = &Paginator{}

// Server ... Define struct
type Server struct {
	RunMode                     string
	HTTPPort                    int
	ReadTimeout                 time.Duration
	WriteTimeout                time.Duration
	JWTAccessTokenSecretKey     string
	JWTAccessTokenExpireMinutes time.Duration
	JWTAccessTokenRedisPrefix   string
	JWTRefreshTokenSecretKey    string
	JWTRefreshTokenExpireHours  time.Duration
	JWTRefreshTokenRedisPrefix  string
}

// ServerSetting ... Init object
var ServerSetting = &Server{}

// Database ... Define struct
type Database struct {
	Type               string
	User               string
	Password           string
	Host               string
	Name               string
	TablePrefix        string
	MaxIdleConnections int
	MaxOpenConnections int
}

// DatabaseSetting ... Init object
var DatabaseSetting = &Database{}

// Redis ... Define struct
type Redis struct {
	RedisAddress  string
	RedisPassword string
	RedisDB       int
}

// RedisSetting ... Init object
var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	// log.Printf("port: %d", ServerSetting.HTTPPort)
}

// mapTo ... Map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
