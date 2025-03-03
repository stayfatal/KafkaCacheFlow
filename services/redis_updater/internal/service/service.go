package service

import (
	"kafkacacheflow/services/redis_updater/internal/pool"
	"time"
)

type ServiceManager struct {
	pool *pool.Pool
}

func NewServiceManager() *ServiceManager {
	sm := &ServiceManager{pool: pool.NewPool()}
	sm.consume()
	return sm
}

func (sm *ServiceManager) consume() {
	for {
		time.Sleep(time.Duration(time.Now().UnixNano()%5) * time.Second)
		sm.pool.NewTask(func() {
			time.Sleep(time.Duration(time.Now().UnixNano()%10) * time.Second)
		})
	}
}
