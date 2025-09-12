package pkg

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/12 下午1:54
* @Package:
 */

import (
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
* @Email: zouyx@knowsec.com
* @Date:   2025/9/11 下午3:45
* @Package:
 */

type VManager interface {
	List(kind consts.VersionKind) ([]*version.Version, error)
}

func NewManager(r bool) VManager {
	if r {
		return &remote{}
	}
	return &local{}
}

type local struct {
}

func (l local) List(kind consts.VersionKind) ([]*version.Version, error) {
	goRoots := viper.GetStringSlice(consts.CONFIG_GOROOT)
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
			v.Installed = true
			v.Path = root
			versions = append(versions, v)
		}
	}
	return versions, nil
}

type remote struct{}

func (r remote) List(kind consts.VersionKind) (versions []*version.Version, err error) {
	rg, err := registry.NewRegistry()
	if err != nil {
		return nil, err
	}
	installVersions, _ := local{}.List(kind)
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
	for i, v := range versions {
		for _, installVersion := range installVersions {
			if v.Equal(installVersion) {
				versions[i].Installed = true
				versions[i].Path = installVersion.Path
			}
		}
	}
	return versions, err
}

func SwitchVersion(version *version.Version) error {
	targetV := filepath.Join(version.Path, version.String())
	// Recreate symbolic link
	_ = os.Remove(consts.GO_ROOT)
	if err := utils.Symlink(targetV, consts.GO_ROOT); err != nil {
		return err
	}
	if output, err := exec.Command(filepath.Join(consts.GO_ROOT, "bin", "go"), "version").Output(); err == nil {
		fmt.Printf("Now using %s", strings.TrimPrefix(string(output), "go version "))
	}
	return nil
}
