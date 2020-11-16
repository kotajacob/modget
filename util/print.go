/*
Copyright © 2020 Dakota Walsh

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package util

import (
	"fmt"
	"strings"

	"git.sr.ht/~kota/modget/curse"
)

// ShowMods prints a list of mods that will be installed.
func ShowMods(files []curse.File) {
	fmt.Println("The following mods will be installed:")
	var s string
	var d int
	for _, file := range files {
		s += " " + file.FileName
		d += file.FileLength
	}
	fmt.Printf("%v\n", s)
	fmt.Printf("After this operation, %d of additional disk space will be used.\n", d)
}

// Ask prompts the user with a Yes/No question about continuing
func Ask() bool {
	fmt.Printf("Do you want to continue? [Y/n] ")
	var answer string
	fmt.Scanln(&answer)
	answer = strings.ToLower(answer)
	if answer == "y" || answer == "" {
		return true
	}
	return false
}
