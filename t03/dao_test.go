package t03

import (
	"fmt"
	"log"
	"testing"

	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ypg"
)

var a = new(AbcDao)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           true,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_pg_cnt(t *testing.T) {
	bean := &Abc{
		Name: "333",
		Age:  "3333",
	}
	ypg.X.AutoMigrate(bean)

	db := ypg.X.Save(bean)
	fmt.Println("----", db.Error)
	fmt.Println("-----", db.RowsAffected)
	fmt.Println("-----", bean.Id)
}

func Test_get(t *testing.T) {
	//此处相当于service层
	//根据id查找
	abc, err := a.Get1(15)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("查到的记录是:", abc)

}
func Test_ypg(t *testing.T) {
	println("输出测试!")
	//println(ypg.X.DB())
	println("输出测试!")
}
func Test_auto1(t *testing.T) {
	//先从redis数据库中查找,如果不存在则在postgres中查找,并将其结果写入Redis(如果有)
	abc, err := a.GetAuto(13)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(abc)
}
func Test_update(t *testing.T) {
	bean1 := &Abc{
		Id:   13,
		Name: "李武",
		Age:  "13",
	}
	a.Update(bean1)
}
func Test_del(t *testing.T) {
	a.Rm(13)
}
func Test_list(t *testing.T) {
	//list, err := a.List("zhangsan")
	list1, err := a.ListAll()
	if err != nil {
		log.Println(err)
	}
	for _, i := range list1 {
		fmt.Println(i)
	}
}
func Test_insert(t *testing.T) {
	//添加记录
	bean := &Abc{
		Id:   122,
		Name: "张三",
		Age:  "13",
	}
	a.Insert(bean)
}
func Test_pg(t *testing.T) {
	pb := &ypage.PageBeanMap{
		PageNo:      0,
		RowsPerPage: 3,
		PagesCount:  0,
		RowsCount:   0,
	}
	pb, err := a.Page(pb)
	if err != nil {
		log.Println(err)
	}
	println("当前数据总条数有:", pb.RowsCount, "条")
	println("当前数据总页数:", pb.PagesCount, "页")

	//类型数组的获取-----已知类型
	data := pb.Data.([]*Abc)
	//从json中获取数据
	//data:=json.Unmarshal([]byte, &pb.Data)

	for k, i := range data {
		fmt.Printf("第 %d 条数据的id是 %d name是%s age是 %s \n", k+1, i.Id, i.Name, i.Age)
	}
	//r:=json.Unmarshal(pb.Data, &data)
	println("查到的数据分别是", pb.Data)
}
func Test_jsontest(t *testing.T) {

}
