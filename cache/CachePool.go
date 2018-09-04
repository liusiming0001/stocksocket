package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisPool struct {
	RedisClient *redis.Pool
	RedisHost   string
	RedisDb     int
	RedisAuth   string
	RedisCon    redis.Conn
}

func (pool *RedisPool) GetPool() {
	pool.RedisClient = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", pool.RedisHost)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("AUTH", pool.RedisAuth)
			// 选择db
			c.Do("SELECT", pool.RedisDb)
			return c, nil
		},
	}
	pool.RedisCon = pool.RedisClient.Get()
}

//string存储
func (pool *RedisPool) StringRedisSet(key string, val string) {
	//_redisLock.Lock()
	pool.RedisCon.Do("SET", key, val)
	//_redisLock.Unlock()
}
func (pool *RedisPool) StringRedisGet(key string) *string {
	//_redisLock.Lock()
	result, _ := redis.String(pool.RedisCon.Do("GET", key))
	//_redisLock.Unlock()
	return &result
}

//List
func (pool *RedisPool) MsgRedisListRightAppend(key string, content []byte) {
	//_redisLock.Lock()
	pool.RedisCon.Do("RPUSH", key, content)
	//_redisLock.Unlock()
}
func (pool *RedisPool) MsgRedisListLeftPop(key string) *[]byte {
	//_redisLock.Lock()
	result, _ := redis.Bytes(pool.RedisCon.Do("LPOP", key))
	//_redisLock.Unlock()
	if result == nil {
		return nil
	}
	return &result
}
func (pool *RedisPool) MsgRedisListCount(key string) int {
	//_redisLock.Lock()
	result, _ := redis.Int(pool.RedisCon.Do("LLEN", key))
	//_redisLock.Unlock()
	return result
}
func (pool *RedisPool) MsgRedisListRemove(key string) {
	//_redisLock.Lock()
	pool.RedisCon.Do("DEL", key)
	//_redisLock.Unlock()
}
func (pool *RedisPool) ListRightPush(key string, content string) {
	//_redisLock.Lock()
	pool.RedisCon.Do("RPUSH", key, content)
	//_redisLock.Unlock()
}
func (pool *RedisPool) ListLefPush(key string, content string) {
	//_redisLock.Lock()
	pool.RedisCon.Do("LPUSH", key, content)
	//_redisLock.Unlock()
}
func (pool *RedisPool) ListRightPop(key string) *string {
	//_redisLock.Lock()
	result, _ := redis.String(pool.RedisCon.Do("RPOP", key))
	//_redisLock.Unlock()
	return &result
}
func (pool *RedisPool) ListLeftPop(key string) *string {
	//_redisLock.Lock()
	result, _ := redis.String(pool.RedisCon.Do("LPOP", key))
	//_redisLock.Unlock()
	return &result
}
func (pool *RedisPool) GetListIndexValue(key string, index int) *string {
	//_redisLock.Lock()
	result, _ := redis.String(pool.RedisCon.Do("LINDEX", key, index))
	//_redisLock.Unlock()
	return &result
}
func (pool *RedisPool) SetListIndexValue(key string, index int, val string) {
	//_redisLock.Lock()
	pool.RedisCon.Do("LSET", key, index, val)
	//_redisLock.Unlock()
}
func (pool *RedisPool) GetListRange(key string, start int, end int) []interface{} {
	//_redisLock.Lock()
	values, _ := redis.Values(pool.RedisCon.Do("LRANGE", key, "0", "100"))
	//_redisLock.Unlock()
	return values
}
func (pool *RedisPool) Pipeline(command string, key string, val []byte) {
	//_redisLock.Lock()
	_redis.Send(command, key, val)
	//_redisLock.Unlock()
}
func (pool *RedisPool) PipelineFlush() {
	_redis.Flush()
	_redis.Receive()
}

//hash
func (pool *RedisPool) MsgRedisHashSet(key string, field string, val string) {
	//_redisLock.Lock()

	pool.RedisCon.Do("hset", key, field, val)
	//_redisLock.Unlock()
}

func (pool *RedisPool) MsgRedisHashGet(key string, field string) *string {
	//_redisLock.Lock()

	result, _ := redis.String(pool.RedisCon.Do("hget", key, field))
	//_redisLock.Unlock()
	return &result
}
