package core

import (
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/utils"
	"os"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/25 下午4:11
* @Package:
 */
type VersionFunc func(version string) error

var (
	UninstallVersion VersionFunc
	InstallVersion   VersionFunc
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
