//go:build windows && !appengine
// +build windows,!appengine

package color

import (
	"golang.org/x/sys/windows"
	"os"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/9 下午4:46
* @Package:
 */

func init() {
	// Opt-in for ansi color support for current process.
	// https://learn.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences#output-sequences
	var outMode uint32
	out := windows.Handle(os.Stdout.Fd())
	if err := windows.GetConsoleMode(out, &outMode); err != nil {
		return
	}
	outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(out, outMode)
}
