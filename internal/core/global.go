package core

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/utils"
	"io"
	"os"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/25 下午4:11
* @Package:
 */

var (
	UninstallVersion func(version string) error
	InstallVersion   func(version string) error
	InstallVersion2  func(versionName string, writer io.Writer, fn func(int642 int64)) error
)

func SwitchVersion(versionDir string) error {
	os.Remove(consts.GO_ROOT)
	_, err := os.Stat(versionDir)
	if err != nil {
		return err
	}
	if err = utils.Symlink(versionDir, consts.GO_ROOT); err != nil {
		return err
	}
	return nil
}

var (
	NewSpinnerProgram    func(options ...tea.ProgramOption) *tea.Program
	NewSimpleListProgram func(items []string, title string, options ...tea.ProgramOption) *tea.Program
)
