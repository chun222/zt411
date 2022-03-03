# golang实现wincc直接驱动斑马打印机zt411打印标签

#### 介绍
后台服务通过opc方式检测wincc变量变化，自动出标签

其中读取wincc变量是用的我另一个c#写的opc服务；本项目主要是go实现斑马打印

## 关键代码
 
```go
	var err error
	var printstr string

	conn, err = net.Dial("tcp", config.Instance().App.PrinerIpPort)

	if err != nil {
		fmt.Println("err :", err)
		return
	}

	//打印wb
	err, printstr = tempaltefile(_PrintClass, "./template/wb.txt")
	if err != nil {
		fmt.Println("tempaltefile failed, err:", err)
		return
	}
	fmt.Println(printstr)
	_, err = conn.Write([]byte(printstr)) //发送打印zpl文本

	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	} 

	fmt.Println("打印成功!")
	defer conn.Close() // 关闭连接

```