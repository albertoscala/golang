package main

import (
	"time"
)

// NON HO CAPITO COME DOVREI AGGIUNGERE LA DATA DI NASCITA AD UN NUMERO
func yourBirthYear(number int64) int64 {
	return number + 2002
}

func temporalCoordinatesDerivation(yourBirthYear func(int64) int64) func(time.Time) (int64, int64) {
	x := 16342 * yourBirthYear(52634)
	y := 88345 * yourBirthYear(35237)
	return func(t time.Time) (int64, int64) {
		return x + t.Unix(), y + t.Unix()
	}
}

func main() {

}
