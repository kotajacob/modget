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

	"git.sr.ht/~kota/modget/curse"
)

// Print some debug info about a curse.File
func DebugFilePrint(file curse.File) {
	fmt.Println(file.FileName)
	fmt.Println(file.FileDate)
	fmt.Println(file.ID)
	for _, fileVersion := range file.GameVersion {
		fmt.Println(fileVersion)
	}
}

// Print some debug info about a curse.Addon
func DebugAddonPrint(addon curse.Addon) {
	fmt.Println(addon.Name)
	fmt.Println(addon.Slug)
	fmt.Println(addon.ID)
	fmt.Printf("%d\n", int(addon.DownloadCount))
}
