package main

import (
	"fmt"
	"runtime"
)

type PrintUnit int8

const (
	PrintUnitB PrintUnit = iota + 1
	PrintUnitKB
	PrintUnitMB
)

func printAlloc(units ...PrintUnit) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	if len(units) == 0 {
		units = append(units, PrintUnitMB)
	}

	for _, unit := range units {
		switch unit {
		case PrintUnitB:
			fmt.Printf("%d B\n", m.Alloc)
		case PrintUnitKB:
			fmt.Printf("%d KB\n", m.Alloc/1024)
		case PrintUnitMB:
			fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
		default:
			fmt.Println("Unknown unit. Please use MB, KB, or B.")
		}
	}
	fmt.Println()
}
