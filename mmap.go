package mmap

func MapFile(fileName string, size int64) (*[]byte, error) {

	return internalMmap(fileName, size)
}

func UnmapFile(bytes *[]byte) {
	if bytes != nil {
		internalUnmmap(bytes)
	}
}