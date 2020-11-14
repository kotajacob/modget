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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"git.sr.ht/~kota/modget/curse"
)

// Default file writing mode
var Mode os.FileMode = 0644

// Get a []byte from a file
func get(filename string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Overwrite a file with a []byte
func put(filename string, dat []byte) error {
	err := ioutil.WriteFile(filename, dat, Mode)
	if err != nil {
		return err
	}
	return nil
}

// Read the whole database returning a []curse.File
func DatabaseRead(p string) ([]curse.File, error) {
	var files []curse.File
	b, err := get(p)
	if os.IsNotExist(err) {
		fmt.Println("FILE DOES NOT EXIST")
		return files, nil
	} else if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Add a new mod to the database. Loads the database into memory, checks if the
// mod is already in the database, adds or updates it depending.
func DatabaseAdd(file curse.File, p string) error {
	files, err := DatabaseRead(p)
	if err != nil {
		return err
	}
	files = append(files, file)
	b, err := json.Marshal(files)
	if err != nil {
		return err
	}
	err = put(p, b)
	if err != nil {
		return err
	}
	return nil
}
