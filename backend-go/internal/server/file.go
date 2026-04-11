package server

import (
	"os"
)

func RemoveFile(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(bytes []byte, filePath string) error {
	if err := RemoveFile(filePath); err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0400); err != nil {
		return err
	}
	return nil
}
