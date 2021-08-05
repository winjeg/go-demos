package ipc

import (
	"fmt"
	"github.com/winjeg/go-commons/str"
	"log"
	"syscall"
	"time"
	"unsafe"
)

func Write() {
	file, _ := syscall.UTF16PtrFromString("ShareMemory")
	size := 65536 // I’ve tried unsafe.Sizeof(MumbleData{}) but that didn’t work.
	handle, err := syscall.CreateFileMapping(0, nil, syscall.PAGE_READWRITE, 0, uint32(size), file)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.CloseHandle(handle)

	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_WRITE, 0, 0, 0)

	if err != nil {
		log.Fatal(err)
	}

	for {
		td := []byte(str.RandomAlphabetsLower(4))
		data := (*Pipe)(unsafe.Pointer(addr))
		for i, v := range td {
			data.out[i] = v
		}
		time.Sleep(1 * time.Second)
		// fmt.Printf("ava %v cam %v id %v\n", data.Avatar.Position, data.Camera, data.Identity)
		fmt.Printf("str: %s\n", string(data.str[:]))
	}

}
