package ipc

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)


func Read() {
	file, _ := syscall.UTF16PtrFromString("ShareMemory")
	size := 65536 // I’ve tried unsafe.Sizeof(MumbleData{}) but that didn’t work.
	handle, err := syscall.CreateFileMapping(0, nil, syscall.PAGE_READONLY, 0, uint32(size), file)

	if err != nil {
		log.Fatal(err)
	}
	defer syscall.CloseHandle(handle)
	fmt.Println(syscall.GetLastError())
	addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_READ, 0, 0, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer syscall.UnmapViewOfFile(addr)
	for {
		data := (*Pipe)(unsafe.Pointer(addr))
		time.Sleep(1 * time.Second)
		// fmt.Printf("ava %v cam %v id %v\n", data.Avatar.Position, data.Camera, data.Identity)
		fmt.Printf("str: %s\n", string(data.in[:]))
	}
}

