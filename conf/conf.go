package conf

import "time"

var RedisTimeout time.Duration = 100 * time.Second

var MaxTimeToResponse int64 = 10 * 1000 // 20 second
