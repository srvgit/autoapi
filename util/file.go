package util

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyDir(src string, dst string, extPattern string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			if err = os.MkdirAll(dstPath, fileInfo.Mode()); err != nil {
				return err
			}
			if err = CopyDir(srcPath, dstPath, extPattern); err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(entry.Name(), extPattern) {
				if err = copyFile(srcPath, dstPath); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
