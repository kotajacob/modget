/* modget Copyright (C) 2020 Dakota Walsh */
package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(get(310806))
}
