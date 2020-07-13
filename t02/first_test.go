package t02

import (
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
)

func Test_test(t *testing.T) {
	fmt.Println("hello, pzy!")
}
func Test_test2(t *testing.T) {
	fmt.Println("1weqw")
}

var MAX_POOL_SIZE = 20

var redisPool chan redis.Conn

func InitRedis(network, address string) redis.Conn {
	redisPool = make(chan redis.Conn, MAX_POOL_SIZE)
	if len(redisPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				c, err := redis.Dial(network, address)
				if err != nil {
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPool
}

func putRedis(conn redis.Conn) {
	if redisPool == nil {
		redisPool = make(chan redis.Conn, MAX_POOL_SIZE)
	}
	if len(redisPool) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	redisPool <- conn
}

func main() {
	fmt.Println()

	c := InitRedis("tcp", "192.168.13.200:6379")

	//test uuid
	fmt.Println(time.Now())
	startTime := time.Now()
	var Success, Failure int
	for i := 0; i < 100000; i++ {
		if ok, _ := redis.Bool(c.Do("HSET", "payVerify:session", uuid.New(), "aaaa")); ok {
			Success++
			// break
		} else {
			Failure++
		}
	}
	fmt.Println(time.Now())
	fmt.Println("用时：", time.Now().Sub(startTime), "总计：100000,成功：", Success, "失败：", Failure)
}
func Test_redis(t *testing.T) {
	host := "localhost" //控制台显示的地址
	port := 6379
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	//如果redis有密码,则在auth后面的参数输入密码,否则c.Do直接省略
	//_, err = c.Do("auth", "1234")
	//if err != nil {
	//	fmt.Println("redis auth failed: ", err)
	//	return
	//}
	_, err = c.Do("SET", "key", "jcloud-redis")
	if err != nil {
		fmt.Println("redis set failed: ", err)
		return
	}
	_, err = c.Do("GET", "key")
	if err != nil {
		fmt.Println("redis get failed: ", err)
		return
	}
	//do other command...
}
