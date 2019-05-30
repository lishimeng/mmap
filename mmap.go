package mmap

func MapFile(fileName string, size int64) (*[]byte, error) {

	return internalMmap(fileName, size)
}

func UnmapFile(datas *[]byte) {
	if datas != nil {
		internalUnmmap(datas)
	}
}