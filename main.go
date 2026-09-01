package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	if !menuAwal(reader) {
		return
	}

	dummmy()
	menuUtama(reader)
}