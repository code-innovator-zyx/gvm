package utils

import (
	"os"
	"testing"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/15 下午2:05
* @Package:
 */

func Test_Download(t *testing.T) {
	sourceUrl := "https://golang.google.cn/dl/go1.24.6.darwin-arm64.tar.gz"
	download, err := DownloadFile(sourceUrl, "./go1.24.6.darwin-arm64.tar.gz", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(download)
}
