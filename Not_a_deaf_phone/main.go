package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
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

	fileoutput, err := os.Create("OUTPUT.TXT")
	if err != nil {
		log.Fatal(err)
	}
	defer fileoutput.Close()

	_, err = fileoutput.Write(bufer.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
