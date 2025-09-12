package autoindex

import (
	"fmt"
	"testing"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/12 上午10:02
* @Package:
 */
const OfficialDownloadPageURL = "https://mirrors.ustc.edu.cn/golang/"

func Test_Parse(t *testing.T) {
	r, err := NewRegistry(OfficialDownloadPageURL)
	if nil != err {
		panic(err)
	}
	versions, err := r.AllVersions()
	if err != nil {
		panic(err)
	}
	for _, version := range versions {
		fmt.Println(version.String())
		for _, artifact := range version.Artifacts {
			fmt.Println(artifact.OS, artifact.Arch, artifact.Kind, artifact.FileName, artifact.URL)
			fmt.Println()
		}
	}
}
