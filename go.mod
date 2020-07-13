module github.com/vhaoran/go-practice

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/gorilla/mux v1.6.2
	github.com/jinzhu/gorm v1.9.11
	github.com/lib/pq v1.1.1
	github.com/vhaoran/vchat v1.9.9
	google.golang.org/genproto v0.0.0-20190425155659-357c62f0e4bb
	gopkg.in/gin-gonic/gin.v1 v1.3.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
	xorm.io/core v0.7.2
)

replace github.com/vhaoran/vchat => ../vchat

replace github.com/vhaoran/vchatintf => ../vchatintf
