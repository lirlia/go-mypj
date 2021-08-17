package main

import "fmt"

type RedisCache struct {
	RedisCacheInterface
}

type RedisCacheInterface interface {
	Delete()
	Set()
}

type RedisCacheAllStrategyImpl struct{}
type RedisCacheModuloStrategyImpl struct{}

func (r *RedisCacheAllStrategyImpl) Delete()    { fmt.Println("delete all") }
func (r *RedisCacheModuloStrategyImpl) Delete() { fmt.Println("delete specific") }
func (r *RedisCacheAllStrategyImpl) Set()       { fmt.Println("set all") }
func (r *RedisCacheModuloStrategyImpl) Set()    { fmt.Println("set specific") }

func NewRedisCache(i RedisCacheInterface) RedisCacheInterface {
	return &RedisCache{i}
}

type UpdateSystemConfigService struct {
	RedisCacheInterface
}

func NewUpdateSystemConfigService(i RedisCacheInterface) *UpdateSystemConfigService {
	return &UpdateSystemConfigService{i}
}

func main() {
	rall := &RedisCacheAllStrategyImpl{}
	rmod := &RedisCacheModuloStrategyImpl{}

	rc := NewRedisCache(rall)
	rc.Set()
	rc.Delete()

	rc = NewRedisCache(rmod)
	rc.Set()
	rc.Delete()

	sys := NewUpdateSystemConfigService(rall)
	sys.Set()
	sys.Delete()
}
