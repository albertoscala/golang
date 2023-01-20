package main

import "fmt"

func bruteForceAtom(ini int, end int, msg string, sem chan int) {
	cc := CaeserCipher{}
	for i := ini; i <= end; i++ {
		cc.Key = i
		fmt.Fprintf(&cc, msg)
	}
	sem <- 1
}

func bruteForce(msg string) {
	var done int
	sem := make(chan int, 4)

	go bruteForceAtom(1, 7, msg, sem)
	go bruteForceAtom(8, 14, msg, sem)
	go bruteForceAtom(15, 21, msg, sem)
	go bruteForceAtom(22, 26, msg, sem)

	for done != 4 {
		done += <-sem
	}
}

func main() {
	cc := CaeserCipher{Key: 5}
	fmt.Print("Phrase crypted: ")
	fmt.Fprintf(&cc, "Veni vidi vici")
	fmt.Println("===========================")
	bruteForce(cc.Cache)
}
