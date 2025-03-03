package main

import (
	"kafkacacheflow/services/redis_updater/internal/service"
	"time"
)

func main() {
	service.NewServiceManager()
	time.Sleep(time.Minute * 2)
}
