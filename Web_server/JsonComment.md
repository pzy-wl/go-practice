目前结论:
方法一   // 初始化请求变量结构
	formData := make(map[string]interface{})
	// 调用json包的解析，解析请求body
	json.Unmarshal(jsonData, &formData)
方法二 
     // 初始化请求变量结构
     	formData := make(map[string]interface{})
     	// 调用json包的解析，解析请求body
     	json.NewDecoder(r.Body).Decode(&formData)
     得出的formData没有差别,输出方式不同可能输出内容也有所不同(一个事直接输出formData,另一种是将其遍历输出)
  *.(type)只能在switch语句中使用
  //go1.8关闭服务器功能   待测试----
  	quit := make(chan os.Signal)
  	signal.Notify(quit, os.Interrupt)
go func() {

		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close Server", err)
		}
	}()