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

package database

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"

	"git.sr.ht/~kota/modget/curse"
)

// Mode sets the default file writing mode
var Mode os.FileMode = 0644

// Mod contains relevant information about a mod that needs stored in the
// database.
type Mod struct {
	ID              int
	Name            string
	Authors         []curse.Author
	Summary         string
	DownloadCount   float64
	FileID          int
	FileName        string
	FileLength      int
	FileDate        string
	Status          int
	Slug            string
	PopularityScore float64
	PrimaryLanguage string
	DateModified    string
	DateCreated     string
	DateReleased    string
}

// Database is the format of the .modget file.
type Database struct {
	Version   string
	Minecraft string
	Loader    string
	Mods      []Mod
}

// AddMod adds a mod to a Database. Requires the mod's Addon and File.
func (db Database) AddMod(addon curse.Addon, file curse.File) Database {
	var mod Mod
	mod.ID = addon.ID
	mod.Name = addon.Name
	mod.Authors = addon.Authors
	mod.Summary = addon.Summary
	mod.DownloadCount = addon.DownloadCount
	mod.FileID = file.ID
	mod.FileName = file.FileName
	mod.FileLength = file.FileLength
	mod.FileDate = file.FileDate
	mod.Status = addon.Status
	mod.Slug = addon.Slug
	mod.PopularityScore = addon.PopularityScore
	mod.PrimaryLanguage = addon.PrimaryLanguage
	mod.DateModified = addon.DateModified
	mod.DateCreated = addon.DateCreated
	mod.DateReleased = addon.DateReleased
	db.Mods = append(db.Mods, mod)
	return db
}

// DelMod removed a Mod from a Database by its index
func (db Database) DelMod(ID int) Database {
	mods := db.Mods
	for i, mod := range mods {
		if mod.ID == ID {
			db.Mods = remove(db.Mods, i)
		}
	}
	return db
}

// Write saves a Database to a file at a path.
func (db Database) Write(p string) error {
	stream := &bytes.Buffer{}
	en := gob.NewEncoder(stream)
	err := en.Encode(db)
	if err != nil {
		return err
	}
	err = put(p, stream.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// Create a Database.
func Create(version, minecraft, loader string) Database {
	var db Database
	db.Version = version
	db.Minecraft = minecraft
	db.Loader = loader
	return db
}

// Load a Database from a path.
func Load(p string) (Database, error) {
	var db Database
	b, err := get(p)
	if err != nil {
		return db, err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&db)
	return db, err
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

// Remove an element from a slice by index
func remove(slice []Mod, i int) []Mod {
	return append(slice[:i], slice[i+1:]...)
}
