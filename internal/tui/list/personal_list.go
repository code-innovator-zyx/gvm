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
* @Email: 1003941268@qq.com
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

type Model struct {
	list     list.Model
	keys     *keyMap
	index    int
	local    bool
	quitting bool
}

func NewListModel(items []list.Item, title title) Model {
	keys := newKeyMap()
	isLocal := title == LOCAL
	if isLocal {
		keys.install.SetEnabled(false)
	}
	versionList := list.New(items, delegate{keys: keys}, 0, 0)
	versionList.Title = string(title)
	versionList.Styles.Title = titleStyle
	helpKeys := func() []key.Binding {
		return []key.Binding{
			keys.install, keys.uninstall, keys.use,
		}
	}
	versionList.AdditionalShortHelpKeys = helpKeys
	versionList.AdditionalFullHelpKeys = helpKeys
	return Model{list: versionList, keys: keys, local: title == LOCAL}
}

func (m Model) Index() int {
	return m.index
}
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var ite item

	if i, ok := m.list.SelectedItem().(item); ok {
		ite = i
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
			if ite.dirname != "" {
				return m, nil
			}
			err := core.InstallVersion(ite.version)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			ite.dirname = filepath.Join(consts.VERSION_DIR, ite.version)
			setCmd := m.list.SetItem(m.list.Index(), ite)
			statusCmd := m.list.NewStatusMessage(successMessageStyle("success install " + ite.version))
			return m, tea.Batch(setCmd, statusCmd)

		case key.Matches(msg, m.keys.uninstall):
			if ite.currentUse {
				return m, m.list.NewStatusMessage(warnMessageStyle("can not uninstall current used version " + ite.version))
			}
			if ite.dirname == "" {
				return m, nil
			}
			err := core.UninstallVersion(ite.dirname)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			if m.local {
				m.list.RemoveItem(m.list.Index())
			} else {
				ite.dirname = ""
				m.list.SetItem(m.list.Index(), ite)
			}
			return m, m.list.NewStatusMessage(successMessageStyle("success uninstall " + ite.version))
		case key.Matches(msg, m.keys.use):
			if ite.currentUse || ite.dirname == "" {
				return m, nil
			}
			err := core.SwitchVersion(ite.dirname)
			if err != nil {
				return m, m.list.NewStatusMessage(failMessageStyle(err.Error()))
			}
			for i, v := range m.list.Items() {
				if vi := v.(item); vi.currentUse {
					vi.currentUse = false
					m.list.SetItem(i, vi)
				}
			}
			ite.currentUse = true
			setCmd := m.list.SetItem(m.list.Index(), ite)
			statusCmd := m.list.NewStatusMessage(successMessageStyle("current use " + ite.version))
			return m, tea.Batch(setCmd, statusCmd)
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return "\n"
	}
	return appStyle.Render(m.list.View())
}
