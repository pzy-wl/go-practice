package t03

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/lib/yredis"
)

var db = newDB()

//todo
// to gorm
// to ypg.X

type (
	Abc struct {
		Id   int64
		Name string
		Age  string
	}

	AbcDao struct {
	}
)

func check(err error) {
	if err != nil {
		fmt.Println("过程出错!")
		//panic(err)
	}
}

func newDB() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", "localhost", "postgres", "test", "123456")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	db.DB().SetMaxOpenConns(500)
	db.LogMode(true)
	return db
}
func (r *AbcDao) Get(id int64) (*Abc, error) {
	//通过bean获取
	var a1, a2 Abc
	//last是按照主键查找最后一个,last是按照主键查找最后一个
	db.Last(&a1)
	fmt.Println(a1)
	db.First(a2, "id=?", id)
	return &a2, nil

}
func (r *AbcDao) Get1(id int64) (*Abc, error) {
	//通过id获取记录
	//ypg.X.Save();
	defer db.Close()
	sqlStr := "select * from abcs where Id=$1"
	//此函数返回值是啥?
	//res:=db.Exec(sqlStr, id)
	rows, err := db.DB().Query(sqlStr, id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for rows.Next() {
		var abc Abc
		err = rows.Scan(&abc.Id, &abc.Name, &abc.Age)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println(abc)
		return &abc, err
	}
	return nil, nil
}
func (r *AbcDao) Insert(bean *Abc) (int64, error) {
	//ypg.X.Save();

	defer db.Close()
	//这个X是在哪定义的其内容?如何定义的内容?
	//ypg.X.Save(bean)
	db.Save(bean)
	return 0, nil
}

func (r *AbcDao) GetAuto(id int64) (*Abc, error) {
	//查询数据,先在redis进行查询,若是查不到,则进入posegres,并且将查询到的结果写入redis
	//把每个数据库表当做一个hashtable, 表明为key, 字段名为filed
	v, err := yredis.CacheAutoGetH(new(Abc), id,
		func() (interface{}, error) {
			//回调函数
			log.Println("redis-get:", id)
			//如果在redis中获取失败,则转回常规方式(直接访问数据库)获取
			return r.Get(id)
		})
	if err != nil {
		return nil, err
	}
	u := v.(*Abc)
	return u, err
}

func (r *AbcDao) Update(bean *Abc) error {
	//通过bean来对数据进行更新
	//ypg.X.Save();

	defer db.Close()
	db.Save(bean)
	return nil
}

func (r *AbcDao) Rm(id int64) error {
	//ypg.X.Save();

	defer db.Close()
	bean := &Abc{Id: 13}
	ret := db.Delete(bean)

	//if err := ypg.DBErr(ret); err != nil {
	//	ylog.Debug("--------shopUserDao.go-->Exist  err----", err)
	//	return err
	//}

	return nil
}

func (r *AbcDao) Page(pb *ypage.PageBeanMap) (*ypage.PageBeanMap, error) {
	//ypg.X.Save();
	defer db.Close()
	abs := make([]*Abc, 0)
	//Find是查找所有的记录
	db.Find(&abs, "1=1")
	pb.RowsCount = int64(len(abs))
	pb.RowsPerPage = 3
	i := pb.RowsCount / pb.RowsPerPage
	if i != 0 {
		i += 1
	}
	pb.PagesCount = i
	pb.Data = abs
	pb.PageNo = 0

	return pb, nil
}

func (r *AbcDao) List(name string) ([]*Abc, error) {
	//模糊查询
	//ypg.X.Save();
	abs := make([]*Abc, 0)
	//Find是查找所有的记录
	db.Find(&abs, "name=?", name)
	return abs, nil
}
