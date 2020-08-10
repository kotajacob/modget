/* modget Copyright (C) 2020 Dakota Walsh */
package main

import (
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	modid, err := strconv.Atoi(os.Args[1])
	check(err)
	parse(get(modid))
}
