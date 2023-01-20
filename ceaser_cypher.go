package main

import "fmt"

type CaeserCipher struct {
	Key   int
	Cache string
}

func convChar(b byte, key int) byte {
	c := int(b)

	// Uppercase chars
	if c >= 65 && c <= 90 {
		c += key
		if c > 90 {
			c = (c - 90) + 64
		}
		return byte(rune(c))
	}

	// Lowercase chars
	if c >= 97 && c <= 122 {
		c += key
		if c > 122 {
			c = (c - 122) + 96
		}
		return byte(rune(c))
	}

	return b
}

func (cc *CaeserCipher) Write(p []byte) (int, error) {

	for i := 0; i < len(p); i++ {
		p[i] = convChar(p[i], cc.Key)
	}

	s := string(p)

	fmt.Println(s)
	cc.Cache = s

	return len(p), nil
}

func main() {
	caesarCypher := CaeserCipher{Key: 3}
	fmt.Fprintf(&caesarCypher, "tl:dr")
	fmt.Println(caesarCypher.Cache)
	fmt.Fprintf(&caesarCypher, "alberto")
}
