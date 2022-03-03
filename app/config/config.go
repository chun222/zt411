/*
 * @Date: 2022-02-14 12:00:49
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-03-03 14:41:17
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
OpcServer = "localhost:2225"
PrinerIpPort = "192.168.1.10:9100"
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
