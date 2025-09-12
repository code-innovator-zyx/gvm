package consts

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/10 下午5:38
* @Package:
 */
var (
	GVM_HOME    string
	GO_ROOT     string
	VERSION_DIR string
)

func init() {
	homeDir, _ := os.UserHomeDir()
	GVM_HOME = filepath.Join(homeDir, ".gvm")
	GO_ROOT = filepath.Join(GVM_HOME, "go")
	VERSION_DIR = filepath.Join(GVM_HOME, "sdk")
}

const (
	// config keys
	CONFIG_MIRROR = "mirror"
	CONFIG_GOROOT = "goroots"

	EMPTY_INFO     = "<set-correct-info>"
	DEFAULT_MIRROR = "https://golang.google.cn/dl/"
	DEFAULT_GOROOT = "/.gvm/versions"
)

type VersionKind string

const (
	All      VersionKind = "all"
	Stable   VersionKind = "stable"
	Unstable VersionKind = "unstable"
	Archived VersionKind = "archived"
)

func ParseVersionKind(s string) (VersionKind, error) {
	switch s {
	case string(Stable):
		return Stable, nil
	case string(Unstable):
		return Unstable, nil
	case string(Archived):
		return Archived, nil
	case string(All):
		return All, nil
	default:
		return "", fmt.Errorf("invalid version kind: %s, must be stable | unstable | archived", s)
	}
}
