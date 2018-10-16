package main

import "fmt"

func RunTowers() {
	nDisks := 10
	doTowers("A", "B", "C", nDisks)
}

func doTowers(from, inter, to string, topN int) {
	if topN == 1 {
		fmt.Printf("Disk 1 from %s to %s\n", from, to)
	} else {
		doTowers(from, to, inter, topN-1)

		fmt.Printf("Disk %d from %s to %s\n", topN, from, to)

		doTowers(inter, from, to, topN-1)
	}

}
