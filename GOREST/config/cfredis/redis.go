package cfredis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis ...
type Redis struct {
	URL  string
	pool *redis.Pool
}

// NewPool ...
// Make New Pool in Redis
func NewPool(server string, dbnum int) *Redis {
	redis := &Redis{
		pool: &redis.Pool{
			MaxIdle:     100,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server, redis.DialDatabase(dbnum))
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
	return redis
}

func (t *Redis) CleanupHook() {
	t.pool.Close()
}

func (t *Redis) Ping() error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func (t *Redis) Get(key string) ([]byte, error) {
	conn := t.pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func (t *Redis) Set(key string, value []byte) error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func (t *Redis) HGet(hkey string, key string) ([]byte, error) {
	conn := t.pool.Get()
	defer conn.Close()
	var data []byte
	data, err := redis.Bytes(conn.Do("HGET", hkey, key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func (t *Redis) HSet(hkey string, key string, value []byte) error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", hkey, key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func (t *Redis) HDelete(hkey string, key string) error {
	conn := t.pool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", hkey, key)
	return err
}

func (t *Redis) Exists(key string) (bool, error) {
	conn := t.pool.Get()
	defer conn.Close()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func (t *Redis) HExists(hkey string, key string) (bool, error) {
	conn := t.pool.Get()
	defer conn.Close()
	ok, err := redis.Bool(conn.Do("HEXISTS", hkey, key))
	if err != nil {
		return ok, fmt.Errorf("error checking if hkey %s, key %s hexists: %v", hkey, key, err)
	}
	return ok, err
}

func (t *Redis) Expire(key string, sec int) error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := redis.Bool(conn.Do("EXPIRE", key, sec))
	return err

}

func (t *Redis) Delete(key string) error {

	conn := t.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func (t *Redis) GetKeysScan(pattern string) ([]string, error) {
	conn := t.pool.Get()
	defer conn.Close()
	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}
		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}
	return keys, nil
}

func (t *Redis) Incr(key string) (int, error) {
	conn := t.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func (t *Redis) Zcount(key string) (int, error) {
	conn := t.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("ZCOUNT", key, "-inf", "+inf"))
}

func (t *Redis) Zrem(zset string, zkey int) error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := conn.Do("ZREM", zset, zkey)
	return err
}

func (t *Redis) Zadd(zset string, score int, zkey int) error {
	conn := t.pool.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", zset, score, zkey)
	return err
}
