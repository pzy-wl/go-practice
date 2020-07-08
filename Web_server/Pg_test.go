package Web_server

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func Test_pg1(t *testing.T) {
	/*插入*/
	//insert("panzhenying","研发", time.Now() )
	/*更新*/
	//	update("lizhihao", 16)
	/*删除id小于13的所有的用户信息*/
	del(13)
	//检索
	selectAll()
	//db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	//checkErr(err)
	//// 插入数据
	//fmt.Println("----------inserting------------")
	//stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	//checkErr(err)
	//
	//res, err := stmt.Exec("panzhenying123", "开发岗位", time.Now())
	//checkErr(err)
	//fmt.Println("----------insert done------------")
	////var lastInsertId int
	////err = db.QueryRow("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	////checkErr(err)
	////fmt.Println("最后插入id =", lastInsertId)
	//
	//// 更新数据
	//fmt.Println("----------updateing------------")
	//stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	//checkErr(err)
	//
	//res, err = stmt.Exec("panzhenying",11)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//fmt.Println("----------update done------------")
	//fmt.Printf("执行更新后数据有%d行受到影响\n",affect)
	//
	//// 查询数据
	//println("-----------selecting-------------------")
	//rows, err := db.Query("SELECT * FROM userinfo")
	//checkErr(err)
	//println("-----------result-------------------")
	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var department string
	//	var created string
	//	err = rows.Scan(&uid, &username, &department, &created)
	//	checkErr(err)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//	fmt.Println(department)
	//	fmt.Println(created)
	//}
	//
	//// 删除数据
	//println("-----------deleteing-------------------")
	//stmt, err = db.Prepare("delete from userinfo where uid=$1")
	//checkErr(err)
	//
	//res, err = stmt.Exec(13)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//println("-----------delete done-------------------")
	//fmt.Printf("执行删除操作后现有数据有%d行收到影响\n",affect)
	//fmt.Println(affect)
	//db.Close()

}
func insert(name string, de string, day time.Time) {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	checkErr(err)
	// 插入数据
	fmt.Println("----------inserting------------")
	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("panzhenying123123", "开发岗位", time.Now())
	checkErr(err)
	println(res)
	fmt.Println("----------insert done------------")
	//fmt.Printf("共有%d行被影响",res.RowsAffected())
}
func update(name string, uid int) {
	// 更新数据
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	checkErr(err)
	fmt.Println("----------updateing------------")
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec(name, uid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("----------update done------------")
	fmt.Printf("执行更新后数据有%d行受到影响\n", affect)
}
func selectAll() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	checkErr(err)
	// 查询数据
	println("-----------selecting-------------------")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	println("-----------result-------------------")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}
func del(num int) {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=test sslmode=disable")
	checkErr(err)
	// 删除数据
	println("-----------deleteing-------------------")
	stmt, err := db.Prepare("delete from userinfo where uid<$1")
	checkErr(err)

	res, err := stmt.Exec(num)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	println("-----------delete done-------------------")
	fmt.Printf("执行删除操作后现有数据有%d行收到影响\n", affect)
	fmt.Println(affect)
	db.Close()

}
func checkErr(err error) {
	if err != nil {
		fmt.Println("过程出错!")
		//panic(err)
	}
}
func Test_outputDate(t *testing.T) {
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now())
	fmt.Println(time.Local)

}
