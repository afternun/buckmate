package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDirectory(scrDir string, dest string) error {
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		// stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		// if !ok {
		// return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		// }

		// fileMode := fileInfo.Mode() & os.ModeType

		if fileInfo.IsDir() {
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		//TODO: Verify if needed
		// if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
		// return err
		// }

	}
	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}
