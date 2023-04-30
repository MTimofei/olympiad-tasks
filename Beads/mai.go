package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var n uint64

	fileinput, err := os.Open("INPUT.TXT")
	if err != nil {
		log.Fatal(err)
	}
	defer fileinput.Close()

	bufer := bytes.Buffer{}
	scan := bufio.NewScanner(fileinput)

	for scan.Scan() {
		bufer.WriteString(scan.Text())
	}

	x, err := strconv.Atoi(bufer.String())
	if err != nil {
		log.Fatal(err)
	}

	n = uint64(x) + 1

	fileuotput, err := os.Create("OUTPUT.TXT")
	if err != nil {
		log.Fatal(err)
	}
	_, err = fileuotput.WriteString(fmt.Sprintf("%d", n))
	if err != nil {
		log.Fatal(err)
	}
}
