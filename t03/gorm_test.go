package t03

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("过程出错!")
		//panic(err)
	}
}
func newDBCnt() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", "localhost", "postgres", "test", "123456")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxOpenConns(500)
	db.LogMode(true)
	return db
}
func Test_gorm1(t *testing.T) {
	//插入记录
	db := newDBCnt()
	defer db.Close()
	sqlStr := "INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid"
	db.Exec(sqlStr, "paznhenying", "开发", time.Now())
	fmt.Println()
	fmt.Println("----------insert done------------")
}
func Test_gorm2(t *testing.T) {
	//创建表
	bean := &Abc{}
	db := newDBCnt()
	defer db.Close()
	if db.HasTable(bean) {
		fmt.Println("该表已存在,正在删除----")
		if err := db.DropTable(bean).Error; err != nil {
			log.Println("err:", err)
			return
		}
		fmt.Println("删除成功!")
	}
	db.CreateTable(bean)
	db.AutoMigrate(bean)

}
func Test_gorm3(t *testing.T) {
	//查询记录
	db := newDBCnt()
	sqlStr := "select * from userinfo"
	defer db.Close()
	println(db.Exec(sqlStr))
}
func Test_gorm5(t *testing.T) {
	//多条件模糊查询
	db := newDBCnt()
	sqlStr := "select * from company where name like $1 and id>$2"
	defer db.Close()
	fmt.Println(db.Exec(sqlStr, "%m%", "2"))

}
func Test_gorm6(t *testing.T) {
	//根据类型写入数据
	db := newDBCnt()
	defer db.Close()
	for i := int64(0); i < 100; i++ {
		bean := &Abc{
			Id:   i,
			Name: fmt.Sprintf("aaa%d", i),
			Age:  "23",
		}
		if err := db.Save(bean).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("append ok")
		}
	}

}
func Test_gorm4(t *testing.T) {
	//更新记录
	db := newDBCnt()
	sqlStr := "update userinfo set username=$1 where uid=$2"
	defer db.Close()
	db.Exec(sqlStr, "pzy", "47")
}
