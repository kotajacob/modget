/* modget Copyright (C) 2020 Dakota Walsh */
package main

import ()

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	parse(get(308702))
}
