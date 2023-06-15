package main

// import (
// 	"fmt"
// 	"strings"
//
// )

// func main() {
// 	var a = "11"
// 	var b = "1"

// 	fmt.Println(addBinary(a, b))
// }

// func addBinary(a string, b string) string {
// 	as := strings.Split(a, "")
// 	bs := strings.Split(b, "")
// 	o := ""
// 	n := len(as) - 1
// 	for i := len(bs) - 1; i >= 0; i-- {
// 		if (as[n-i] == "1") && (bs[i] == "1") {
// 			as[n-i] = "0"
// 			if n-i == 0 {
// 				as = append(as[0:], as[1:])
// 			}
// 			as = pp(as, n-i-1)

// 			continue
// 		}
// 		if (as[i] == "0") && (bs[i] == "1") {
// 			as[i] = "1"
// 			continue
// 		}
// 		if (as[i] == "1") && (bs[i] == "0") {
// 			as[i] = "1"
// 			continue
// 		}
// 		if (as[i] == "0") && (bs[i] == "0") {
// 			as[i] = "0"
// 			continue
// 		}
// 	}

// 	for i := range as {
// 		o += as[i]
// 	}

// 	return o
// }

// func pp(as []string, a int) []string {
// 	if as[a] == "1" {
// 		as[a] = "0"
// 		if a == 0 {
// 			copy(as[0:], as[1:])
// 		}
// 		return as
// 	}
// 	if as[a] == "0" {
// 		as[a] = "1"
// 		return as
// 	}
// 	return as
// }
