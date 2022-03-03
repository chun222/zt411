/*
 * @Date: 2022-03-03 14:20:28
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-03-03 17:35:18
 * @FilePath: \zt-printer\app\print\print.go
 */

// tcp/server/main.go

// TCP server端
package print

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"text/template"
	"time"
	"zt-printer/app/config"
	"zt-printer/util/convert"

	//"time"

	"io/ioutil"
	"net/http"
	"strings"
)

var conn net.Conn

type DataResult struct {
	Code interface{}            `json:"code"`
	Msg  interface{}            `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type PrintClass struct {
	Lots     string
	Produced string
	Expired  string
	Gross    string
	Net      string
	Bucket   string
}

func Run() {
	startSign := config.Instance().App.PrintStart
	BucketSign := config.Instance().App.Bucket
	GrossSign := config.Instance().App.Gross
	LotsSign := config.Instance().App.Lots
	NetSign := config.Instance().App.Net

	err, data := readTags([]string{startSign, BucketSign, GrossSign, LotsSign, NetSign})
	obj := DataResult{}
	if err == nil {
		if err := json.Unmarshal([]byte(data), &obj); err != nil {
			fmt.Println("err>>", err)
		} else {
			if sign, err := convert.Float(obj.Data[startSign]); err != nil {
				fmt.Println("err>>", err)
			} else {
				fmt.Println("sign>>", sign)
				if sign == 1 {
					_PrintClass := PrintClass{}
					_PrintClass.Lots = convert.String(obj.Data[LotsSign])
					_PrintClass.Produced = time.Now().Format("2006-01-02 15:04:05")
					_PrintClass.Expired = time.Now().AddDate(1, 0, 0).Format("2006-01-02 15:04:05") //一年后时间
					_PrintClass.Gross = convert.String(obj.Data[GrossSign])
					_PrintClass.Net = convert.String(obj.Data[NetSign])
					_PrintClass.Bucket = convert.String(obj.Data[BucketSign])
					Ztprint(_PrintClass)
					//重置信号
					write(startSign, 0)
				}
			}
		}
	}
}

func write(tag string, value float64) {

	url := "http://" + config.Instance().App.OpcServer + "/api/opc/write"
	m := map[string]interface{}{"tags": []string{tag}, "values": []float64{value}}
	jsonbyte, _ := json.Marshal(m)
	jsonstr := string(jsonbyte)
	// jsonstr := `{ "tags":["` + tag + `"],"values":[` + fmt.Sprintf("%f", value) + `]}`
	sendpost(url, jsonstr)
}

func readTags(tags []string) (error, string) {
	url := "http://" + config.Instance().App.OpcServer + "/api/opc/read"

	m := map[string]interface{}{"tags": tags}
	jsonbyte, _ := json.Marshal(m)
	jsonstr := string(jsonbyte)
	return sendpost(url, jsonstr)
}

func Ztprint(_PrintClass PrintClass) {
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

	//打印nb
	err, printstr = tempaltefile(_PrintClass, "./template/nb.txt")
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
}

func tempaltefile(_PrintClass PrintClass, file string) (error, string) {

	tmpl, err := template.ParseFiles(file)
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, _PrintClass)

	return err, buf.String()
}

func sendpost(url string, jsonstr string) (error, string) {
	method := "POST"
	payload := strings.NewReader(jsonstr)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	return nil, string(body)
}
