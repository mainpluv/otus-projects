package main

import (
	"fmt"

	"github.com/ozgio/strutil"
)

func main() {
	str := "Hello, OTUS!"
	str = strutil.Reverse(str)
	fmt.Println(str)
}
