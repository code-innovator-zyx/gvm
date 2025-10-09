package github

import (
	"fmt"
	"os"
	"testing"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/29 下午2:51
* @Package:
 */
func TestCheck(t *testing.T) {
	i, y, err := NewReleaseUpdater().CheckForUpdates()
	if err != nil {
		t.Error(err)
		return
	}
	if y {
		assert, err := i.FindAsset()
		if err != nil {
			t.Error(err)
			return
		}
		defer assert.Clean()
		os.Setenv("http_proxy", "127.0.0.1:7890")
		os.Setenv("https_proxy", "127.0.0.1:7890")
		fmt.Println(assert.Download())
		fmt.Println(assert.Unzip())
		assert.Install()
		fmt.Println(assert)
	} else {
		fmt.Println("not need to upgrade")
	}
}
