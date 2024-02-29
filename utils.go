package main

import "os"

// exists returns whether the given file or directory exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
