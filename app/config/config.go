/*
 * @Date: 2022-02-14 12:00:49
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-03-03 17:23:05
 * @FilePath: \zt-printer\app\config\config.go
 */
package config

import (
	"io/ioutil" //
	"log"
	"zt-printer/util/file"

	"github.com/spf13/viper" //配置
)

var c *conf

const configInit string = `
## 其他杂项配置
[app]
OpcServer = "localhost:2225" ##localhost:2225
PrinerIpPort = "192.168.1.10:9100"  ##192.168.1.10:9100 默认端口9100
PrintStart ="test.AAA"    ##开始打印位号
Lots ="test.AAA"    ##批次位号
Gross ="test.AAA"    ##毛重位号
Net ="test.AAA"    ##净重位号
Bucket ="test.AAA"    ##桶号位号

`

func Instance() *conf {
	if c == nil {
		InitConfig("./config.toml")
	}
	return c
}

type conf struct {
	App SettingConf
}

type SettingConf struct {
	OpcServer    string
	PrinerIpPort string
	PrintStart   string
	Lots         string
	Gross        string
	Net          string
	Bucket       string
}

func InitConfig(tomlPath ...string) {
	if len(tomlPath) > 1 {
		log.Fatal("配置路径数量不正确")
	}
	if file.CheckNotExist(tomlPath[0]) {
		err := ioutil.WriteFile(tomlPath[0], []byte(configInit), 0777)
		if err != nil {
			log.Fatal("无法写入配置模板", err.Error())
		}
		log.Fatal("配置文件不存在，已在根目录下生成示例模板文件【config.toml】，请修改后重新启动程序！")
	}
	v := viper.New()
	v.SetConfigFile(tomlPath[0])
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件读取失败: ", err.Error())
	}
	err = v.Unmarshal(&c) //将解析到c地址
	if err != nil {
		log.Fatal("配置解析失败:", err.Error())
	}
}
