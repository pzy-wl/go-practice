# go-practice
以下工程可供參考
  github.com/vhaoran/vchat.git 
    該工程中有本地服務器配置參數
    及redis/postgres/mongo/etcd/rabbitMQ/emq
    的具體使用案例。
  github.com/vhaoran/goStudy.git
    这里有基本的go学习demo,及其它一些常见类库的使用

# 主要练习内容
  ## 基础语法部分
     基本数据类型
     int,int32,int64区别
     float32/flot64区别
     fmt输出内容到控制台
     string转int/int64
     strring转float/flot64
     for
     fmt.Println();
     fmt.SPrint
     fmt.Sprintf();
  ## go routine
    go func(){
    }()
  ## 函数作为参数（回调）
    go func(func f(a int)error){
        a(3)
    }  
  ## chan的用法 
     也解即可。
     
## 数据库及其它服务器部件的使用
 ## postgres数据库
    要求能使用CRUD,分面技术
 ## redis
    会用get/set
以及了解Redis可以存储什么数据
