package t03

import (
	"github.com/vhaoran/vchat/common/ypage"
)

type (
	Abc struct {
		Id   int64
		Name string
		Age  string
	}

	AbcDao struct {
	}
)

func (r *AbcDao) Insert(bean *Abc) (int64, error) {
	//ypg.X.Save();
	return nil, nil
}

func (r *AbcDao) Get(id int64) (*Abc, error) {
	//ypg.X.Save();
	return nil, nil
}

func (r *AbcDao) GetAuto(id int64) (*Abc, error) {
	//v, err := yredis.CacheAutoGetH(new(UserInfo), uid,
	//	func() (interface{}, error) {
	//		log.Println("redis-get:", uid)
	//		return r.Get(uid)
	//	})
	//if err != nil {
	//	return nil, err
	//}
	//
	//u := v.(*UserInfo)
	//u.Pwd = "*********"
	//return u, err
	return nil, nil
}

func (r *AbcDao) Update(bean *Abc) error {
	//ypg.X.Save();
	return nil
}

func (r *AbcDao) Rm(id int64) error {
	//ypg.X.Save();
	return nil
}

func (r *AbcDao) Page(pb *ypage.PageBeanMap) (*ypage.PageBeanMap, error) {
	//ypg.X.Save();
	return nil, nil
}

func (r *AbcDao) List(name string) ([]*Abc, error) {
	//ypg.X.Save();
	return nil, nil
}
