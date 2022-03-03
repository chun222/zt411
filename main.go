/*
 * @Date: 2022-03-03 09:34:03
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-03-03 17:32:19
 * @FilePath: \zt-printer\main.go
 */
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
		time.Sleep(3 * time.Second) //时间太短可能导致未重置又开始打印
	}

}
