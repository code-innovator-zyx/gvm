package list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/code-innovator-zyx/gvm/internal/consts"
	"github.com/code-innovator-zyx/gvm/internal/core"
	"path/filepath"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/25 上午10:46
* @Package:
 */
type title string

const (
	Remote title = "remote versions"
	LOCAL  title = "local versions"
)

var (
	appStyle = lipgloss.NewStyle().Padding(0, 0).Margin(1, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 0)

	successMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	warnMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#FFD700", Dark: "#FFD700"}).
				Render

	failMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#FF4B4B", Dark: "#FF4B4B"}).
				Render
)

type VersionModel struct {
	list     list.Model
	keys     *keyMap
	index    int
	local    bool
	quitting bool
}

func NewVersionModel(items []list.Item, title title) VersionModel {
	keys := newKeyMap()
	isLocal := title == LOCAL
	if isLocal {
		keys.install.SetEnabled(false)
	}
	versionList := list.New(items, versionDelegate{keys: keys}, 0, 0)
	versionList.Title = string(title)
	versionList.Styles.Title = titleStyle
	helpKeys := func() []key.Binding {
		return []key.Binding{
			keys.install, keys.uninstall, keys.use,
		}
	}
	versionList.AdditionalShortHelpKeys = helpKeys
	versionList.AdditionalFullHelpKeys = helpKeys
	return VersionModel{list: versionList, keys: keys, local: title == LOCAL}
}

func (m VersionModel) Index() int {
	return m.index
}
func (m VersionModel) Init() tea.Cmd {
	return nil
}

func (m VersionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var item versionItem

	if i, ok := m.list.SelectedItem().(versionItem); ok {
		item = i
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch {
		case key.Matches(msg, m.keys.install):
			if item.dirname != "" {
				return m, nil
			}
			err := core.InstallVersion(item.version)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			item.dirname = filepath.Join(consts.VERSION_DIR, item.version)
			setCmd := m.list.SetItem(m.list.Index(), item)
			statusCmd := m.list.NewStatusMessage(successMessageStyle("success install " + item.version))
			return m, tea.Batch(setCmd, statusCmd)

		case key.Matches(msg, m.keys.uninstall):
			if item.currentUse {
				return m, m.list.NewStatusMessage(warnMessageStyle("can not uninstall current used version " + item.version))
			}
			if item.dirname == "" {
				return m, nil
			}
			err := core.UninstallVersion(item.dirname)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			if m.local {
				m.list.RemoveItem(m.list.Index())
			} else {
				item.dirname = ""
				m.list.SetItem(m.list.Index(), item)
			}
			return m, m.list.NewStatusMessage(successMessageStyle("success uninstall " + item.version))
		case key.Matches(msg, m.keys.use):
			if item.currentUse || item.dirname == "" {
				return m, nil
			}
			err := core.SwitchVersion(item.dirname)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			for i, v := range m.list.Items() {
				if vi := v.(versionItem); vi.currentUse {
					vi.currentUse = false
					m.list.SetItem(i, vi)
				}
			}
			item.currentUse = true
			setCmd := m.list.SetItem(m.list.Index(), item)
			statusCmd := m.list.NewStatusMessage(successMessageStyle("current use " + item.version))
			return m, tea.Batch(setCmd, statusCmd)
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m VersionModel) View() string {
	if m.quitting {
		return "\n"
	}
	return appStyle.Render(m.list.View())
}
