package mmap

import (
	"fmt"
	"os"
	"syscall"
)

func internalMmap(fileName string, size int64) (*[]byte, error) {
	fmt.Println("create file")
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	fmt.Println("extend file")
	_, err = f.WriteAt([]byte{byte(0)}, size - 1)

	if err != nil {
		return nil, err
	}

	fmt.Println("map memory")
	data, err := syscall.Mmap(int(f.Fd()), 0, 1<<8, syscall.PROT_WRITE, syscall.MAP_SHARED)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func internalUnmmap(data *[]byte) (error) {
	return syscall.Munmap(*data)
}