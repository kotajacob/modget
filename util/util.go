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
	"os"

	"git.sr.ht/~kota/modget/curse"
)

// GetModid takes a string, which is meant to be an addon's slug and attempts
// to convert it to a MODID. It returns an error on failure.
func GetModid(s string) (int, error) {
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

// EnsureDir creates a directory if missing.
func EnsureDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}
