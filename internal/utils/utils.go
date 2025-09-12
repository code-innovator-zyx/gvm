package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
	"os/exec"
	"runtime"
)

/*
* @Author: zouyx
* @Email: 1003941268@qq.com
* @Date:   2025/9/10 上午11:33
* @Package:
 */

func Unique[T comparable](in []T) []T {
	m := make(map[T]struct{}, len(in))
	out := make([]T, 0, len(in))
	for _, v := range in {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Algorithm checksum algorithm type
type Algorithm string

const (
	// SHA256 algorithm using SHA-256 hash function.
	SHA256 Algorithm = "SHA256"
	// SHA1 algorithm using SHA-1 hash function.
	SHA1 Algorithm = "SHA1"
)

// VerifyFile validates file integrity against expected checksum.
func VerifyFile(algo Algorithm, expectedChecksum, filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var h hash.Hash
	switch algo {
	case SHA256:
		h = sha256.New()
	case SHA1:
		h = sha1.New()
	default:
		return errors.New("unsupported checksum algorithm")
	}

	if _, err = io.Copy(h, f); err != nil {
		return err
	}

	if expectedChecksum != hex.EncodeToString(h.Sum(nil)) {
		return errors.New("file checksum does not match the computed checksum")
	}
	return nil
}
func Symlink(oldname, newname string) (err error) {
	if runtime.GOOS == "windows" {
		// Windows 10下无特权用户无法创建符号链接，优先调用mklink /j创建'目录联接'
		if err = exec.Command("cmd", "/c", "mklink", "/j", newname, oldname).Run(); err == nil {
			return nil
		}
	}
	if err = os.Symlink(oldname, newname); err != nil {
		return err
	}
	return nil
}
