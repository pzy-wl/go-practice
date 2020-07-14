module github.com/vhaoran/go-practice

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/jinzhu/gorm v1.9.11
	github.com/lib/pq v1.1.1
	github.com/vhaoran/vchat v1.9.9
)

replace github.com/vhaoran/vchat => ../vchat
