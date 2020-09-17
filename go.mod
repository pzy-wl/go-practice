module hyl

go 1.15

require (
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.8.0
	github.com/vhaoran/go-practice v0.0.0-20200720062337-51835b1cf2c5
	github.com/vhaoran/vchat v1.9.9
)

replace github.com/vhaoran/vchat v1.9.9 => ../../../work/vchat
