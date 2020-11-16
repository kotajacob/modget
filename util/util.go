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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"git.sr.ht/~kota/modget/curse"
	"git.sr.ht/~kota/modget/database"
)

// getModid takes a string, which is meant to be an addon's slug and attempts
// to convert it to a MODID. It returns an error on failure.
func getModid(s string) (int, error) {
	var search curse.Search
	search.GameID = 432     // Set game to minecraft
	search.SectionID = 6    // Set section to mods
	search.SearchFilter = s // Search string
	addons, err := curse.AddonSearch(search)
	if err != nil {
		return 0, err
	}
	for _, addon := range addons {
		if addon.Slug == s {
			return addon.ID, nil
		}
	}
	err = errors.New("Could not find: " + s)
	return 0, err
}

// ToID converts a list of strings to MODIDs
func ToID(s []string) []int {
	var mods []int
	for i := 0; i < len(s); i++ {
		id, err := strconv.Atoi(s[i])
		if err != nil {
			// Attempt to convert slug to modid
			id, err = getModid(s[i])
			if err != nil {
				fmt.Printf("Failed to find: %v\n", s[i])
				os.Exit(1)
			}
		}
		mods = append(mods, id)
	}
	return mods
}

// FindDatabase finds the .modget database at the path. Create the database if missing.
func FindDatabase(path string) (database.Database, error) {
	var db database.Database
	if path == "" {
		path = "."
	}
	err := EnsureDir(path)
	if err != nil {
		return db, err
	}
	path = filepath.Join(path, ".modget")
	db, err = database.Load(path)
	return db, err
}

// EnsureDir creates a directory if missing.
func EnsureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
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
