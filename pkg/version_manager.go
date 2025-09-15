package pkg

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/12 下午1:54
* @Package:
 */

import (
	"errors"
	"fmt"
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/registry"
	"github.com/code-innovator-zyx/gvm/internal/utils"
	"github.com/code-innovator-zyx/gvm/internal/version"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/11 下午3:45
* @Package:
 */

type VManager interface {
	// List 列出所有版本号
	List(kind consts.VersionKind) ([]*version.Version, error)
	// Install 安装指定版本号到本地
	Install(versionName string) error
	// Uninstall 从本地卸载指定版本号
	Uninstall(versionName string) error
}

type ManagerOption struct {
	WithLocal bool
}

func WithLocal() func(option *ManagerOption) {
	return func(option *ManagerOption) {
		option.WithLocal = true
	}
}

// NewManager
// r
func NewManager(r bool, opts ...func(option *ManagerOption)) VManager {
	opt := &ManagerOption{}
	for _, o := range opts {
		o(opt)
	}
	if r {
		return &remote{withLocal: opt.WithLocal}
	}
	return &local{}
}

type local struct {
}

func (l local) List(kind consts.VersionKind) ([]*version.Version, error) {
	goRoots := viper.GetStringSlice(consts.CONFIG_GOROOT)
	currentVersion := l.currentUsedVersion()
	if len(goRoots) == 0 {
		return nil, nil
	}
	var versions []*version.Version
	for _, root := range goRoots {
		versionDirs, err := os.ReadDir(root)
		if err != nil {
			continue
		}
		for _, versionDir := range versionDirs {
			if !versionDir.IsDir() {
				continue
			}
			v, err := version.NewVersion(strings.TrimPrefix(versionDir.Name(), "go"))
			if err != nil || v == nil {
				continue
			}
			v.CurrentUsed = v.String() == currentVersion
			v.Installed = true
			v.Path = root
			v.DirName = versionDir.Name()
			versions = append(versions, v)
		}
	}
	return versions, nil
}
func (l local) Install(versionName string) error {
	return errors.New("not support")
}
func (l local) currentUsedVersion() string {
	p, _ := os.Readlink(consts.GO_ROOT)
	versionName := filepath.Base(p)
	return strings.TrimPrefix(versionName, "go")
}

func (l local) Uninstall(versionName string) error {
	if versionName == l.currentUsedVersion() {
		return fmt.Errorf("cannot uninstall version %s: it is currently in use\n", versionName)
	}
	targetDir := filepath.Join(consts.VERSION_DIR, fmt.Sprintf("go%s", versionName))
	if finfo, err := os.Stat(targetDir); err != nil || !finfo.IsDir() {
		return fmt.Errorf("version %q is not installed\n", versionName)
	}

	if err := os.RemoveAll(targetDir); err != nil {
		return fmt.Errorf("uninstall failed: %s\n", err.Error())
	}
	return nil
}

type remote struct {
	withLocal bool
}

func (r remote) mergeInstalled(remoteVers []*version.Version, localVers []*version.Version) {
	m := make(map[string]*version.Version)
	for _, v := range localVers {
		m[v.Original()] = v // 用原始版本号做 key
	}
	for _, v := range remoteVers {
		if lv, ok := m[v.Original()]; ok {
			v.Installed = true
			v.CurrentUsed = lv.CurrentUsed
			v.Path = lv.Path
		}
	}
}

func (r remote) List(kind consts.VersionKind) (versions []*version.Version, err error) {
	rg, err := registry.NewRegistry()
	if err != nil {
		return nil, err
	}
	switch kind {
	case consts.Stable:
		versions, err = rg.StableVersions()
	case consts.Unstable:
		versions, err = rg.UnstableVersions()
	case consts.Archived:
		versions, err = rg.ArchivedVersions()
	default:
		versions, err = rg.AllVersions()
	}
	if r.withLocal {
		installVersions, _ := local{}.List(kind)
		r.mergeInstalled(versions, installVersions)
	}
	return versions, err
}

func (r remote) Install(versionName string) error {
	versions, err := (&remote{withLocal: false}).List(consts.All)
	if err != nil {
		return err
	}
	v, err := version.NewFinder(versions).Find(versionName)
	if err != nil {
		return err
	}
	artifact, err := v.FindArtifact()
	if nil != err {
		return err
	}
	return artifact.Install(versionName)
}

func (r remote) Uninstall(version string) error {
	//TODO implement me
	return errors.New("not support")
}

/*
*
软连接go指定版本的本地目录
*/
func SwitchVersion(version *version.Version) error {
	_ = os.Remove(consts.GO_ROOT)
	if err := utils.Symlink(version.LocalDir(), consts.GO_ROOT); err != nil {
		return err
	}
	if output, err := exec.Command(filepath.Join(consts.GO_ROOT, "bin", "go"), "version").Output(); err == nil {
		fmt.Printf("Now using %s", strings.TrimPrefix(string(output), "go version "))
	}
	return nil
}

func LocalInstalled(versionName string) bool {
	installVersions, _ := local{}.List(consts.All)
	for _, installVersion := range installVersions {
		if installVersion.String() == versionName {
			return true
		}
	}
	return false
}
