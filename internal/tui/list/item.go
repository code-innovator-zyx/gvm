package list

import "github.com/charmbracelet/bubbles/list"

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/25 上午10:42
* @Package:
 */

type versionItem struct {
	version    string
	dirname    string
	currentUse bool
}

func NewVersionItem(version, dirname string, used bool) list.Item {
	return versionItem{
		version:    version,
		dirname:    dirname,
		currentUse: used,
	}
}

func (i versionItem) Title() string       { return i.version }
func (i versionItem) Description() string { return i.dirname }
func (i versionItem) FilterValue() string { return i.version }
