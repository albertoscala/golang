package main

import (
	"fmt"
)

type CeaserCipher struct {
	Key   int
	Cache string
}

func conv_char(b byte, key int) byte {
	c := int(b)

	// Uppercase chars
	if c >= 65 && c <= 90 {
		c += key
		if c > 90 {
			c = ((c - 90) + 64)
		}
		return byte(rune(c))
	}

	// Lowercase chars
	if c >= 97 && c <= 122 {
		c += key
		if c > 122 {
			c = ((c - 122) + 96)
		}
		return byte(rune(c))
	}

	return b
}

func (cc *CeaserCipher) Write(p []byte) (int, error) {

	for i := 0; i < len(p); i++ {
		p[i] = conv_char(p[i], cc.Key)
	}

	s := string(p)

	fmt.Println(s)
	cc.Cache = s

	return len(p), nil
}

func bruteForceAtom(ini int, end int, msg string, sem chan int) {
	cc := CeaserCipher{}
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
	cc := CeaserCipher{Key: 5}
	fmt.Print("Phrase crypted: ")
	fmt.Fprintf(&cc, "Veni vidi vici")
	fmt.Println("===========================")
	bruteForce(cc.Cache)
}
