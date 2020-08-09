package utils

import "io/ioutil"

type FileUtils interface {
	Write(filename string, content string) error
}

type fileUtils struct {
}

func NewFileUtils() FileUtils {
	return fileUtils{}
}

func (fileUtils) Write(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}
