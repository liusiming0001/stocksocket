package cache

import (
	"github.com/garyburd/redigo/redis"
	"sync"
)

var Redis redis.Conn
var _redis redis.Conn
var _redisLock sync.Mutex

func InitRedis() {
	_redis, _ = redis.Dial("tcp", "127.0.0.1:6379")
	_redis.Do("AUTH", "123qwe")
}

//string存储
func StringRedisSet(key string, val string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("SET", key, val)
	_redisLock.Unlock()
}
func StringRedisGet(key string, db int) *string {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.String(_redis.Do("GET", key))
	_redisLock.Unlock()
	return &result
}

//List
func MsgRedisListRightAppend(key string, content []byte, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("RPUSH", key, content)
	_redisLock.Unlock()
}
func MsgRedisListLeftPop(key string, db int) *[]byte {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.Bytes(_redis.Do("LPOP", key))
	_redisLock.Unlock()
	if result == nil {
		return nil
	}
	return &result
}
func MsgRedisListCount(key string, db int) int {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.Int(_redis.Do("LLEN", key))
	_redisLock.Unlock()
	return result
}
func MsgRedisListRemove(key string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("DEL", key)
	_redisLock.Unlock()
}
func ListRightPush(key string, content string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("RPUSH", key, content)
	_redisLock.Unlock()
}
func ListLefPush(key string, content string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("LPUSH", key, content)
	_redisLock.Unlock()
}
func ListRightPop(key string, db int) *string {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.String(_redis.Do("RPOP", key))
	_redisLock.Unlock()
	return &result
}
func ListLeftPop(key string, db int) *string {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.String(_redis.Do("LPOP", key))
	_redisLock.Unlock()
	return &result
}
func GetListIndexValue(key string, index int, db int) *string {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.String(_redis.Do("LINDEX", key, index))
	_redisLock.Unlock()
	return &result
}
func SetListIndexValue(key string, index int, val string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("LSET", key, index, val)
	_redisLock.Unlock()
}
func GetListRange(key string, start int, end int, db int) []interface{} {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	values, _ := redis.Values(_redis.Do("LRANGE", key, "0", "100"))
	_redisLock.Unlock()
	return values
}
func Pipeline(command string, key string, val []byte, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Send(command, key, val)
	_redisLock.Unlock()
}
func PipelineFlush() {
	_redis.Flush()
	_redis.Receive()
}

//hash
func MsgRedisHashSet(key string, field string, val string, db int) {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	_redis.Do("hset", key, field, val)
	_redisLock.Unlock()
}

func MsgRedisHashGet(key string, field string, db int) *string {
	_redisLock.Lock()
	_redis.Do("SELECT", db)
	result, _ := redis.String(_redis.Do("hget", key, field))
	_redisLock.Unlock()
	return &result
}
