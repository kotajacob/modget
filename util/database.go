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
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"

	"git.sr.ht/~kota/modget/curse"
)

// Default file writing mode
var Mode os.FileMode = 0644

type Database struct {
	Version   int
	Minecraft string
	Loader    string
	Files     []curse.File
}

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

// Load database into memory.
func DatabaseLoad(p string) (Database, error) {
	var db Database
	b, err := get(p)
	if os.IsNotExist(err) {
		return db, nil
	} else if err != nil {
		return db, err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&db)
	if err != nil {
		return db, err
	}
	return db, nil
}

// Add a new mod to the database. Loads the database into memory, checks if the
// mod is already in the database, adds or updates it depending.
func DatabaseAdd(file curse.File, p string) error {
	db, err := DatabaseLoad(p)
	if err != nil {
		return err
	}
	db.Files = append(db.Files, file)
	stream := &bytes.Buffer{}
	en := gob.NewEncoder(stream)
	err = en.Encode(db)
	if err != nil {
		return err
	}
	err = put(p, stream.Bytes())
	if err != nil {
		return err
	}
	return nil
}
