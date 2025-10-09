package pipeline

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/10/9 上午10:44
* @Package:
 */

type model struct {
	pipeline upgradePipeline
	steps    int
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func NewtProgram() *tea.Program {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return tea.NewProgram(model{
		spinner:  s,
		pipeline: searchPipeline{},
	})
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.pipeline.Do(), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case pipelineMsg:
		if msg.stage == nil {
			m.done = true
			info := tea.Printf("%s %s", checkMark, m.pipeline.String())
			if msg.info != "" {
				m.done = false
				info = tea.Println(msg.info)
			}
			return m, tea.Sequence(tea.Sequence(info, tea.Quit))
		}
		previous := m.pipeline.String()
		m.pipeline = msg.stage
		return m, tea.Batch(
			tea.Printf("%s %s", checkMark, previous), m.pipeline.Do())

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return doneStyle.Render(fmt.Sprintf("success upgrade to gvm %s\n", m.pipeline.Version()))
	}

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog))
	pkgName := currentPkgNameStyle.Render(m.pipeline.String())
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render(pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog))
	gap := strings.Repeat(" ", cellsRemaining)
	return spin + info + gap + prog
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
