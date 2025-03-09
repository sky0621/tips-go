package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	printAlloc(PrintUnitKB)
	uuid := "232B4A2F-7007-4D52-BFBB-F19CBB44AD38"
	printAlloc(PrintUnitKB)
	uuid2 := uuid[:36]
	fmt.Println(uuid2)
	printAlloc(PrintUnitKB)
	uuid2 = ""
	runtime.GC()
	printAlloc(PrintUnitKB)
	uuid3 := strings.Clone(uuid)
	fmt.Println(uuid3)
	printAlloc(PrintUnitKB)
	uuid = ""
	runtime.GC()
	fmt.Println(uuid3)
	printAlloc(PrintUnitKB)
	uuid3 = ""
	runtime.GC()
	printAlloc(PrintUnitKB)
}
