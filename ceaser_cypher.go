package main

import "fmt"

type CeaserCipher struct {
	Key   int
	Cache string
}

func (cc *CeaserCipher) Write(p []byte) (int, error) {

	for i := 0; i < len(p); i++ {
		c := int(p[i])

		if c >= 65 && c <= 90 {
			c += cc.Key
			if c > 90 {
				c = ((c - 90) + 64)
			}
			p[i] = byte(rune(c))
		}

		if c >= 97 && c <= 122 {
			c += cc.Key
			if c > 122 {
				c = ((c - 122) + 96)
			}
			p[i] = byte(rune(c))
		}
	}

	fmt.Println(string(p))

	return len(p), nil
}

func main() {
	ceaserCipher := CeaserCipher{Key: 1}
	fmt.Fprintf(&ceaserCipher, "zzz")
}
