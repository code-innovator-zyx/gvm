package list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type VersionModel struct {
	list     list.Model
	keys     *keyMap
	index    int
	quitting bool
}

func NewVersionModel(items []list.Item, title title) VersionModel {
	keys := newKeyMap()
	if title == LOCAL {
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
	return VersionModel{list: versionList, keys: keys}
}

func (m VersionModel) Index() int {
	return m.index
}
func (m VersionModel) Init() tea.Cmd {
	return nil
}

func (m VersionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var version string

	if i, ok := m.list.SelectedItem().(versionItem); ok {
		version = i.Title()
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
			return m, m.list.NewStatusMessage(statusMessageStyle("install " + version))

		case key.Matches(msg, m.keys.uninstall):
			index := m.Index()
			m.list.RemoveItem(index)
			if len(m.list.Items()) == 0 {
				m.keys.uninstall.SetEnabled(false)
			}
			return m, m.list.NewStatusMessage(statusMessageStyle("uninstall " + version))
		case key.Matches(msg, m.keys.use):
			return m, m.list.NewStatusMessage(statusMessageStyle("use " + version))
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
