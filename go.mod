module hyl

go 1.15

require (
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.8.0
	github.com/lifei6671/gocaptcha v0.1.1
	github.com/prometheus/common v0.7.0
	github.com/vhaoran/go-practice v0.0.0-20200720062337-51835b1cf2c5
	github.com/vhaoran/vchat v1.9.9
	github.com/vhaoran/yiintf v0.0.0-20201202075015-3b1c41704035
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200927032502-5d4f70055728
	google.golang.org/grpc/examples v0.0.0-20201204235607-0d6a24f68a5f // indirect
	gopl.io v0.0.0-20200323155855-65c318dde95e
)

replace github.com/vhaoran/vchat v1.9.9 => ../../../work/vchat

// 处理etcd编译出错
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
