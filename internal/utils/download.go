package utils

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/*
* @Author: zouyx
* @Email: zouyx@knowsec.com
* @Date:   2025/9/15 下午1:57
* @Package:
 */

// DownloadFile saves the remote resource to local file with progress support.
func DownloadFile(srcURL string, filename string, flag int, perm fs.FileMode) (size int64, err error) {
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

	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("Downloading "+filepath.Base(filename)),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionShowBytes(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			_, _ = fmt.Fprint(ansi.NewAnsiStdout(), "\n")
		}),
	)
	_ = bar.RenderBlank()
	dst := io.MultiWriter(f, bar)
	return io.Copy(dst, resp.Body)
}
