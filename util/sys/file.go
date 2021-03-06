/*
 * @Date: 2022-02-22 11:29:54
 * @LastEditors: 春贰
 * @Desc:
 * @LastEditTime: 2022-03-03 14:03:04
 * @FilePath: \zt-printer\util\sys\file.go
 */
package sys

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"zt-printer/util/str"
)

func RealPath(f string) string {
	p, err := filepath.Abs(f)
	if err != nil {
		log.Panicln("Get absolute path error.")
	}
	p = strings.Replace(p, "\\", "/", -1)
	l := strings.LastIndex(p, "/") + 1
	return str.Substr(p, 0, l)
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsFile(f string) bool {
	b, err := os.Stat(f)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return false
	}
	return true
}

func IsDir(p string) bool {
	b, err := os.Stat(p)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return true
	}
	return false
}
