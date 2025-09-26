package list

import "github.com/charmbracelet/bubbles/list"

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/25 上午10:42
* @Package:
 */

type item struct {
	version    string
	dirname    string
	currentUse bool
}

func NewVersionItem(version, dirname string, used bool) list.Item {
	return item{
		version:    version,
		dirname:    dirname,
		currentUse: used,
	}
}

func (i item) Title() string       { return i.version }
func (i item) Description() string { return i.dirname }
func (i item) FilterValue() string { return i.version }
