package utils

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

const (
	padding    = 1
	maxWidth   = 100
	speedQueue = 5
)

type progressMsg struct {
	ratio      float64
	speed      float64
	remain     time.Duration
	written    int64
	totalBytes int64
}
type model struct {
	progress   progress.Model
	err        error
	cancel     bool
	speed      float64
	remain     time.Duration
	written    int64
	totalBytes int64
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tea.KeyMsg:
		m.cancel = true
		return m, tea.Quit

	case progressMsg:
		m.speed = msg.speed
		m.remain = msg.remain
		m.written = msg.written
		m.totalBytes = msg.totalBytes
		return m, m.progress.SetPercent(msg.ratio)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func formatSpeed(speed float64) string {
	if speed >= 1024*1024 {
		return fmt.Sprintf("%.2f MB/s", speed/1024/1024)
	}
	return fmt.Sprintf("%.2f KB/s", speed/1024)
}

func formatETA(d time.Duration) string {
	min := int(d.Minutes())
	sec := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}

// 格式化文件大小
func formatSize(bytes int64) string {
	if bytes >= 1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(bytes)/1024/1024)
	}
	return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	if m.err != nil {
		return pad + "❌ download failed: " + m.err.Error() + "\n"
	}
	sizeInfo := fmt.Sprintf("%s/%s", formatSize(m.written), formatSize(m.totalBytes))
	return "\n" + pad + m.progress.View() + "\n\n" +
		pad + helpStyle(fmt.Sprintf("press any key to  cancel | speed：%s | ETA：%s | %s",
		formatSpeed(m.speed), formatETA(m.remain), sizeInfo))
}

type progressWriter struct {
	total        int64
	written      int64
	start        time.Time
	speedHistory []float64
	onProgress   func(progressMsg)
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.written += int64(n)

	now := time.Now()
	elapsed := now.Sub(pw.start).Seconds()
	if elapsed <= 0 {
		elapsed = 0.001
	}

	instSpeed := float64(n) / elapsed
	pw.speedHistory = append(pw.speedHistory, instSpeed)
	if len(pw.speedHistory) > speedQueue {
		pw.speedHistory = pw.speedHistory[1:]
	}

	var sum float64
	for _, s := range pw.speedHistory {
		sum += s
	}
	avgSpeed := sum / float64(len(pw.speedHistory))
	remain := time.Duration(float64(pw.total-pw.written)/avgSpeed) * time.Second

	if pw.onProgress != nil && pw.total > 0 {
		pw.onProgress(progressMsg{
			ratio:      float64(pw.written) / float64(pw.total),
			speed:      avgSpeed,
			remain:     remain,
			written:    pw.written,
			totalBytes: pw.total,
		})
	}

	pw.start = now
	return n, nil
}

func DownloadFile(srcURL, filename string, flag int, perm fs.FileMode) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, srcURL, nil)
	if err != nil {
		return 0, fmt.Errorf("resource(%s) download failed ==> %s", srcURL, err.Error())
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("resource(%s) download failed ==> %s", srcURL, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("URL %q is unreachable  ==> %d", srcURL, resp.StatusCode)
	}

	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return 0, fmt.Errorf("resource(%s) download failed ==> %s", srcURL, err.Error())
	}
	defer f.Close()

	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}
	p := tea.NewProgram(m, tea.WithAltScreen())

	pw := &progressWriter{
		total:        resp.ContentLength,
		start:        time.Now(),
		speedHistory: make([]float64, 0, speedQueue),
		onProgress:   func(msg progressMsg) { p.Send(msg) },
	}

	go func() {
		_, _ = io.Copy(io.MultiWriter(f, pw), resp.Body)
		p.Quit() // 下载完主动退出 TUI
	}()

	finalModel, err := p.Run()
	if finalModel.(model).cancel {
		os.Remove(filename)
		return 0, fmt.Errorf("cancel download %s", srcURL)
	}
	return resp.ContentLength, err
}
