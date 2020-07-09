package t03

import (
	"fmt"
	"log"
	"testing"

	"github.com/vhaoran/vchat/common/ypage"
)

var a = new(AbcDao)

func Test_Dao(t *testing.T) {
	//此处相当于service层
	//根据id查找
	abc, err := a.Get(13)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(&abc)
	//添加记录
	bean := &Abc{
		Id:   13,
		Name: "张三",
		Age:  "13",
	}
	a.Insert(bean)

}
func Test_auto1(t *testing.T) {
	//先从redis数据库中查找,如果不存在则在postges中查找,并将其结果写入Redis(如果有)
	abc, err := a.GetAuto(13)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(abc)
}
func Test_update(t *testing.T) {
	bean1 := &Abc{
		Id:   13,
		Name: "李四",
		Age:  "13",
	}
	a.Update(bean1)
}
func Test_del(t *testing.T) {
	a.Rm(13)
}
func Test_list(t *testing.T) {
	list, err := a.List("zhangsan")
	if err != nil {
		log.Println(err)
	}
	for _, i := range list {
		fmt.Println(i)
	}
}
func Test_pg(t *testing.T) {

	a.Page()
}
