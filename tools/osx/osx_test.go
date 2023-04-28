package osx

import "testing"

func TestCopy(t *testing.T) {
	Copy("D:\\a.txt", "D:\\b.txt")
}

func TestMove(t *testing.T) {
	Move("D:\\a.txt", "D:\\my\\a.txt")
}

func TestRemove(t *testing.T) {
	Remove("D:\\my\\w.txt")
}

func TestRename(t *testing.T) {
	Rename("D:\\my\\a.txt", "D:\\my\\w.txt")
}
