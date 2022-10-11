package utils

import (
	"errors"
	"os"
	"path"
	"strings"
)

// FileExt 获取文件后缀
func FileExt(fname string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(fname), "."))
}

func FileExist(fpath string) bool {
	_, err := os.Stat(fpath) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func FileDirCreate(fDir string) error {
	f, err := os.Stat(fDir)
	if err == nil && f.IsDir() {
		return nil
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(fDir, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if f != nil && !f.IsDir() {
		return errors.New("exist a file, cannot cover a file")
	}
	return err
}

func FileNameFilter(name string) string {
	return strings.TrimSpace(strings.NewReplacer(
		`  `, ` `,
		`   `, ` `,
		`\`, ``,
		`/`, ``,
		`:`, ``,
		`*`, ``,
		`?`, ``,
		`"`, ``,
		`<`, ``,
		`>`, ``,
		`|`, ``,
		"\t", ``,
		"\n", ``,
		"\v", ``,
		"\f", ``,
		"\r", ``).Replace(name))
}
