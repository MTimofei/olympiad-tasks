package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a int64

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

	split := strings.Split(bufer.String(), " ")

	for i := range split {
		x, err := strconv.Atoi(split[i])
		i := int64(x)
		if err != nil {
			log.Fatal(err)
		}
		a += i

	}

	fileoutput, err := os.Create("OUTPUT.TXT")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fileoutput.WriteString(fmt.Sprintf("%d", a))
	if err != nil {
		log.Fatal(err)
	}

}
