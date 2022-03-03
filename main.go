/*
 * @Date: 2022-03-03 09:34:03
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-03-03 14:34:25
 * @FilePath: \zt-printer\main.go
 */
// tcp/server/main.go

// TCP server端
package main

import (
	"time"
	"zt-printer/app/config"
	"zt-printer/app/print"
)

func main() {
	config.Instance()
	for {
		go print.Run()
		time.Sleep(1 * time.Second)
	}

}
