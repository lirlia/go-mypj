package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Config struct {

	// 接続先のインスタンスのIPアドレス群です。
	Instances []string

	// 接続先のインスタンスのポートです。
	Port string

	// Redisへデータを格納する際に使用するKeyのPrefixです。
	// 例: NamespaceがExampleでTestというKeyをSetする際は Example:Test となります
	Namespace string

	// プール内のアイドル接続の最大数です。
	MaxIdle int

	// 特定の時間にプールによって割り当てられた接続の最大数です。
	// ゼロの場合、プール内の接続数に制限はありません。
	MaxActive int

	// この間アイドル状態を維持した後、接続を閉じます。
	// 値がゼロの場合、アイドル状態の接続は閉じられません。
	IdleTimeout time.Duration

	// 一定時間経過した接続を閉じます。
	// 値がゼロの場合プールは時間で接続を閉じません。
	MaxConnLifetime time.Duration

	// Read時のタイムアウトですを指定します。
	ReadTimeout time.Duration

	// Write時のタイムアウトですを指定します。
	WriteTimeout time.Duration

	// Redisサーバーに接続するためのタイムアウトを指定します。
	DialConnectTimeout time.Duration
}

// 複数のRedisの接続情報の構造体です
type RedisPools struct {
	Pool []*redis.Pool
}

// Redisインスタンスへの接続を全てクローズします
func (pools *RedisPools) Close() error {
	for _, pool := range pools.Pool {
		if err := pool.Close(); err != nil {
			return err
		}
	}
	return nil
}

// newPoolを呼び出しConfigに指定された(複数の)Redisの接続情報を返却します
func NewPool(c *Config) (*RedisPools, error) {

	pools := &RedisPools{}

	// コンフィグに記載されたRedisインスタンス数でループを行い
	// 対象のRedisインスタンスへの接続を pools に追加する
	for _, instance := range c.Instances {
		pool, err := newPool(instance, c)
		if err != nil {
			return nil, err
		}

		pools.Pool = append(pools.Pool, pool)
	}

	return pools, nil
}

// 指定のRedis Instanceに対する接続を作成
func newPool(host string, c *Config) (*redis.Pool, error) {

	// Redis接続時のOptionを作成
	options := []redis.DialOption{
		redis.DialReadTimeout(c.ReadTimeout),
		redis.DialWriteTimeout(c.WriteTimeout),
		redis.DialConnectTimeout(c.DialConnectTimeout),
	}

	// Redisへの接続インスタンスを作成
	addr := host + c.Port
	redis_pool := &redis.Pool{
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     c.IdleTimeout,
		MaxConnLifetime: c.MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, options...)
		},
	}

	// Redisへの接続を取得
	conn := redis_pool.Get()
	defer conn.Close()

	// 接続したRedisへの疎通確認を行いレスポンスがなければエラーを返す
	_, err := conn.Do("PING")
	if err != nil {
		return nil, fmt.Errorf("can't connect(PING) to %s : %w", host, err)
	}

	return redis_pool, nil

}

// Redisインスタンスに値を格納します
func Set(pool *redis.Pool, c *Config, key string, value string) ([]byte, error) {
	conn := pool.Get()
	bytes, err := redis.Bytes(conn.Do("SET", c.Namespace+":"+key, value))

	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Redisインスタンスから値を取得します
func Get(pool *redis.Pool, c *Config, key string) error {
	conn := pool.Get()
	bytes, err := redis.Bytes(conn.Do("GET", c.Namespace+":"+key))

	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func preStart() {
	_, f, _, _ := runtime.Caller(0)
	bytes, _ := ioutil.ReadFile(f)
	fmt.Println("## source")
	fmt.Println()
	fmt.Println("```go")
	fmt.Println(string(bytes))
	fmt.Println("```")
	fmt.Println()
	fmt.Println("## Result")
	fmt.Println()
	fmt.Println("```sh")
}
func preEnd() { fmt.Println("```") }

func main() {

	preStart()
	defer preEnd()

	c := &Config{}
	c.Instances = append(c.Instances, "127.0.100.1")
	c.Instances = append(c.Instances, "127.0.100.2")
	c.Instances = append(c.Instances, "127.0.100.3")
	c.Namespace = "test"
	c.Port = ":6379"
	c.IdleTimeout = 240 * time.Second
	c.DialConnectTimeout = 2 * time.Second
	c.MaxActive = 0
	c.MaxConnLifetime = 2 * time.Second
	c.MaxIdle = 3
	c.ReadTimeout = 2 * time.Second
	c.WriteTimeout = 2 * time.Second

	pools, err := NewPool(c)
	if err != nil {
		fmt.Println(err)
	}

	for i, pool := range pools.Pool {
		random, _ := MakeRandomStr(10)
		_, _ = Set(pool, c, "test"+strconv.Itoa(i), random)
		_ = Get(pool, c, "test"+strconv.Itoa(i))
	}

	// r, err := redis.Int(conn.Do("GET", "temperature"))
	// if err != nil {
	// 	panic(err)

}
