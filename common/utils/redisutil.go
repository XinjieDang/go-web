package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisClientWrapper 提供封装后的 Redis 操作方法。
type RedisClientWrapper struct {
	client *redis.Client
}

// NewRedisClientWrapper 创建一个新的 RedisClientWrapper 实例-构造函数
func NewRedisClientWrapper(client *redis.Client) *RedisClientWrapper {
	return &RedisClientWrapper{client: client}
}

// Set 设置一个键值对，带过期时间。
func (rcw *RedisClientWrapper) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 设置操作超时时间
	defer cancel()
	return rcw.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取一个键的值。
func (rcw *RedisClientWrapper) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	val, err := rcw.client.Get(ctx, key).Result()
	if err == redis.Nil {
		err = nil // 若键不存在，返回空字符串而不是错误
	}
	return val, err
}

// Delete 删除个键。
func (rcw *RedisClientWrapper) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rcw.client.Del(ctx, key).Err()
}

// Close 关闭 Redis 客户端连接。
func (rcw *RedisClientWrapper) Close() error {
	return rcw.client.Close()
}
