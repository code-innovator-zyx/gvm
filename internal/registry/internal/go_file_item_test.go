// Copyright (c) 2019 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package internal

import (
	"fmt"
	"testing"
)

func Test_getGoVersion(t *testing.T) {
	items := []*struct {
		In       *GoFileItem
		Expected string
	}{
		{
			In: &GoFileItem{
				FileName: "go1.18beta1.darwin-amd64.pkg",
				URL:      "https://mirrors.aliyun.com/golang/go1.18beta1.darwin-amd64.pkg",
				Size:     "136.9 MB",
			},
			Expected: "1.18beta1",
		},
		{
			In: &GoFileItem{
				FileName: "go1.18beta2.freebsd-386.tar.gz",
				URL:      "https://mirrors.aliyun.com/golang/go1.18beta2.freebsd-386.tar.gz",
				Size:     "107.2 MB",
			},
			Expected: "1.18beta2",
		},
		{
			In: &GoFileItem{
				FileName: "go1.18rc1.darwin-amd64.tar.gz",
				URL:      "https://mirrors.aliyun.com/golang/go1.18rc1.darwin-amd64.tar.gz",
				Size:     "137.0 MB",
			},
			Expected: "1.18rc1",
		},
		{
			In: &GoFileItem{
				FileName: "go1.18.windows-arm64.zip",
				URL:      "https://mirrors.aliyun.com/golang/go1.18.windows-arm64.zip",
				Size:     "118.0 MB",
			},
			Expected: "1.18",
		},
		{
			In: &GoFileItem{
				FileName: "go1.18.1.linux-386.tar.gz",
				URL:      "https://mirrors.aliyun.com/golang/go1.18.1.linux-386.tar.gz",
				Size:     "107.6 MB",
			},
			Expected: "1.18.1",
		},
		{
			In: &GoFileItem{
				FileName: "go1.18.1.src.tar.gz",
				URL:      "https://mirrors.aliyun.com/golang/go1.18.1.src.tar.gz",
				Size:     "21.8 MB",
			},
			Expected: "1.18.1",
		},
	}
	t.Run("从文件名中获取go版本号", func(t *testing.T) {
		for _, item := range items {
			fmt.Println(item.In.getGoVersion())
		}
	})
}
