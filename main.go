package main

import (
	"fmt"
	"github.com/winjeg/demo/ipc"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("args must be 2")
		return
	}

	if strings.EqualFold("r", args[1]) {
		ipc.Read()
	} else if strings.EqualFold("w", args[1]) {
		ipc.Write()
	}
}
