package main

import (
	"twof/blog-api/initiator"
	zaplog "twof/blog-api/internal/log"
)

func main() {
	zaplog.InitLogger()
	defer zaplog.SugerLogger.Sync()
	initiator.Init()
}
