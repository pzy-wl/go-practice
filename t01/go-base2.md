# 语言基本特性
  ## interface{}数据类型
  从interface中取出来原始数据
  func f(i interface{}){
      i,ok := i.(int)
      fmt.Prihntln(ok,i)
  }
  f(3)
  ## inface{}中數據類型判斷 
  switch(i.(type)){
    case int,int32:
      fmt.println("int32")
    case float,float64:
      fmt.println("int32")
  }
  ## slive使用
  l := make([]int,0)
  l = append(l,3,4,5)
  for i,v :- range l{
     fmt.Println(v,i)
  }
  ## 指针及struct的声明
   type Test struct {
      A int
      b String
   }
   type TestChild struct {
      Test
      A int
      b String
   }
