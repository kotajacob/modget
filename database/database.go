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
	"fmt"
	"io/ioutil"
	"os"

	"git.sr.ht/~kota/modget/curse"
)

// Mode sets the default file writing mode
var Mode os.FileMode = 0644

// Mod contains relevant information about a mod that needs stored in the
// database.
type Mod struct {
	Name            string
	Authors         []curse.Author
	Summary         string
	DownloadCount   float64
	FileID          int
	FileName        string
	FileLength      int
	FileDate        string
	GameVersion     []string
	DownloadURL     string
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
	Mods      map[int]Mod
}

// NewMod creates a new Mod from an Addon and File.
func NewMod(addon curse.Addon, file curse.File) Mod {
	var mod Mod
	mod.Name = addon.Name
	mod.Authors = addon.Authors
	mod.Summary = addon.Summary
	mod.DownloadCount = addon.DownloadCount
	mod.FileID = file.ID
	mod.FileName = file.FileName
	mod.FileLength = file.FileLength
	mod.FileDate = file.FileDate
	mod.GameVersion = file.GameVersion
	mod.DownloadURL = file.DownloadURL
	mod.Status = addon.Status
	mod.Slug = addon.Slug
	mod.PopularityScore = addon.PopularityScore
	mod.PrimaryLanguage = addon.PrimaryLanguage
	mod.DateModified = addon.DateModified
	mod.DateCreated = addon.DateCreated
	mod.DateReleased = addon.DateReleased
	return mod
}

// Add inserts a mod into a Database from the mod's Addon and File.
func (db *Database) Add(ID int, mod Mod) {
	db.Mods[ID] = mod
}

// Del removes a Mod from a Database by its index
func (db *Database) Del(ID int) {
	delete(db.Mods, ID)
}

// Write saves a Database to a file at a path.
func (db *Database) Write(p string) error {
	stream := &bytes.Buffer{}
	en := gob.NewEncoder(stream)
	err := en.Encode(db)
	if err != nil {
		return fmt.Errorf("encoding database: %v", err)
	}
	err = ioutil.WriteFile(p, stream.Bytes(), Mode)
	if err != nil {
		return fmt.Errorf("writing database: %v", err)
	}
	return nil
}

// Load a Database from a path.
func Load(p string) (*Database, error) {
	var db Database
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading database: %v", err)
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&db)
	if err != nil {
		return nil, fmt.Errorf("decoding database: %v", err)
	}
	return &db, err
}
