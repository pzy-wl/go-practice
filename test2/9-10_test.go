package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func Test_Path(t *testing.T) {
	l := []string{
		"../",
		"../../",
		"../../../",
		"../../../../",
		"../../../../../"}

	pwd, _ := os.Getwd()
	for _, v := range l {
		res := path.Join(pwd, v)
		println("结果是:", res)
		//	结果是逐层返回上一级目录,而不是简单地字符串相加,而是寻址操作
	}
}
func Test_path1(t *testing.T) {
	fmt.Println("Path操作-----------------")
	fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg")) //aa.jpg
	fmt.Println(path.Clean("c:\\file//abc///aa.jpg"))          //c:\file/abc/aa.jpg
	fmt.Println(os.Getwd())                                    //D:\Projects\GoPath\source\demo\syntax\path <nil>
	fmt.Println(path.Dir("http://www.baidu.com/aa/aaa.jpg"))   //http:/www.baidu.com/aa
	fmt.Println(path.Dir("c:/a/b/c/d.txt"))                    //c:/a/b/c
	fmt.Println(path.Dir("c:\\a/b.txt"))                       //c:\a
	fmt.Println(path.Ext("c:\\a/b.txt"))                       //.txt
	fmt.Println(path.IsAbs("c:/wind/aa/bb/b.txt"))             //false
	fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))         //c:/aa/bb/cc.txt
	isMatch, err := path.Match("c:/windows/*/", "c:/windows/system/")
	fmt.Println(isMatch, err)                            //true <nil>
	fmt.Println(path.Split("c:/windows/system/aaa.jpg")) //c:/windows/system/ aaa.jpg
	//FilePath操作
	fmt.Println("FilePath操作-----------------")
	fmt.Println(filepath.IsAbs("c:\\wind\\aa\\bb\\b.txt"))                 //true
	fmt.Println(filepath.Abs("."))                                         //D:\Projects\GoPath\source\demo\syntax\path <nil>
	fmt.Println(filepath.Base("c:\\aa\\baa.exe"))                          //baa.exe
	fmt.Println(filepath.Clean("c:\\\\aa/c\\baa.exe"))                     //c:\aa\c\baa.exe
	fmt.Println(filepath.Clean("aa/c\\baa.exe"))                           //aa\c\baa.exe
	fmt.Println(filepath.Dir("aa/c\\baa.exe"))                             //aa\c
	fmt.Println(filepath.EvalSymlinks("./path.exe"))                       //可以用来判断文件或文件夹是否存在。 //path.exe <nil>
	fmt.Println(filepath.Ext("./path.exe"))                                //.exe
	fmt.Println(filepath.FromSlash("c:\\windows\\aa//bb/cc//path.exe"))    //将路径中的\\更换为/  //c:\windows\aa\\bb\cc\\path.exe
	fmt.Println(filepath.ToSlash("c:\\windows\\aa/bb/cc/path.exe"))        //将路径中的/替换为\\   //c:/windows/aa/bb/cc/path.exe
	fmt.Println(filepath.VolumeName("c:\\windows\\"))                      //获取卷标   //c:
	fmt.Println(filepath.Glob("c:\\windows\\*.exe"))                       //获取所有c:\\windows\\目录下exe文件。
	fmt.Println(filepath.HasPrefix("c:\\aa\\bb", "c:\\"))                  //true
	fmt.Println(filepath.IsAbs("http://www.baidu.com/aa.jpg"))             //false
	fmt.Println(filepath.Join("a", "\\bb\\", "cc", "/d", "e\\", "ff.txt")) //a\bb\cc\d\e\ff.txt
	fmt.Println(filepath.Match("c:/windows/*/", "c:/windows/system/"))     //true <nil>
	fmt.Println(filepath.Rel("c:/windows", "c:/windows/system/"))          //取得第二参的路径中，相对于前面的路径的相对路径。  //system <nil>
	fmt.Println(string(filepath.Separator))                                // windows下返回\\
	fmt.Println(filepath.Split("c:/windows/system/abc.exe"))               //c:/windows/system/ abc.exe
	fmt.Println(filepath.SplitList("c:/windows/system/abc.exe"))           //[c:/windows/system/abc.exe]
	filepath.Walk("../../syntax", WalkFunc)
}
func WalkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println("File:", path, "IsDir:", info.IsDir(), "size:", info.Size())
	return nil
}
func Test_vchat(t *testing.T) {
	fileanme := os.Getenv("vchat_yml_file")
	println("文件名是", fileanme)
}
func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	count := 0
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
		}
		count++
	}
	fmt.Printf("目前有数据%d条", count)
	return err
}
func TestGetFileName(t *testing.T) {
	err := GetAllFile("/Users/ccc/work/yicms/dao/img")
	if err != nil {
		return
	}
}

func TestFileRename(t *testing.T) {
	//对于某目录下的所有文件进行重命名
	fileInfoList, err := ioutil.ReadDir("/Users/ccc/work/yicms/dao/img")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
		name := strings.Split(fileInfoList[i].Name(), "-")[1]
		os.Rename("/Users/ccc/work/yicms/dao/img/"+fileInfoList[i].Name(), "/Users/ccc/work/yicms/dao/img/"+name)
	}

}

func TestRame(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	//path=path+"/1-sffaf.txt"
	fmt.Println(path)
	println(strings.Split("1-sffaf.txt", "-")[1])
	os.Rename(path+"/1-sffaf.txt", path+"/"+strings.Split("1-sffaf.txt", "-")[1])
}

var result []string

func TestImgNum(t *testing.T) {
	//	用于找出多出来的48张图片去哪了
	//	step1: 将就json中的数据取出,存在数组中, 得到数组长度
	//	step2: 将数组去重后得到长度
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	pwd = path.Join(pwd, "../")
	fmt.Println(pwd + "/file/yicms.json")
	res := readJson(pwd + "/file/yicms.json")

	fmt.Printf("共有图片%d张", len(res))
	println()
	after := RemoveRepeatedElement(res)
	//after := RemoveRepeatedElement2(res)
	fmt.Printf("去重后共有图片%d张", len(after))
	println()
	for _, v := range after {
		fmt.Println(v)
	}
}
func RemoveRepeatedElement(arr []string) (newArr []string) {
	//数组或切片去重
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				//fmt.Println(i)
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
func RemoveRepeatedElement2(arr []string) (newArr []string) {
	//第二种方法数组去重
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
func readJson(filePath string) (result []string) {
	//读取json并且取出图片的url
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		//将数据分条
		s, err := buf.ReadString('\n')
		//content分离出来
		s = strings.Replace(s, "[", "*", -1)
		s = strings.Replace(s, "]", "*", -1)
		//分段
		l := strings.SplitAfterN(s, "{", -1)
		for _, v := range l {
			//根据关键字找内容
			if strings.Contains(v, "file") || strings.Contains(v, "uploadimages") || strings.Contains(v, "uploadimg") {
				l1 := strings.SplitAfterN(v, "content\":\"", -1)
				l2 := strings.Split(l1[1], "\"},{")
				l3 := strings.Split(l2[0], "\"}*}}")
				if !strings.Contains(l3[0], "http") {
					l3[0] = "https://www.jpkcnet.com" + l3[0]
				}
				result = append(result, l3[0])
			}
			//fmt.Println(v)
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}
	return result
}
