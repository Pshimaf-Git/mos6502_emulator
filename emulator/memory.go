package emulator

const MAX_MEM_SIZE = 1024 * 64 // 65536

type Memory struct {
	Data [MAX_MEM_SIZE]byte
}

func (mem *Memory) Get(index Word) byte {
	return mem.Data[index]
}
