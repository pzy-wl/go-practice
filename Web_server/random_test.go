package Web_server

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//随机生成名字
var (
	familiyNames = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳"}
	firstNames   = []string{"金", "木", "水", "火", "土", "春", "夏", "秋", "冬", "山", "石", "田", "天", "地", "玄", "黄", "宇", "宙", "洪", "荒"}
	//辈分
	generationNameMap = make(map[string][]string)
	names             = make([]string, 0)
)

func init() {
	generationNameMap["欧阳"] = []string{"宗", "的", "永", "其", "光"}
	for _, ln := range familiyNames {
		if ln != "欧阳" {
			generationNameMap[ln] = []string{"飞", "前", "茂", "百", "方", "书", "生", "无", "一", "用"}
		}
	}
}

func GetRandomName() (name string) {
	familyName := familiyNames[GetRandomInt(0, len(familiyNames)-1)]
	middleName := generationNameMap[familyName][GetRandomInt(0, len(generationNameMap[familyName])-1)]
	firstName := firstNames[GetRandomInt(0, len(firstNames)-1)]
	return familyName + middleName + firstName
}

//生成随机数
func GetRandomInt(start, end int) int {
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + r.Intn(end-start)
}

func Test_random1(t *testing.T) {
	for i := 0; i < 100; i++ {
		names = append(names, GetRandomName())
	}
	fmt.Println(names)

}
