package main

import "os"

func initDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0o770)
		}
		return err
	}
	return nil
}
