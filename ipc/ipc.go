package ipc

type Pipe struct {
	processTable map[int]int
	memTable     map[int]int
	str          [32]byte
	in           [1024]byte
	out          [1024]byte
}

type MemoryPool struct {
	Max  int
	Size int
	Used map[int][64]byte
	Free map[int][64]byte
}

func (mp *MemoryPool) Return(p int) {
	delete(mp.Used, p)
}

func (mp *MemoryPool) Borrow() [64]byte {
	mp.Used[2] = mp.Free[0]
	delete(mp.Free, 0)
	if len(mp.Free) < 10 {
	}
	return [64]byte{}
}
