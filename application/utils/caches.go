package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var MemCache = cache.New(5*time.Minute, 10*time.Minute)
