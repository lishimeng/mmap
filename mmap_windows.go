package mmap

import "github.com/jeanphorn/log4go"

func internalMmap(fileName string, size int64) (*[]byte, error) {

	log4go.Debug("dummy mmap ", fileName)
	return nil, nil
}

func internalUnmmap(data *[]byte) (error) {

	log4go.Debug("dummy munmap ")
	return nil
}