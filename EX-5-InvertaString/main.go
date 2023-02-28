package main

import "fmt"

func main() {
	str := "Hello, world!"
	inversao := inverter(str)
	fmt.Println(inversao)
}

func inverter(str string) string {
	slicePalavra := []rune(str)
	for i, x := 0, len(slicePalavra)-1; i < len(slicePalavra)/2; i, x = i+1, x-1 {
		slicePalavra[i], slicePalavra[x] = slicePalavra[x], slicePalavra[i]
	}
	return string(slicePalavra)
}
