package official

import (
	"fmt"
	"testing"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/12 上午10:02
* @Package:
 */
const OfficialDownloadPageURL = "https://golang.google.cn/dl/"

func Test_Parse(t *testing.T) {
	r, err := NewRegistry(OfficialDownloadPageURL)
	if nil != err {
		panic(err)
	}
	versions, err := r.ArchivedVersions()
	if err != nil {
		panic(err)
	}
	for _, version := range versions {
		fmt.Println(version.String())
	}
}
