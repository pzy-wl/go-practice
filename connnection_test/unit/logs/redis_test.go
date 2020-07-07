package logs

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yredis"
)

func Test_redis1(t *testing.T) {
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadJwt:          false,
		LoadMq:           false,
		LoadMongo:        false,
		//here
		LoadRedis:    true,
		LoadPg:       false,
		LoadEtcd:     false,
		LoadRabbitMq: false,
	}
	_, err := lib.InitModulesOfOptions(opt)
	if err != nil {
		log.Println(err)
		return
	}
	key := "pzy"
	ret, err := yredis.X.Set(key, "panzhenying1", time.Second*1000).Result()
	fmt.Println("---ret---", ret, "-----------")
	fmt.Println("---err---", err, "-----------")
	fmt.Println("------", "demo get", "-----------")
	str, err := yredis.X.Get(key).Result()
	fmt.Println("----err--", err, "-----------")
	log.Println("key value:", str)
}
