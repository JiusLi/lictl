package osx

import (
	"io"
	"os"
)

// Copy 复制文件
func Copy(source, target string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

// Move 移动文件
func Move(source, target string) error {
	err := os.Rename(source, target)
	if err != nil {
		return err
	}
	// 删除源文件
	Remove(source)
	return nil
}

// Rename 重命名文件
func Rename(old, new string) error {
	err := os.Rename(old, new)
	if err != nil {
		return err
	}
	return nil
}

// Remove 删除文件
func Remove(source string) error {
	err := os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}
