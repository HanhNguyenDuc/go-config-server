package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App ... Define struct
type App struct {
	StatusSuccess int `env:"APP_STATUS_SUCCESS"`
	StatusError   int `env:"APP_STATUS_ERROR"`
}

// AppSetting ... Init object
var AppSetting = &App{}

// Server ... Define struct
type Server struct {
	RunMode                     string        `env:"SERVER_RUN_MODE"`
	HTTPPort                    int           `env:"SERVER_HTTP_PORT"`
	ReadTimeout                 time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout                time.Duration `env:"SERVER_WRITE_TIMEOUT"`
	JWTAccessTokenSecretKey     string        `env:"SERVER_JWT_ACCESS_TOKEN_SECRET_KEY"`
	JWTAccessTokenExpireMinutes time.Duration `env:"SERVER_JWT_ACCESS_TOKNE_EXPIRE_MINUTES"`
	JWTAccessTokenRedisPrefix   string        `env:"SERVER_JWT_ACCESS_TOKEN_REDIS_PREFIX"`
	JWTRefreshTokenSecretKey    string        `env:"SERVER_JWT_REFRESH_TOKEN_SECRET_KEY"`
	JWTRefreshTokenExpireHours  time.Duration `env:"SERVER_JWT_REFRESH_TOKEN_EXPIRE_HOURS"`
	JWTRefreshTokenRedisPrefix  string        `env:"SERVER_JWT_REFRESH_TOKEN_REDIS_PREFIX"`
}

// ServerSetting ... Init object
var ServerSetting = &Server{}

// Database ... Define struct
type Database struct {
	Type               string `env:"DATABASE_TYPE"`
	User               string `env:"DATABASE_USER"`
	Password           string `env:"DATABASE_PASSWORD"`
	Host               string `env:"DATABASE_HOST"`
	Name               string `env:"DATABASE_NAME"`
	TablePrefix        string `env:"DATABASE_TABLE_PREFIX"`
	MaxIdleConnections int    `env:"DATABASE_IDLE_CONNECTIONS"`
	MaxOpenConnections int    `env:"DATBASE_MAX_OPEN_CONNECTIONS"`
}

// DatabaseSetting ... Init object
var DatabaseSetting = &Database{}

// Redis ... Define struct
type Redis struct {
	RedisAddress  string `env:"REDIS_ADDRESS"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`
}

// RedisSetting ... Init object
var RedisSetting = &Redis{}

var cfg *ini.File

func Setup(isDeploy *bool) {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	if isDeploy != nil && !*isDeploy {
		mapTo("app", AppSetting)
		mapTo("server", ServerSetting)
		mapTo("database", DatabaseSetting)
		mapTo("redis", RedisSetting)
	} else {
		parseFromEnvs(AppSetting)
		parseFromEnvs(ServerSetting)
		parseFromEnvs(DatabaseSetting)
		parseFromEnvs(RedisSetting)
	}
	log.Printf("[info] serverSetting: %v", *ServerSetting)
	// log.Printf("port: %d", ServerSetting.HTTPPort)
}

// mapTo ... Map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
