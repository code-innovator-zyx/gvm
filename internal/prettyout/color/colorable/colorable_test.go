package colorable

import (
	"os"
	"runtime"
	"testing"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/9 下午3:40
* @Package:
 */

func TestColorable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skipf("skip this test on windows")
	}
	_, ok := NewColorableStdout().(*os.File)
	if !ok {
		t.Fatalf("should os.Stdout on UNIX")
	}
	_, ok = NewColorableStderr().(*os.File)
	if !ok {
		t.Fatalf("should os.Stdout on UNIX")
	}
	_, ok = NewColorable(os.Stdout).(*os.File)
	if !ok {
		t.Fatalf("should os.Stdout on UNIX")
	}
}
