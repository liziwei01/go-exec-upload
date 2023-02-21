/*
 * @Author: liziwei01
 * @Date: 2022-03-04 13:52:11
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-16 20:44:56
 * @Description: file content
 */
package redis

import (
	"context"
	"sync"
	"time"

	r "github.com/go-redis/redis"
	"github.com/gogf/gf/util/gconv"
)

var (
	// 初始化互斥锁
	mu sync.Mutex
)

type Client interface {
	// GetStr 获取value
	Get(ctx context.Context, key string) (value string, err error)
	// Set 将字符串值 value 关联到 key
	// 如果 key 已经存在， SET 将覆盖旧值 无视类型
	// 过期时间为 nanoseconds 纳秒
	Set(ctx context.Context, key string, value string, expireTime ...time.Duration) error
	// Del
	Del(ctx context.Context, keys ...string) error
	// Determine if a key exists
	Exists(ctx context.Context, keys ...string) (bool, error)
	// Expired
	// Expired(ctx context.Context, key string) (bool, error)

	connect(ctx context.Context) (*r.Client, error)

	name() string
	host() string
	port() string
	password() string
	dbname() int
}

type client struct {
	conf *Config
	db   *r.Client
}

func New(config *Config) Client {
	c := &client{
		conf: config,
	}
	return c
}

func (c *client) connect(ctx context.Context) (*r.Client, error) {
	var err error
	if c.db != nil {
		if err = c.db.Ping().Err(); err == nil {
			return c.db, nil
		}
	}
	mu.Lock()
	defer mu.Unlock()
	c.db, err = c.open()
	return c.db, err
}

func (c *client) open() (*r.Client, error) {
	var (
		db  *r.Client
		err error
	)
	// 内含 retry 2
	db = r.NewClient(&r.Options{
		Addr:     c.host() + ":" + c.port(),
		Password: c.password(),
		DB:       c.dbname(),
	})
	return db, err
}

func (c *client) name() string {
	return c.conf.Name
}

func (c *client) dbname() int {
	return c.conf.Redis.DB
}

func (c *client) host() string {
	return c.conf.Resource.Manual.Host
}

func (c *client) port() string {
	return gconv.String(c.conf.Resource.Manual.Port)
}

func (c *client) password() string {
	return c.conf.Redis.Password
}
